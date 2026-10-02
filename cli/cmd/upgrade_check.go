package cmd

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/glob"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog/log"
)

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
	config := managedplugin.Config{
		Name:       sourceSpec.Name,
		Path:       sourceSpec.Path,
		Version:    version,
		Registry:   SpecRegistryToPlugin(sourceSpec.Registry),
		DockerAuth: sourceSpec.DockerRegistryAuthToken,
	}
	client, err := managedplugin.NewClient(ctx, managedplugin.PluginSource, config, opts...)
	if err != nil {
		return upgradeSourceTables{}, fmt.Errorf("failed to start source %s@%s: %w", sourceSpec.Name, version, err)
	}
	defer func() {
		if err := client.Terminate(); err != nil {
			log.Warn().Err(err).Str("source", sourceSpec.Name).Str("version", version).Msg("failed to terminate source")
		}
	}()

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
