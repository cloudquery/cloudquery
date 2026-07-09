package cmd

import (
	"bytes"
	"cmp"
	"context"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strings"

	cqapi "github.com/cloudquery/cloudquery-api-go"
	cqauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/analytics"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/platform"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const (
	initShort = `Generate a configuration file for a sync`
	initLong  = `Generate a configuration file for a sync

### Modes

The ` + "`init`" + ` command operates in one of three modes depending on your authentication state and flags:

**AI-assisted mode** (default when logged in)

Activates when you are logged in to a team (` + "`cloudquery login`" + `) and don't specify ` + "`--source`" + ` or ` + "`--destination`" + `. Launches an interactive AI chat session that walks you through the setup process — selecting integrations, generating YAML configuration files, testing connections, and giving you some example queries.

Type ` + "`exit`" + ` or ` + "`quit`" + ` to end the conversation. Use ` + "`--resume-conversation`" + ` to continue a previous session instead of starting a new one.

**Basic interactive mode**

Activates when you pass ` + "`--disable-ai`" + `, or as a fallback if the AI assistant is unavailable. Presents a searchable picker to select source and destination integrations, then generates a configuration file from their default templates.

**Non-interactive mode**

Activates when both ` + "`--source`" + ` and ` + "`--destination`" + ` are specified. Generates the configuration file directly without prompts.

Authentication via ` + "`cloudquery login`" + ` is required for AI-assisted and basic interactive modes.`
	initExample = `# Display prompts to select source and destination plugins and generate a configuration file from them
cloudquery init
# Generate a configuration file for a sync from aws to bigquery
cloudquery init --source aws --destination bigquery
# Display a prompt to select a source plugin and generate a configuration file for a sync from it to bigquery
cloudquery init --destination bigquery
# Display a prompt to select a destination plugin and generate a configuration file for a sync from aws to it
cloudquery init --source aws
# Accept all defaults and generate a configuration file for a sync from the first source and destination plugins
cloudquery init --yes`
)

var (
	sourcesOrder      = []string{"aws", "azure", "gcp"}
	destinationsOrder = []string{"postgresql", "bigquery", "s3"}
	bold              = color.New(color.Bold)
	successful        = color.New(color.Bold, color.FgGreen)
	link              = color.New(color.Bold, color.FgCyan)
	errorColor        = color.New(color.Bold, color.FgRed)
)

func newCmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   initShort,
		Long:    initLong,
		Example: initExample,
		Args:    cobra.ExactArgs(0),
		RunE:    initCmd,
	}
	cmd.Flags().String("source", "", "Source plugin name or path")
	cmd.Flags().String("destination", "", "Destination plugin name or path")
	cmd.Flags().String("spec-path", "", "Output spec file path")
	cmd.Flags().Bool("yes", false, "Accept all defaults")
	cmd.Flags().Bool("disable-ai", false, "Disable AI assistant")
	cmd.Flags().Bool("disable-platform", false, "Skip CloudQuery Platform sync scaffolding")
	cmd.Flags().Bool("resume-conversation", false, "Resume existing AI conversation instead of starting a new one")
	return cmd
}

func normalizePluginPath(pluginNameOrPath string) (string, error) {
	parts := strings.Split(pluginNameOrPath, "/")
	if len(parts) == 1 {
		return "cloudquery/" + pluginNameOrPath, nil
	}
	if len(parts) != 2 {
		return "", errors.New("invalid plugin path")
	}
	return pluginNameOrPath, nil
}

