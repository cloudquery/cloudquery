package cmd

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	cqplatform "github.com/cloudquery/cloudquery/cli/v6/internal/platform"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/glob"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	upgradeCheckShort = "Preview how upgrading a source plugin affects your destinations, without migrating or writing anything"
	upgradeCheckLong  = upgradeCheckShort + `

Exit codes:
  0  no action needed: no schema change, automatically migratable changes, or no output difference
  3  action needed: manual migration or rebuild, selected tables removed, or file schema or output changed
  4  unknown: a destination or source could not be assessed, or coverage is incomplete
  1  the check failed with an error

With several destinations, the highest of 0, 3 and 4 is used. The report is always printed first.`
	upgradeCheckExample = `# Check how upgrading the datadog source to v6.0.0 affects the destinations in config.yml
cloudquery upgrade check ./config.yml --source datadog --to v6.0.0

# Compare a source that runs from a local binary, gRPC server or Docker image with cloudquery/semgrep v3.1.0 from the CloudQuery Hub
cloudquery upgrade check ./config.yml --source semgrep --to v3.1.0 --to-path cloudquery/semgrep
`
	destinationNoProtocolV3Reason = "destination does not support plugin protocol v3"
	destinationNoAssessmentReason = "destination version does not support assessment"

	upgradeOutputText = "text"
	upgradeOutputJSON = "json"
)

func newCmdUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Plugin upgrade commands",
	}
	cmd.AddCommand(newCmdUpgradeCheck())
	return cmd
}

func newCmdUpgradeCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "check [files or directories]",
		Short:   upgradeCheckShort,
		Long:    upgradeCheckLong,
		Example: upgradeCheckExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    upgradeCheck,
	}
	cmd.Flags().String("source", "", "Name of the source to upgrade, as set in the configuration")
	cmd.Flags().String("to", "", "Source plugin version to upgrade to")
	cmd.Flags().String("to-path", "", "CloudQuery Hub path of the source plugin to upgrade to, as <team>/<name>. Only for sources from the local, grpc or docker registry. Default: cloudquery/<name reported by the configured source>")
	_ = cmd.MarkFlagRequired("source")
	_ = cmd.MarkFlagRequired("to")
	cmd.Flags().String("output", upgradeOutputText, "Output format. One of: text, json")
	cmd.Flags().String("license", "", "set offline license file")
	cmd.Flags().Bool("cq-columns-not-null", false, "Force CloudQuery internal columns to be NOT NULL. This feature is in Preview. Please provide feedback to help us improve it.")
	_ = cmd.Flags().MarkHidden("cq-columns-not-null")
	return cmd
}

func upgradeCheck(cmd *cobra.Command, args []string) error {
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}
	sourceName, err := cmd.Flags().GetString("source")
	if err != nil {
		return err
	}
	toVersion, err := cmd.Flags().GetString("to")
	if err != nil {
		return err
	}
	toPath, err := cmd.Flags().GetString("to-path")
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	if output != upgradeOutputText && output != upgradeOutputJSON {
		return fmt.Errorf("invalid output format %q. One of: %s, %s", output, upgradeOutputText, upgradeOutputJSON)
	}
	licenseFile, err := cmd.Flags().GetString("license")
	if err != nil {
		return err
	}
	cqColumnsNotNull, err := cmd.Flags().GetBool("cq-columns-not-null")
	if err != nil {
		return err
	}

	out := cmd.OutOrStdout()
	if output == upgradeOutputJSON {
		// Plugin downloads print progress to os.Stdout, which would corrupt the JSON report.
		stdout := os.Stdout
		os.Stdout = os.Stderr
		defer func() { os.Stdout = stdout }()
	}

	ctx := cmd.Context()
	loaded, err := loadUpgradeCheckSpecs(ctx, args, sourceName)
	if err != nil {
		return err
	}
	sourceSpec, destinationSpecs, transformersForDestination := loaded.source, loaded.destinations, loaded.transformersForDestination
	dlToken, teamName := loaded.downloadToken, loaded.teamName

	opts := []managedplugin.Option{
		managedplugin.WithLogger(log.Logger),
		managedplugin.WithAuthToken(dlToken),
		managedplugin.WithTeamName(teamName),
		managedplugin.WithLicenseFile(licenseFile),
	}
	if logConsole {
		opts = append(opts, managedplugin.WithNoProgress())
	}
	if cqDir != "" {
		opts = append(opts, managedplugin.WithDirectory(cqDir))
	}
	if disableSentry {
		opts = append(opts, managedplugin.WithNoSentry())
	}

	from, to, err := loadUpgradeSourceTables(ctx, *sourceSpec, upgradeTarget{Version: toVersion, Path: toPath}, managedplugin.NewClient, opts...)
	if err != nil {
		return err
	}
	var reports []upgradeReport
	exitCode := 0
	err = upgradeReports(ctx, *sourceSpec, destinationSpecs, transformersForDestination, cqColumnsNotNull, from, to, func(report upgradeReport) error {
		reports = append(reports, report)
		exitCode = max(exitCode, upgradeExitCode(report))
		if output == upgradeOutputText {
			return renderUpgradeReport(out, report)
		}
		return nil
	}, opts...)
	if err != nil {
		return err
	}
	if output == upgradeOutputJSON {
		if err := renderUpgradeReportsJSON(out, reports); err != nil {
			return err
		}
	}
	return upgradeExitError(cmd, exitCode)
}

