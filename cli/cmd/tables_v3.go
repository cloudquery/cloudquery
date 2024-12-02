package cmd

import (
	"context"
	"fmt"
	gosync "sync"

	"github.com/cloudquery/cloudquery/cli/v6/internal/docs"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

var registerOnce = gosync.OnceValue(types.RegisterAllExtensions)

func getTables(ctx context.Context, sourcePbClient pluginPb.PluginClient, req *pluginPb.GetTables_Request) (schema.Tables, error) {
	getTablesResp, err := sourcePbClient.GetTables(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get tables: %w", err)
	}
	schemas, err := pluginPb.NewSchemasFromBytes(getTablesResp.Tables)
	if err != nil {
		return nil, fmt.Errorf("failed to parse schemas: %w", err)
	}
	tables, err := schema.NewTablesFromArrowSchemas(schemas)
	if err != nil {
		return nil, fmt.Errorf("failed to convert schemas to tables: %w", err)
	}

	return tables, nil
}

func tablesV3(ctx context.Context, sourceClient *managedplugin.Client, sourceSpec *specs.Source, path, format, filter string) error {
	err := registerOnce()
	if err != nil {
		return err
	}
	sourcePbClient := pluginPb.NewPluginClient(sourceClient.Conn)
	if err := initPlugin(ctx, sourcePbClient, sourceSpec.Spec, true, invocationUUID.String()); err != nil {
		return fmt.Errorf("failed to init source: %w", err)
	}

	name, err := sourcePbClient.GetName(ctx, &pluginPb.GetName_Request{})
	if err != nil {
		return fmt.Errorf("failed to get source name: %w", err)
	}

	req := &pluginPb.GetTables_Request{
		Tables: []string{"*"},
	}

	if filter == "spec" {
		req.Tables = sourceSpec.Tables
		req.SkipTables = sourceSpec.SkipTables
		req.SkipDependentTables = *sourceSpec.SkipDependentTables
	}

	tables, err := getTables(ctx, sourcePbClient, req)
	if err != nil {
		return fmt.Errorf("failed to convert schemas to tables: %w", err)
	}
	topLevelTables, err := tables.UnflattenTables()
	if err != nil {
		return fmt.Errorf("failed to unflatten tables: %w", err)
	}

	g := docs.NewGenerator(name.Name, topLevelTables)
	f := docs.FormatMarkdown
	if format == "json" {
		f = docs.FormatJSON
	}
	return g.Generate(path, f)
}