func parseFlags(cmd *cobra.Command) (source, destination, specPath string, acceptDefaults, disableAI, resumeConversation, disablePlatform bool, allErrors error) {
	source, err := cmd.Flags().GetString("source")
	allErrors = errors.Join(allErrors, err)
	if source != "" {
		source, err = normalizePluginPath(source)
		allErrors = errors.Join(allErrors, err)
	}
	destination, err = cmd.Flags().GetString("destination")
	allErrors = errors.Join(allErrors, err)
	if destination != "" {
		destination, err = normalizePluginPath(destination)
		allErrors = errors.Join(allErrors, err)
	}
	specPath, err = cmd.Flags().GetString("spec-path")
	allErrors = errors.Join(allErrors, err)

	acceptDefaults, err = cmd.Flags().GetBool("yes")
	allErrors = errors.Join(allErrors, err)

	disableAI, err = cmd.Flags().GetBool("disable-ai")
	allErrors = errors.Join(allErrors, err)

	resumeConversation, err = cmd.Flags().GetBool("resume-conversation")
	allErrors = errors.Join(allErrors, err)

	disablePlatform, err = cmd.Flags().GetBool("disable-platform")
	allErrors = errors.Join(allErrors, err)
	return source, destination, specPath, acceptDefaults, disableAI, resumeConversation, disablePlatform, allErrors
}

func pluginFilter(pluginPath string, kind cqapi.PluginKind) func(plugin cqapi.ListPlugin) bool {
	return func(plugin cqapi.ListPlugin) bool {
		return plugin.TeamName+"/"+plugin.Name == pluginPath && plugin.Kind == kind && plugin.LatestVersion != nil
	}
}

func pluginName(plugin cqapi.ListPlugin, _ int) string {
	return plugin.Name
}

func officialReleasedPluginsByKind(kind cqapi.PluginKind) func(plugin cqapi.ListPlugin, _ int) bool {
	return func(plugin cqapi.ListPlugin, _ int) bool {
		return plugin.Kind == kind && plugin.Official && plugin.ReleaseStage != cqapi.PluginReleaseStageComingSoon && plugin.LatestVersion != nil
	}
}

func pluginsSorter(prioritySlice []string) func(a, b cqapi.ListPlugin) int {
	return func(a, b cqapi.ListPlugin) int {
		indexOfA := lo.IndexOf(prioritySlice, a.Name)
		indexOfB := lo.IndexOf(prioritySlice, b.Name)
		if indexOfA == -1 && indexOfB != -1 {
			return 1
		}
		if indexOfA != -1 && indexOfB == -1 {
			return -1
		}
		if indexOfA == -1 && indexOfB == -1 {
			return cmp.Compare(a.Name, b.Name)
		}
		return cmp.Compare(indexOfA, indexOfB)
	}
}

func extractYamlFromMarkdownCodeBlock(markdown string) string {
	re := regexp.MustCompile("```yaml.*?\n([\\s\\S]+?)\n```")

	matches := re.FindStringSubmatch(markdown)
	if len(matches) < 2 {
		return ""
	}

	return matches[1]
}

func defaultConfigForPlugin(plugin cqapi.ListPlugin) *strings.Builder {
	tmpl := `kind: {{.Kind}}
spec:
  name: {{.Name}}
  path: {{.TeamName}}/{{.Name}}
  version: {{.LatestVersion}}
`
	var buf bytes.Buffer
	t := template.Must(template.New("config").Parse(tmpl))
	_ = t.Execute(&buf, plugin)

	sb := strings.Builder{}
	sb.WriteString(buf.String())
	return &sb
}

// withRecommendedTables sets the source spec's `tables` selection to the
// platform's recommended tables, overriding whatever form the example config used
// — wildcard `['*']`, an inline list like `["aws_s3_buckets"]`, or a block list —
// and adding a `tables:` key if the example omitted one. It edits the parsed YAML
// node tree in place and re-marshals, so the rest of the scaffold (comments, key
// order, auth-field stubs) is preserved. Returns the spec unchanged when the
// recommended set is empty, the input isn't parseable, or it has no source
// `spec:` mapping to write into.
func withRecommendedTables(yamlSpec string, tables []string) string {
	if len(tables) == 0 {
		return yamlSpec
	}
	// Decode every document — an example config's ```yaml block may ship more than
	// one (e.g. a companion destination doc). Editing only the first and dropping
	// the rest would truncate the spec, so round-trip them all.
	dec := yaml.NewDecoder(strings.NewReader(yamlSpec))
	var docs []*yaml.Node
	for {
		var doc yaml.Node
		err := dec.Decode(&doc)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return yamlSpec
		}
		docs = append(docs, &doc)
	}
	changed := false
	for _, doc := range docs {
		if setSpecTables(doc, tables) {
			changed = true
		}
	}
	if !changed {
		return yamlSpec
	}
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2) // match the example configs' 2-space convention
	for _, doc := range docs {
		if err := enc.Encode(doc); err != nil {
			return yamlSpec
		}
	}
	if err := enc.Close(); err != nil {
		return yamlSpec // don't return a half-flushed buffer
	}
	return buf.String()
}