func upgradeReports(ctx context.Context, sourceSpec specs.Source, destinationSpecs []*specs.Destination, transformersForDestination map[string][]*specs.Transformer, cqColumnsNotNull bool, from, to upgradeSourceTables, emit func(upgradeReport) error, opts ...managedplugin.Option) error {
	report := upgradeReport{
		SourceName:  sourceSpec.Name,
		FromVersion: from.Version,
		ToVersion:   to.Version,
		FromOrigin:  from.Origin,
		ToOrigin:    to.Origin,
		SourceGaps:  upgradeSourceGaps(sourceSpec.Name, from, to),
	}
	if from.UnknownReason != "" || to.UnknownReason != "" {
		report.SourceUnknown = true
		return emit(report)
	}

	selection, err := selectUpgradeTables(sourceSpec, from.Tables, to.Tables)
	if err != nil {
		return err
	}
	report.RemovedTables = selection.RemovedTables
	report.UnmatchedTablePatterns = selection.UnmatchedTablePatterns
	for _, destinationSpec := range destinationSpecs {
		tables, findings, err := assessDestination(ctx, sourceSpec, *destinationSpec, transformersForDestination[destinationSpec.Name], cqColumnsNotNull, selection.Pairs, opts...)
		if err != nil {
			return err
		}
		report.Destination = destinationSpec
		report.Findings = findings
		if report.Tables, err = upgradeTablesByName(tables); err != nil {
			return err
		}
		if err := emit(report); err != nil {
			return err
		}
	}
	return nil
}

func upgradeExitError(cmd *cobra.Command, exitCode int) error {
	if exitCode == 0 {
		return nil
	}
	cmd.SilenceErrors = true
	return &ExitCodeError{Code: exitCode}
}

type upgradeCheckSpecs struct {
	source                     *specs.Source
	destinations               []*specs.Destination
	transformersForDestination map[string][]*specs.Transformer
	downloadToken              string
	teamName                   string
}

