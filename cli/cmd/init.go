package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"regexp"
	"sort"
	"strings"

	cqapi "github.com/cloudquery/cloudquery-api-go"
	cqauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/analytics"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

const (
	initShort   = `Generate a configuration file for a sync`
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
)

func newCmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Short:   initShort,
		Long:    initShort,
		Example: initExample,
		Args:    cobra.ExactArgs(0),
		RunE:    initCmd,
	}
	cmd.Flags().String("source", "", "Source plugin name or path")
	cmd.Flags().String("destination", "", "Destination plugin name or path")
	cmd.Flags().String("spec-path", "", "Output spec file path")
	cmd.Flags().Bool("yes", false, "Accept all defaults")
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

func parseFlags(cmd *cobra.Command) (source, destination, specPath string, acceptDefaults bool, allErrors error) {
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
	return source, destination, specPath, acceptDefaults, allErrors
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

func pluginsSorter(plugins []cqapi.ListPlugin, prioritySlice []string) func(a, b int) bool {
	return func(a, b int) bool {
		indexOfA := lo.IndexOf(prioritySlice, plugins[a].Name)
		indexOfB := lo.IndexOf(prioritySlice, plugins[b].Name)
		if indexOfA == -1 && indexOfB != -1 {
			return false
		}
		if indexOfA != -1 && indexOfB == -1 {
			return true
		}
		if indexOfA == -1 && indexOfB == -1 {
			return plugins[a].Name < plugins[b].Name
		}
		return indexOfA < indexOfB
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

func selectSource(allPlugins []cqapi.ListPlugin, acceptDefaults bool) (string, error) {
	officialSources := lo.Filter(allPlugins, officialReleasedPluginsByKind(cqapi.PluginKindSource))
	sort.SliceStable(officialSources, pluginsSorter(officialSources, sourcesOrder))
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
	sort.SliceStable(officialDestinations, pluginsSorter(officialDestinations, destinationsOrder))
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
	return link.Sprintf("https://hub.cloudquery.io/plugins/%s/%s/%s", plugin.Kind, plugin.TeamName, plugin.Name)
}

func initCmd(cmd *cobra.Command, args []string) (initCommandError error) {
	ctx := cmd.Context()
	source, destination, specPath, acceptDefaults, err := parseFlags(cmd)
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

	apiClient, err := api.NewAnonymousClient()
	if err != nil {
		return err
	}
	if user != nil {
		apiClient, err = api.NewClient(token.Value)
		if err != nil {
			return err
		}
	}

	fmt.Println("Fetching plugins...")
	allPlugins, err := api.ListAllPlugins(apiClient)
	if err != nil {
		return err
	}

	var notFoundPluginsErrors error
	if source != "" {
		sourcePluginFilter := pluginFilter(source, cqapi.PluginKindSource)
		sourceFound := lo.SomeBy(allPlugins, sourcePluginFilter)
		if !sourceFound {
			notFoundPluginsErrors = errors.Join(notFoundPluginsErrors, fmt.Errorf("source plugin %q not found", source))
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
		source, err = selectSource(allPlugins, acceptDefaults)
		if err != nil {
			return err
		}
		source, _ = normalizePluginPath(source)
	}
	_, sourceIndex, _ := lo.FindIndexOf(allPlugins, pluginFilter(source, cqapi.PluginKindSource))

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