// setSpecTables sets the source spec's `tables` to a block sequence of the given
// table names — replacing the existing value node in place, or adding a `tables:`
// key when the spec has none. Only touches a `kind: source` document (so a
// companion destination doc in a multi-doc example isn't given a bogus `tables:`).
// Returns false when the doc isn't a source or has no `spec:` mapping.
func setSpecTables(doc *yaml.Node, tables []string) bool {
	if doc.Kind != yaml.DocumentNode || len(doc.Content) == 0 {
		return false
	}
	root := doc.Content[0]
	if kind := mappingValue(root, "kind"); kind == nil || kind.Value != "source" {
		return false
	}
	spec := mappingValue(root, "spec")
	if spec == nil || spec.Kind != yaml.MappingNode {
		return false
	}
	items := make([]*yaml.Node, 0, len(tables))
	for _, t := range tables {
		items = append(items, &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: t})
	}
	if tablesNode := mappingValue(spec, "tables"); tablesNode != nil {
		tablesNode.Kind = yaml.SequenceNode
		tablesNode.Tag = "!!seq"
		tablesNode.Style = 0 // block, not flow
		tablesNode.Value = ""
		tablesNode.Content = items
		return true
	}
	// No `tables` key in the example — add one.
	spec.Content = append(spec.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: "tables"},
		&yaml.Node{Kind: yaml.SequenceNode, Tag: "!!seq", Content: items},
	)
	return true
}

// mappingValue returns the value node for key in a mapping node, or nil.
func mappingValue(n *yaml.Node, key string) *yaml.Node {
	if n == nil || n.Kind != yaml.MappingNode {
		return nil
	}
	for i := 0; i+1 < len(n.Content); i += 2 {
		if n.Content[i].Value == key {
			return n.Content[i+1]
		}
	}
	return nil
}

func configForSourcePlugin(source cqapi.ListPlugin, version *cqapi.PluginVersionDetails) string {
	exampleConfig := extractYamlFromMarkdownCodeBlock(version.ExampleConfig)
	if exampleConfig != "" {
		return exampleConfig
	}

	defaultConfig := defaultConfigForPlugin(source)
	defaultConfig.WriteString("  tables: ['*']\n")
	defaultConfig.WriteString("  destinations: ['DESTINATION_NAME']")
	return defaultConfig.String()
}

func configForDestinationPlugin(destination cqapi.ListPlugin, version *cqapi.PluginVersionDetails) string {
	exampleConfig := extractYamlFromMarkdownCodeBlock(version.ExampleConfig)
	if exampleConfig != "" {
		return exampleConfig
	}

	defaultConfig := defaultConfigForPlugin(destination)
	return defaultConfig.String()
}

// unsupportedPlatformSourceError returns an error when a platform tenant doesn't
// support the given source path (team/name), else nil. An empty or nil supported
// set (no platform tenant) never rejects.
func unsupportedPlatformSourceError(source string, supported map[string]string) error {
	if len(supported) == 0 {
		return nil
	}
	if _, ok := supported[source]; ok {
		return nil
	}
	return fmt.Errorf("source plugin %q is not supported by your CloudQuery Platform", source)
}