func loadUpgradeCheckSpecs(ctx context.Context, args []string, sourceName string) (upgradeCheckSpecs, error) {
	specReader, err := specs.NewSpecReaderWithoutValidation(args)
	if err != nil {
		return upgradeCheckSpecs{}, fmt.Errorf("failed to load spec(s) from %s. Error: %w", strings.Join(args, ", "), err)
	}
	loaded := upgradeCheckSpecs{source: specReader.GetSourceByName(sourceName), transformersForDestination: make(map[string][]*specs.Transformer)}
	if loaded.source == nil {
		return upgradeCheckSpecs{}, fmt.Errorf("source %q not found in %s", sourceName, strings.Join(args, ", "))
	}

	loaded.downloadToken, loaded.teamName, err = cqplatform.DownloadAuth(ctx, log.Logger, upgradeAuthSources(*loaded.source), specReader.Destinations, specReader.Transformers)
	if err != nil {
		return upgradeCheckSpecs{}, err
	}
	cqplatform.PropagatePluginCredential(loaded.downloadToken)
	destinations, err := cqplatform.MaybeInjectDestination(ctx, log.Logger, loaded.downloadToken, loaded.teamName, specReader.Sources, specReader.Destinations)
	if err != nil {
		return upgradeCheckSpecs{}, err
	}
	if err := specReader.SetDestinationsAndValidate(destinations); err != nil {
		return upgradeCheckSpecs{}, fmt.Errorf("failed to load spec(s) from %s. Error: %w", strings.Join(args, ", "), err)
	}

	transformerSpecsByName := make(map[string]*specs.Transformer)
	for _, transformerSpec := range specReader.Transformers {
		transformerSpecsByName[transformerSpec.Name] = transformerSpec
	}
	for _, destinationSpec := range specReader.Destinations {
		if !slices.Contains(loaded.source.Destinations, destinationSpec.Name) {
			continue
		}
		loaded.destinations = append(loaded.destinations, destinationSpec)
		for _, transformerName := range destinationSpec.Transformers {
			loaded.transformersForDestination[destinationSpec.Name] = append(loaded.transformersForDestination[destinationSpec.Name], transformerSpecsByName[transformerName])
		}
	}
	return loaded, nil
}

func upgradeAuthSources(sourceSpec specs.Source) []*specs.Source {
	sources := []*specs.Source{&sourceSpec}
	if !upgradeVersionSelectsBinary(sourceSpec.Registry) {
		sources = append(sources, &specs.Source{Metadata: specs.Metadata{Name: sourceSpec.Name, Registry: specs.RegistryCloudQuery}})
	}
	return sources
}

func upgradeSourceGaps(sourceName string, from, to upgradeSourceTables) []string {
	fromLabel, toLabel := upgradeSourceLabels(from.Origin, from.Version, to.Origin, to.Version)
	var gaps []string
	for _, side := range []struct {
		label  string
		tables upgradeSourceTables
	}{{fromLabel, from}, {toLabel, to}} {
		if side.tables.UnknownReason != "" {
			gaps = append(gaps, fmt.Sprintf("%s %s: tables could not be listed: %s", sourceName, side.label, side.tables.UnknownReason))
		} else if side.tables.Connected {
			gaps = append(gaps, fmt.Sprintf("%s %s: tables were listed with a connection (metadata only, no rows read)", sourceName, side.label))
		}
	}
	return gaps
}

func managedPluginConfig(metadata specs.Metadata) managedplugin.Config {
	return managedplugin.Config{
		Name:       metadata.Name,
		Path:       metadata.Path,
		Version:    metadata.Version,
		Registry:   SpecRegistryToPlugin(metadata.Registry),
		DockerAuth: metadata.DockerRegistryAuthToken,
	}
}

func terminatePlugin(client *managedplugin.Client) {
	if err := client.Terminate(); err != nil {
		log.Warn().Err(err).Str("plugin", client.Name()).Msg("failed to terminate plugin")
	}
}

type upgradeSourceOrigin struct {
	Registry specs.Registry
	Path     string
}

func (o upgradeSourceOrigin) label(version string) string {
	if upgradeVersionSelectsBinary(o.Registry) {
		return o.Path + "@" + version
	}
	return fmt.Sprintf("(%s: %s, %s)", o.Registry, o.Path, version)
}

func upgradeSourceLabels(from upgradeSourceOrigin, fromVersion string, to upgradeSourceOrigin, toVersion string) (fromLabel, toLabel string) {
	if from == to {
		return fromVersion, toVersion
	}
	return from.label(fromVersion), to.label(toVersion)
}

type upgradeSourceTables struct {
	Origin        upgradeSourceOrigin
	Version       string
	PluginName    string
	Tables        schema.Tables
	Connected     bool
	UnknownReason string
}

type upgradeTarget struct {
	Version string
	Path    string
}

