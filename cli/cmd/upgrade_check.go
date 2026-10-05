package cmd

import (
	"context"
	"fmt"
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
	upgradeCheckShort   = "Preview how upgrading a source plugin affects your destinations, without migrating or writing anything"
	upgradeCheckExample = `# Check how upgrading the datadog source to v6.0.0 affects the destinations in config.yml
cloudquery upgrade check ./config.yml --source datadog --to v6.0.0
`
	destinationNoProtocolV3Reason = "destination does not support plugin protocol v3"
	destinationNoAssessmentReason = "destination version does not support assessment"
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
		Long:    upgradeCheckShort,
		Example: upgradeCheckExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    upgradeCheck,
	}
	cmd.Flags().String("source", "", "Name of the source to upgrade, as set in the configuration")
	cmd.Flags().String("to", "", "Source plugin version to upgrade to")
	_ = cmd.MarkFlagRequired("source")
	_ = cmd.MarkFlagRequired("to")
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
	licenseFile, err := cmd.Flags().GetString("license")
	if err != nil {
		return err
	}
	cqColumnsNotNull, err := cmd.Flags().GetBool("cq-columns-not-null")
	if err != nil {
		return err
	}

	ctx := cmd.Context()
	specReader, err := specs.NewSpecReader(args)
	if err != nil {
		return fmt.Errorf("failed to load spec(s) from %s. Error: %w", strings.Join(args, ", "), err)
	}
	sourceSpec := specReader.GetSourceByName(sourceName)
	if sourceSpec == nil {
		return fmt.Errorf("source %q not found in %s", sourceName, strings.Join(args, ", "))
	}
	transformerSpecsByName := make(map[string]*specs.Transformer)
	for _, transformerSpec := range specReader.Transformers {
		transformerSpecsByName[transformerSpec.Name] = transformerSpec
	}
	var destinationSpecs []*specs.Destination
	var transformerSpecs []*specs.Transformer
	transformersForDestination := make(map[string][]*specs.Transformer)
	for _, destinationSpec := range specReader.Destinations {
		if !slices.Contains(sourceSpec.Destinations, destinationSpec.Name) {
			continue
		}
		destinationSpecs = append(destinationSpecs, destinationSpec)
		for _, transformerName := range destinationSpec.Transformers {
			transformersForDestination[destinationSpec.Name] = append(transformersForDestination[destinationSpec.Name], transformerSpecsByName[transformerName])
		}
		transformerSpecs = append(transformerSpecs, transformersForDestination[destinationSpec.Name]...)
	}

	dlToken, teamName, err := cqplatform.DownloadAuth(ctx, log.Logger, []*specs.Source{sourceSpec}, destinationSpecs, transformerSpecs)
	if err != nil {
		return err
	}
	cqplatform.PropagatePluginCredential(dlToken)

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

	from, to, err := loadUpgradeSourceTables(ctx, *sourceSpec, toVersion, opts...)
	if err != nil {
		return err
	}
	report := upgradeReport{
		SourceName:  sourceSpec.Name,
		FromVersion: from.Version,
		ToVersion:   to.Version,
		SourceGaps:  upgradeSourceGaps(sourceSpec.Name, from, to),
	}
	out := cmd.OutOrStdout()
	if from.UnknownReason != "" || to.UnknownReason != "" {
		report.SourceUnknown = true
		return renderUpgradeReport(out, report)
	}

	selection, err := selectUpgradeTables(*sourceSpec, from.Tables, to.Tables)
	if err != nil {
		return err
	}
	report.RemovedTables = selection.RemovedTables
	for _, destinationSpec := range destinationSpecs {
		tables, findings, err := assessDestination(ctx, *sourceSpec, *destinationSpec, transformersForDestination[destinationSpec.Name], cqColumnsNotNull, selection.Pairs, opts...)
		if err != nil {
			return err
		}
		report.Destination = destinationSpec
		report.Findings = findings
		if report.Tables, err = upgradeTablesByName(tables); err != nil {
			return err
		}
		if err := renderUpgradeReport(out, report); err != nil {
			return err
		}
	}
	return nil
}