// selectSource prompts for a source plugin. When supportedPaths is non-empty
// (a CloudQuery Platform tenant — the caller requires a non-empty set there), the
// list is restricted to the sources the platform supports (its `team/name` keys),
// so the picker never offers plugins, e.g. database sources, a platform sync can't
// use. An empty or nil supportedPaths (no platform tenant) leaves it unfiltered.
func selectSource(allPlugins []cqapi.ListPlugin, acceptDefaults bool, supportedPaths map[string]string) (string, error) {
	officialSources := lo.Filter(allPlugins, officialReleasedPluginsByKind(cqapi.PluginKindSource))
	if len(supportedPaths) > 0 {
		officialSources = lo.Filter(officialSources, func(p cqapi.ListPlugin, _ int) bool {
			_, ok := supportedPaths[p.TeamName+"/"+p.Name]
			return ok
		})
	}
	if len(officialSources) == 0 {
		if len(supportedPaths) > 0 {
			// Platform tenant, but none of its supported sources are official
			// released plugins (near-impossible) — give the same escape hatch.
			return "", errors.New("none of the source plugins your CloudQuery Platform supports are available — please try again, or pass --disable-platform to scaffold a regular source + destination config")
		}
		return "", errors.New("no source plugins available to select")
	}
	slices.SortStableFunc(officialSources, pluginsSorter(sourcesOrder))
	if acceptDefaults {
		return officialSources[0].Name, nil
	}

	prompt := promptui.Select{
		Label:             "Select Source Plugin",
		Items:             lo.Map(officialSources, pluginName),
		Stdin:             os.Stdin,
		Size:              10,
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			return strings.Contains(officialSources[index].Name, input)
		},
	}

	_, source, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("source prompt failed %w", err)
	}

	return source, nil
}

func selectDestination(allPlugins []cqapi.ListPlugin, acceptDefaults bool) (string, error) {
	officialDestinations := lo.Filter(allPlugins, officialReleasedPluginsByKind(cqapi.PluginKindDestination))
	slices.SortStableFunc(officialDestinations, pluginsSorter(destinationsOrder))
	if acceptDefaults {
		return officialDestinations[0].Name, nil
	}

	prompt := promptui.Select{
		Label:             "Select Destination Plugin",
		Items:             lo.Map(officialDestinations, pluginName),
		Stdin:             os.Stdin,
		Size:              10,
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			return strings.Contains(officialDestinations[index].Name, input)
		},
	}

	_, destination, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("destination prompt failed %w", err)
	}

	return destination, nil
}

func linkForPlugin(plugin cqapi.ListPlugin) string {
	return link.Sprintf("https://www.cloudquery.io/hub/plugins/%s/%s/%s", plugin.Kind, plugin.TeamName, plugin.Name)
}