type newPluginClientFunc func(ctx context.Context, typ managedplugin.PluginType, config managedplugin.Config, opts ...managedplugin.Option) (*managedplugin.Client, error)

const upgradeUnknownVersion = "unknown"

func upgradeVersionSelectsBinary(registry specs.Registry) bool {
	return registry == specs.RegistryCloudQuery || registry == specs.RegistryGitHub
}

func loadUpgradeSourceTables(ctx context.Context, sourceSpec specs.Source, target upgradeTarget, newClient newPluginClientFunc, opts ...managedplugin.Option) (from, to upgradeSourceTables, err error) {
	if err := validateUpgradeTarget(sourceSpec, target); err != nil {
		return from, to, err
	}
	if err := registerOnce(); err != nil {
		return from, to, err
	}
	fromOrigin := upgradeSourceOrigin{Registry: sourceSpec.Registry, Path: sourceSpec.Path}
	from, err = loadSourceTables(ctx, newClient, sourceSpec, fromOrigin, managedPluginConfig(sourceSpec.Metadata), opts...)
	if err != nil {
		return from, to, err
	}
	toOrigin, toConfig, err := upgradeTargetConfig(sourceSpec, target, from.PluginName)
	if err != nil {
		return from, to, err
	}
	to, err = loadSourceTables(ctx, newClient, sourceSpec, toOrigin, toConfig, opts...)
	if err != nil && !upgradeVersionSelectsBinary(sourceSpec.Registry) && target.Path == "" {
		return from, to, fmt.Errorf("%w. The CloudQuery Hub path %s was inferred from the name the source reports. Use --to-path <team>/<name> to set a different path", err, toConfig.Path)
	}
	return from, to, err
}

func validateUpgradeTarget(sourceSpec specs.Source, target upgradeTarget) error {
	if target.Path == "" {
		return nil
	}
	if upgradeVersionSelectsBinary(sourceSpec.Registry) {
		return fmt.Errorf("--to-path is supported only for sources from the local, grpc or docker registry, source %s uses the %s registry", sourceSpec.Name, sourceSpec.Registry)
	}
	team, name, ok := strings.Cut(target.Path, "/")
	if !ok || team == "" || name == "" || strings.Contains(name, "/") {
		return fmt.Errorf("invalid --to-path %q. Use <team>/<name>, for example cloudquery/aws", target.Path)
	}
	return nil
}

func upgradeTargetConfig(sourceSpec specs.Source, target upgradeTarget, pluginName string) (upgradeSourceOrigin, managedplugin.Config, error) {
	config := managedPluginConfig(sourceSpec.Metadata)
	config.Version = target.Version
	if upgradeVersionSelectsBinary(sourceSpec.Registry) {
		return upgradeSourceOrigin{Registry: sourceSpec.Registry, Path: config.Path}, config, nil
	}
	config.Registry = managedplugin.RegistryCloudQuery
	config.DockerAuth = ""
	config.Path = target.Path
	if config.Path == "" {
		if pluginName == "" {
			return upgradeSourceOrigin{}, managedplugin.Config{}, fmt.Errorf("could not infer the CloudQuery Hub path to upgrade source %s to, because the source from the %s registry did not report its name. Use --to-path <team>/<name> to set the path, for example --to-path cloudquery/aws", sourceSpec.Name, sourceSpec.Registry)
		}
		config.Path = "cloudquery/" + pluginName
	}
	return upgradeSourceOrigin{Registry: specs.RegistryCloudQuery, Path: config.Path}, config, nil
}

