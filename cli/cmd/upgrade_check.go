package cmd

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
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