// writePlatformSourceOnlySpec scaffolds a source-only spec for a user with a
// CloudQuery Platform tenant: no destination block, since the CLI auto-injects
// the `platform` destination at sync time. It wires the source to that reserved
// destination name and tells the user where the data will land.
func writePlatformSourceOnlySpec(ctx context.Context, apiClient *cqapi.ClientWithResponses, sourcePlugin cqapi.ListPlugin, specPath, platformURL string, tenantInit *platform.TenantInit) error {
	sourcePath := sourcePlugin.TeamName + "/" + sourcePlugin.Name
	// Prefer the platform-pinned source version over the hub's latest, so the
	// scaffolded spec targets a version the tenant will accept (the same version
	// the sync-time gate enforces). Unpinned sources keep LatestVersion.
	if pinned := tenantInit.PinnedSourceVersions[sourcePath]; pinned != "" {
		sourcePlugin.LatestVersion = &pinned
	}
	fmt.Printf("Getting configuration for source plugin %s...\n", bold.Sprintf("%s/%s@%s", sourcePlugin.TeamName, sourcePlugin.Name, *sourcePlugin.LatestVersion))
	sourceVersion, err := api.GetPluginVersion(apiClient, sourcePlugin.TeamName, sourcePlugin.Kind, sourcePlugin.Name, *sourcePlugin.LatestVersion)
	if err != nil {
		return fmt.Errorf("failed to get source plugin %s/%s@%s version %w", sourcePlugin.TeamName, sourcePlugin.Name, *sourcePlugin.LatestVersion, err)
	}

	if specPath == "" {
		specPath = sourcePlugin.Name + "_to_platform.yaml"
	}
	fmt.Printf("Writing spec to %s...\n", bold.Sprint(specPath))
	// Wire the source to the reserved `platform` destination name; the CLI adds
	// the destination itself at sync time, so no destination block is written.
	sourceConfig := configForSourcePlugin(sourcePlugin, sourceVersion)
	yamlSpec := strings.ReplaceAll(sourceConfig, "DESTINATION_NAME", "platform")
	// Set the source's tables to what the platform recommends, so the sync
	// populates the tables the platform ingests. No recommendation → leave the
	// example config's tables as-is.
	if recommended := tenantInit.RecommendedTables(ctx, log.Logger, sourcePath); len(recommended) > 0 {
		yamlSpec = withRecommendedTables(yamlSpec, recommended)
	}
	if err := os.WriteFile(specPath, []byte(yamlSpec), 0644); err != nil {
		return fmt.Errorf("failed to write spec file %w", err)
	}

	successful.Println("Sync spec file generated successfully!")
	fmt.Println()
	// platformURL is empty for a legacy CQ_PLATFORM_TOKEN with no url claim; omit
	// the "at <url>" tail rather than printing a blank one.
	if platformURL != "" {
		fmt.Printf("This sync will write to your CloudQuery Platform at %s\n", bold.Sprint(platformURL))
	} else {
		fmt.Println("This sync will write to your CloudQuery Platform.")
	}
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Printf("1. Review %s and fill in the source's authentication details:\n", bold.Sprint(specPath))
	fmt.Printf("   %s: %s\n", bold.Sprint(sourcePlugin.DisplayName), linkForPlugin(sourcePlugin))
	fmt.Println("2. Run the sync:")
	bold.Printf("cloudquery sync %s\n", specPath)
	return nil
}