func loadSourceTables(ctx context.Context, newClient newPluginClientFunc, sourceSpec specs.Source, origin upgradeSourceOrigin, config managedplugin.Config, opts ...managedplugin.Option) (upgradeSourceTables, error) {
	version := config.Version
	if !upgradeVersionSelectsBinary(origin.Registry) {
		version = upgradeUnknownVersion
	}
	client, err := newClient(ctx, managedplugin.PluginSource, config, opts...)
	if err != nil {
		return upgradeSourceTables{}, fmt.Errorf("failed to start source %s %s: %w", sourceSpec.Name, origin.label(version), err)
	}
	defer terminatePlugin(client)

	versions, err := client.Versions(ctx)
	if err != nil {
		return upgradeSourceTables{}, fmt.Errorf("failed to get protocol versions for %s %s: %w", sourceSpec.Name, origin.label(version), err)
	}
	if findMaxCommonVersion(versions, []int{3}) != 3 {
		return upgradeSourceTables{Origin: origin, Version: version, UnknownReason: "source does not support plugin protocol v3"}, nil
	}

	pluginClient := pluginPb.NewPluginClient(client.Conn)
	pluginName, reportedVersion := reportedPluginIdentity(ctx, pluginClient)
	if !upgradeVersionSelectsBinary(origin.Registry) && reportedVersion != "" {
		version = reportedVersion
	}
	result := listSourceTables(ctx, pluginClient, sourceSpec.Spec)
	result.Origin = origin
	result.Version = version
	result.PluginName = pluginName
	return result, nil
}

func reportedPluginIdentity(ctx context.Context, client pluginPb.PluginClient) (name, version string) {
	if resp, err := client.GetName(ctx, &pluginPb.GetName_Request{}); err == nil {
		name = resp.Name
	} else {
		log.Debug().Err(err).Msg("failed to get source plugin name")
	}
	if resp, err := client.GetVersion(ctx, &pluginPb.GetVersion_Request{}); err == nil {
		version = resp.Version
	} else {
		log.Debug().Err(err).Msg("failed to get source plugin version")
	}
	return name, version
}

func listSourceTables(ctx context.Context, client pluginPb.PluginClient, spec map[string]any) upgradeSourceTables {
	tables, err := initAndGetAllTables(ctx, client, spec, true)
	if err != nil {
		return upgradeSourceTables{UnknownReason: err.Error()}
	}
	if len(tables) > 0 {
		return upgradeSourceTables{Tables: tables}
	}

	tables, err = initAndGetAllTables(ctx, client, spec, false)
	result := upgradeSourceTables{Tables: tables, Connected: true}
	switch {
	case err != nil:
		result.UnknownReason = err.Error()
	case len(tables) == 0:
		result.UnknownReason = "source returned no tables with or without a connection"
	}
	return result
}

func initAndGetAllTables(ctx context.Context, client pluginPb.PluginClient, spec map[string]any, noConnection bool) (schema.Tables, error) {
	if err := initPlugin(ctx, client, spec, noConnection, invocationUUID.String()); err != nil {
		return nil, fmt.Errorf("failed to init source: %w", err)
	}
	return getTables(ctx, client, &pluginPb.GetTables_Request{Tables: []string{"*"}})
}

type upgradeTablePair struct {
	Name string
	From *schema.Table
	To   *schema.Table
}

type upgradeTableSelectedBy string

const (
	upgradeSelectedByName      upgradeTableSelectedBy = "name"
	upgradeSelectedByPattern   upgradeTableSelectedBy = "pattern"
	upgradeSelectedAsDependent upgradeTableSelectedBy = "dependent"
)

type upgradeRemovedTable struct {
	Name       string
	SelectedBy upgradeTableSelectedBy
}

type upgradeTableSelection struct {
	Pairs                  []upgradeTablePair
	RemovedTables          []upgradeRemovedTable
	UnmatchedTablePatterns []string
}

type upgradeTableSchemas struct {
	Name string
	From []byte
	To   []byte
}