func upgradeSourceGaps(sourceName string, versions ...upgradeSourceTables) []string {
	var gaps []string
	for _, version := range versions {
		if version.UnknownReason != "" {
			gaps = append(gaps, fmt.Sprintf("%s %s: tables could not be listed: %s", sourceName, version.Version, version.UnknownReason))
		} else if version.Connected {
			gaps = append(gaps, fmt.Sprintf("%s %s: tables were listed with a connection (metadata only, no rows read)", sourceName, version.Version))
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

type upgradeSourceTables struct {
	Version       string
	Tables        schema.Tables
	Connected     bool
	UnknownReason string
}

func loadUpgradeSourceTables(ctx context.Context, sourceSpec specs.Source, toVersion string, opts ...managedplugin.Option) (from, to upgradeSourceTables, err error) {
	if err := registerOnce(); err != nil {
		return from, to, err
	}
	from, err = loadSourceVersionTables(ctx, sourceSpec, sourceSpec.Version, opts...)
	if err != nil {
		return from, to, err
	}
	to, err = loadSourceVersionTables(ctx, sourceSpec, toVersion, opts...)
	return from, to, err
}

func loadSourceVersionTables(ctx context.Context, sourceSpec specs.Source, version string, opts ...managedplugin.Option) (upgradeSourceTables, error) {
	config := managedPluginConfig(sourceSpec.Metadata)
	config.Version = version
	client, err := managedplugin.NewClient(ctx, managedplugin.PluginSource, config, opts...)
	if err != nil {
		return upgradeSourceTables{}, fmt.Errorf("failed to start source %s@%s: %w", sourceSpec.Name, version, err)
	}
	defer terminatePlugin(client)

	versions, err := client.Versions(ctx)
	if err != nil {
		return upgradeSourceTables{}, fmt.Errorf("failed to get protocol versions for %s@%s: %w", sourceSpec.Name, version, err)
	}
	if findMaxCommonVersion(versions, []int{3}) != 3 {
		return upgradeSourceTables{Version: version, UnknownReason: "source does not support plugin protocol v3"}, nil
	}

	result := listSourceTables(ctx, pluginPb.NewPluginClient(client.Conn), sourceSpec.Spec)
	result.Version = version
	return result, nil
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

type upgradeTableSelection struct {
	Pairs         []upgradeTablePair
	RemovedTables []string
}

type upgradeTableSchemas struct {
	Name string
	From []byte
	To   []byte
}

func selectUpgradeTables(sourceSpec specs.Source, from, to schema.Tables) (upgradeTableSelection, error) {
	selectedFrom, err := selectSourceTables(sourceSpec, from)
	if err != nil {
		return upgradeTableSelection{}, fmt.Errorf("failed to select tables from source version %s: %w", sourceSpec.Version, err)
	}
	selectedTo, err := selectSourceTables(sourceSpec, to)
	if err != nil {
		return upgradeTableSelection{}, fmt.Errorf("failed to select tables from target source version: %w", err)
	}

	var selection upgradeTableSelection
	for _, fromTable := range selectedFrom {
		toTable := selectedTo.Get(fromTable.Name)
		if toTable == nil {
			selection.RemovedTables = append(selection.RemovedTables, fromTable.Name)
			continue
		}
		selection.Pairs = append(selection.Pairs, upgradeTablePair{Name: fromTable.Name, From: fromTable, To: toTable})
	}
	for _, toTable := range selectedTo {
		if selectedFrom.Get(toTable.Name) == nil {
			selection.Pairs = append(selection.Pairs, upgradeTablePair{Name: toTable.Name, To: toTable})
		}
	}
	return selection, nil
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
			return nil, fmt.Errorf("failed to transform table %s from source version %s for destination %s: %w", pair.Name, sourceSpec.Version, destinationSpec.Name, err)
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

func unknownTableFindings(tables []upgradeTableSchemas, reason string) []*pluginPb.AssessTables_TableFinding {
	findings := make([]*pluginPb.AssessTables_TableFinding, len(tables))
	for i, table := range tables {
		findings[i] = &pluginPb.AssessTables_TableFinding{
			TableName:                table.Name,
			Category:                 pluginPb.AssessTables_CATEGORY_UNKNOWN,
			CoverageIncomplete:       true,
			CoverageIncompleteReason: reason,
		}
	}
	return findings
}