func initCmd(cmd *cobra.Command, args []string) (initCommandError error) {
	ctx := cmd.Context()
	source, destination, specPath, acceptDefaults, disableAI, resumeConversation, disablePlatform, err := parseFlags(cmd)
	analytics.TrackInitStarted(ctx, invocationUUID.UUID, analytics.InitEvent{
		Source:         source,
		Destination:    destination,
		AcceptDefaults: acceptDefaults,
		SpecPath:       specPath,
		Error:          err,
	})
	defer func() {
		analytics.TrackInitCompleted(ctx, invocationUUID.UUID, analytics.InitEvent{
			Source:         source,
			Destination:    destination,
			AcceptDefaults: acceptDefaults,
			SpecPath:       specPath,
			Error:          initCommandError,
		})
	}()
	if err != nil {
		return err
	}

	authClient := cqauth.NewTokenClient()
	token, err := authClient.GetToken()
	var user *cqapi.User
	if err == nil {
		user, _ = auth.GetUser(cmd.Context(), token)
	}

	team, _ := auth.GetTeamForToken(cmd.Context(), token)

	// If the user has a CloudQuery Platform tenant (cloud login, or a
	// CQ_PLATFORM_TOKEN), scaffold a source-only spec that targets the platform
	// destination (auto-injected at sync time), skipping the destination prompt
	// and AI. --disable-platform opts out (normal source+destination spec); an
	// explicit --destination also takes the normal path.
	platformURL, platformTenant := "", false
	var tenantInit *platform.TenantInit
	if !disablePlatform {
		// One tenant lookup + mint yields the URL to report, the pinned source
		// versions to scaffold, and a session reused for the recommended-tables
		// lookup below. A detected tenant whose session can't be minted can't sync,
		// so fail rather than scaffold a spec that would break later.
		var err error
		tenantInit, err = platform.DetectTenantForInit(ctx, log.Logger, token.Value, team)
		if err != nil {
			return fmt.Errorf("failed to set up CloudQuery Platform sync (use --disable-platform to scaffold a regular source + destination config): %w", err)
		}
		if tenantInit != nil {
			platformTenant = true
			platformURL = tenantInit.APIURL
		}
	}

	apiClient, err := api.NewAnonymousClient()
	var apiClientWithoutRetries *cqapi.ClientWithResponses
	if err != nil {
		return err
	}
	if user != nil {
		apiClient, err = api.NewClient(token.Value)
		if err != nil {
			return err
		}

		apiClientWithoutRetries, err = api.NewClient(token.Value, cqapi.WithHTTPClient(http.DefaultClient))
		if err != nil {
			return err
		}
	}

	// Check if user and team are set, and if so, run AI command. Platform-tenant
	// users skip AI: their spec is source-only (the platform destination is
	// auto-injected), which the AI flow doesn't produce.
	if user != nil && team != "" && !disableAI && source == "" && destination == "" && !platformTenant {
		err := api.NewConversation(ctx, apiClientWithoutRetries, team, resumeConversation)
		if err != nil && err != api.ErrDisabled {
			return err
		}
		if err != api.ErrDisabled {
			// User and team are set, endpoint is not FF disabled, proceed to run the AI command
			err := aiCmd(ctx, apiClient, team, resumeConversation)

			// This is unintuitive:
			// - if AI works out, we're done
			// - if AI fails, we output an obfuscated error and fallback to basic interactive mode
			if err == nil {
				return nil
			}
			errorColor.Println("There was an issue with the AI assistant. Falling back to basic interactive mode...")
			fmt.Println()
		}
	} else if (user == nil || team == "") && source == "" && destination == "" && !disableAI && !platformTenant {
		return errors.New("authentication required for interactive mode. Please run `cloudquery login` first, or supply source and destination plugins, or else use the --disable-ai flag to run basic interactive mode")
	}

	fmt.Println("Fetching plugins...")
	allPlugins, err := api.ListAllPlugins(apiClient)
	if err != nil {
		return err
	}

	// On a platform tenant, restrict sources to what the platform supports (its
	// supported-source-versions). This is required, not best-effort: offering a
	// source the platform can't ingest scaffolds a config that only fails later at
	// the sync-time gate, so if the support list is unavailable, stop with an
	// actionable error rather than silently listing everything.
	var supportedSourcePaths map[string]string
	if platformTenant {
		if len(tenantInit.PinnedSourceVersions) == 0 {
			return errors.New("couldn't determine which source plugins your CloudQuery Platform supports — please try again, or pass --disable-platform to scaffold a regular source + destination config")
		}
		supportedSourcePaths = tenantInit.PinnedSourceVersions
	}

	var notFoundPluginsErrors error
	if source != "" {
		sourcePluginFilter := pluginFilter(source, cqapi.PluginKindSource)
		if !lo.SomeBy(allPlugins, sourcePluginFilter) {
			notFoundPluginsErrors = errors.Join(notFoundPluginsErrors, fmt.Errorf("source plugin %q not found", source))
		} else if err := unsupportedPlatformSourceError(source, supportedSourcePaths); err != nil {
			notFoundPluginsErrors = errors.Join(notFoundPluginsErrors, err)
		}
	}
	if destination != "" {
		destinationPluginFilter := pluginFilter(destination, cqapi.PluginKindDestination)
		destinationFound := lo.SomeBy(allPlugins, destinationPluginFilter)
		if !destinationFound {
			notFoundPluginsErrors = errors.Join(notFoundPluginsErrors, fmt.Errorf("destination plugin %q not found", destination))
		}
	}

	if notFoundPluginsErrors != nil {
		return notFoundPluginsErrors
	}

	if source == "" {
		source, err = selectSource(allPlugins, acceptDefaults, supportedSourcePaths)
		if err != nil {
			return err
		}
		source, _ = normalizePluginPath(source)
	}
	_, sourceIndex, _ := lo.FindIndexOf(allPlugins, pluginFilter(source, cqapi.PluginKindSource))

	// Platform tenant + no explicit destination → scaffold a source-only spec;
	// the CLI auto-injects the platform destination at sync time.
	if platformTenant && destination == "" {
		return writePlatformSourceOnlySpec(ctx, apiClient, allPlugins[sourceIndex], specPath, platformURL, tenantInit)
	}

	if destination == "" {
		destination, err = selectDestination(allPlugins, acceptDefaults)
		if err != nil {
			return err
		}
		destination, _ = normalizePluginPath(destination)
	}
	_, destinationIndex, _ := lo.FindIndexOf(allPlugins, pluginFilter(destination, cqapi.PluginKindDestination))

	sourcePlugin := allPlugins[sourceIndex]
	fmt.Printf("Getting configuration for source plugin %s...\n", bold.Sprintf("%s/%s@%s", sourcePlugin.TeamName, sourcePlugin.Name, *sourcePlugin.LatestVersion))
	sourceVersion, err := api.GetPluginVersion(apiClient, sourcePlugin.TeamName, sourcePlugin.Kind, sourcePlugin.Name, *sourcePlugin.LatestVersion)
	if err != nil {
		return fmt.Errorf("failed to get source plugin %s/%s@%s version %w", sourcePlugin.TeamName, sourcePlugin.Name, *sourcePlugin.LatestVersion, err)
	}

	destinationPlugin := allPlugins[destinationIndex]
	fmt.Printf("Getting configuration for destination plugin %s...\n", bold.Sprintf("%s/%s@%s", destinationPlugin.TeamName, destinationPlugin.Name, *destinationPlugin.LatestVersion))
	destinationVersion, err := api.GetPluginVersion(apiClient, destinationPlugin.TeamName, destinationPlugin.Kind, destinationPlugin.Name, *destinationPlugin.LatestVersion)
	if err != nil {
		return fmt.Errorf("failed to get destination plugin %s/%s@%s version %w", destinationPlugin.TeamName, destinationPlugin.Name, *destinationPlugin.LatestVersion, err)
	}

	if specPath == "" {
		specPath = sourcePlugin.Name + "_to_" + destinationPlugin.Name + ".yaml"
	}
	fmt.Printf("Writing spec to %s...\n", bold.Sprint(specPath))
	var yamlSpec strings.Builder
	sourceConfig := configForSourcePlugin(sourcePlugin, sourceVersion)
	yamlSpec.WriteString(strings.ReplaceAll(sourceConfig, "DESTINATION_NAME", destinationPlugin.Name))
	yamlSpec.WriteString("\n---\n")
	yamlSpec.WriteString(configForDestinationPlugin(destinationPlugin, destinationVersion))

	if err := os.WriteFile(specPath, []byte(yamlSpec.String()), 0644); err != nil {
		return fmt.Errorf("failed to write spec file %w", err)
	}

	if user != nil {
		successful.Println("Sync spec file generated successfully!")
		fmt.Println()
		fmt.Println("Next steps:")
		fmt.Printf("1. Review the generated config file %s and make sure to fill in all authentication details. Learn more about the plugins configuration at:\n", bold.Sprint(specPath))
		fmt.Printf("   %s: %s\n", bold.Sprint(sourcePlugin.DisplayName), linkForPlugin(sourcePlugin))
		fmt.Printf("   %s: %s\n", bold.Sprint(destinationPlugin.DisplayName), linkForPlugin(destinationPlugin))
		fmt.Println("2. Run the following command to start the sync:")
		bold.Printf("cloudquery sync %s\n", specPath)
	} else {
		successful.Println("Sync spec file generated successfully!")
		fmt.Println()
		fmt.Println("Next steps:")
		fmt.Printf("1. Review the generated config file %s and make sure to fill in all authentication details. Learn more about the plugins configuration at:\n", bold.Sprint(specPath))
		fmt.Printf("   %s: %s\n", bold.Sprint(sourcePlugin.DisplayName), linkForPlugin(sourcePlugin))
		fmt.Printf("   %s: %s\n", bold.Sprint(destinationPlugin.DisplayName), linkForPlugin(destinationPlugin))
		fmt.Println("2. Run the following command to log in:")
		bold.Printf("cloudquery login\n")
		fmt.Println()
		fmt.Println("3. Run the following command to start the sync:")
		bold.Printf("cloudquery sync %s\n", specPath)
	}
	return nil
}