func selectUpgradeTables(sourceSpec specs.Source, from, to schema.Tables) (upgradeTableSelection, error) {
	selectedFrom, err := selectSourceTables(sourceSpec, from)
	if err != nil {
		return upgradeTableSelection{}, fmt.Errorf("failed to select tables from current source version: %w", err)
	}
	selectedTo, err := selectSourceTables(sourceSpec, to)
	if err != nil {
		return upgradeTableSelection{}, fmt.Errorf("failed to select tables from target source version: %w", err)
	}

	var selection upgradeTableSelection
	for _, fromTable := range selectedFrom {
		toTable := selectedTo.Get(fromTable.Name)
		if toTable == nil {
			selection.RemovedTables = append(selection.RemovedTables, upgradeRemovedTable{Name: fromTable.Name, SelectedBy: upgradeSelectedBy(sourceSpec.Tables, fromTable.Name)})
			continue
		}
		selection.Pairs = append(selection.Pairs, upgradeTablePair{Name: fromTable.Name, From: fromTable, To: toTable})
	}
	for _, toTable := range selectedTo {
		if selectedFrom.Get(toTable.Name) == nil {
			selection.Pairs = append(selection.Pairs, upgradeTablePair{Name: toTable.Name, To: toTable})
		}
	}
	selection.UnmatchedTablePatterns = upgradeUnmatchedTablePatterns(sourceSpec.Tables, selection.RemovedTables, to)
	return selection, nil
}

func upgradeSelectedBy(patterns []string, tableName string) upgradeTableSelectedBy {
	switch {
	case slices.Contains(patterns, tableName):
		return upgradeSelectedByName
	case slices.ContainsFunc(patterns, func(pattern string) bool { return glob.Glob(pattern, tableName) }):
		return upgradeSelectedByPattern
	}
	return upgradeSelectedAsDependent
}

func upgradeUnmatchedTablePatterns(patterns []string, removedTables []upgradeRemovedTable, to schema.Tables) []string {
	var unmatched []string
	for _, pattern := range patterns {
		if !strings.Contains(pattern, glob.GLOB) {
			continue
		}
		matchesRemoved := slices.ContainsFunc(removedTables, func(table upgradeRemovedTable) bool { return glob.Glob(pattern, table.Name) })
		if matchesRemoved && !slices.ContainsFunc(to, matchesAnyTablePattern([]string{pattern})) {
			unmatched = append(unmatched, pattern)
		}
	}
	return unmatched
}

func selectSourceTables(sourceSpec specs.Source, tables schema.Tables) (schema.Tables, error) {
	topLevelTables, err := tables.UnflattenTables()
	if err != nil {
		return nil, err
	}
	return topLevelTables.FilterDfsFunc(
		matchesAnyTablePattern(sourceSpec.Tables),
		matchesAnyTablePattern(sourceSpec.SkipTables),
		*sourceSpec.SkipDependentTables,
	).FlattenTables(), nil
}

func matchesAnyTablePattern(patterns []string) func(*schema.Table) bool {
	return func(table *schema.Table) bool {
		return slices.ContainsFunc(patterns, func(pattern string) bool {
			return glob.Glob(pattern, table.Name)
		})
	}
}

func transformUpgradeTables(ctx context.Context, sourceSpec specs.Source, destinationSpec specs.Destination, transformerClients []pluginPb.PluginClient, cqColumnsNotNull bool, pairs []upgradeTablePair) ([]upgradeTableSchemas, error) {
	syncTime := time.Now().UTC()
	syncGroupId := ""
	if destinationSpec.SyncGroupId != "" {
		syncGroupId = destinationSpec.RenderedSyncGroupId(syncTime, invocationUUID.String())
	}
	recordTransformer := newDestinationRecordTransformer(destinationSpec, sourceSpec.Name, syncTime, syncGroupId, cqColumnsNotNull)

	transform := func(table *schema.Table) ([]byte, error) {
		if table == nil {
			return nil, nil
		}
		return transformSchemaForDestination(ctx, recordTransformer, transformerClients, table.ToArrowSchema())
	}

	transformed := make([]upgradeTableSchemas, len(pairs))
	for i, pair := range pairs {
		from, err := transform(pair.From)
		if err != nil {
			return nil, fmt.Errorf("failed to transform table %s from current source version for destination %s: %w", pair.Name, destinationSpec.Name, err)
		}
		to, err := transform(pair.To)
		if err != nil {
			return nil, fmt.Errorf("failed to transform table %s from target source version for destination %s: %w", pair.Name, destinationSpec.Name, err)
		}
		transformed[i] = upgradeTableSchemas{Name: pair.Name, From: from, To: to}
	}
	return transformed, nil
}

func assessDestination(ctx context.Context, sourceSpec specs.Source, destinationSpec specs.Destination, transformerSpecs []*specs.Transformer, cqColumnsNotNull bool, pairs []upgradeTablePair, opts ...managedplugin.Option) ([]upgradeTableSchemas, []*pluginPb.AssessTables_TableFinding, error) {
	if len(pairs) == 0 {
		return nil, nil, nil
	}

	transformerConfigs := make([]managedplugin.Config, len(transformerSpecs))
	for i, transformerSpec := range transformerSpecs {
		transformerConfigs[i] = managedPluginConfig(transformerSpec.Metadata)
	}
	managedTransformerClients, err := managedplugin.NewClients(ctx, managedplugin.PluginTransformer, transformerConfigs, opts...)
	defer func() {
		if err := managedTransformerClients.Terminate(); err != nil {
			log.Warn().Err(err).Str("destination", destinationSpec.Name).Msg("failed to terminate transformers")
		}
	}()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start transformers for destination %s: %w", destinationSpec.Name, err)
	}
	transformerClients := make([]pluginPb.PluginClient, len(transformerSpecs))
	for i, transformerSpec := range transformerSpecs {
		transformerClients[i] = pluginPb.NewPluginClient(managedTransformerClients[i].Conn)
		if err := initPlugin(ctx, transformerClients[i], transformerSpec.Spec, false, invocationUUID.String()); err != nil {
			return nil, nil, fmt.Errorf("failed to init transformer %s: %w", transformerSpec.Name, err)
		}
	}
	tables, err := transformUpgradeTables(ctx, sourceSpec, destinationSpec, transformerClients, cqColumnsNotNull, pairs)
	if err != nil {
		return nil, nil, err
	}

	client, err := managedplugin.NewClient(ctx, managedplugin.PluginDestination, managedPluginConfig(destinationSpec.Metadata), opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start destination %s: %w", destinationSpec.Name, err)
	}
	defer terminatePlugin(client)
	versions, err := client.Versions(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get protocol versions for destination %s: %w", destinationSpec.Name, err)
	}
	if !slices.Contains(versions, 3) {
		return tables, unknownTableFindings(tables, destinationNoProtocolV3Reason), nil
	}
	findings, err := assessTables(ctx, pluginPb.NewPluginClient(client.Conn), destinationSpec, tables)
	return tables, findings, err
}

func assessTables(ctx context.Context, client pluginPb.PluginClient, destinationSpec specs.Destination, tables []upgradeTableSchemas) ([]*pluginPb.AssessTables_TableFinding, error) {
	if err := initPlugin(ctx, client, destinationSpec.Spec, true, invocationUUID.String()); err != nil {
		return nil, fmt.Errorf("failed to init destination %s: %w", destinationSpec.Name, err)
	}
	req := &pluginPb.AssessTables_Request{MigrateForce: destinationSpec.MigrateMode == specs.MigrateModeForced}
	for _, table := range tables {
		req.Tables = append(req.Tables, &pluginPb.AssessTables_TablePair{OldTable: table.From, NewTable: table.To})
	}
	resp, err := client.AssessTables(ctx, req)
	if status.Code(err) == codes.Unimplemented {
		return unknownTableFindings(tables, destinationNoAssessmentReason), nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to assess tables for destination %s: %w", destinationSpec.Name, err)
	}
	return resp.Tables, nil
}

func upgradeDestinationTableName(table upgradeTableSchemas) string {
	for _, schemaBytes := range [][]byte{table.To, table.From} {
		if decoded, err := upgradeTableFromBytes(schemaBytes); err == nil && decoded != nil {
			return decoded.Name
		}
	}
	return table.Name
}

func unknownTableFindings(tables []upgradeTableSchemas, reason string) []*pluginPb.AssessTables_TableFinding {
	findings := make([]*pluginPb.AssessTables_TableFinding, len(tables))
	for i, table := range tables {
		findings[i] = &pluginPb.AssessTables_TableFinding{
			TableName:                upgradeDestinationTableName(table),
			Category:                 pluginPb.AssessTables_CATEGORY_UNKNOWN,
			IncompleteCoverageReason: reason,
		}
	}
	return findings
}
