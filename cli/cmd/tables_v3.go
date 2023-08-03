package cmd

import (
	"context"

	"github.com/cloudquery/cloudquery/cli/internal/docs"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	pluginPb "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func tablesV3(ctx context.Context, sourceClient *managedplugin.Client, path string, format string) error {
	sourcePbClient := pluginPb.NewPluginClient(sourceClient.Conn)
	if _, err := sourcePbClient.Init(ctx, &pluginPb.Init_Request{
		NoConnection: true,
	}); err != nil {
		return err
	}
	getTablesResp, err := sourcePbClient.GetTables(ctx, &pluginPb.GetTables_Request{
		Tables: []string{"*"},
	})
	if err != nil {
		return err
	}
	name, err := sourcePbClient.GetName(ctx, &pluginPb.GetName_Request{})
	if err != nil {
		return err
	}

	schemas, err := pluginPb.NewSchemasFromBytes(getTablesResp.Tables)
	if err != nil {
		return err
	}
	tables, err := schema.NewTablesFromArrowSchemas(schemas)
	if err != nil {
		return err
	}

	g := docs.NewGenerator(name.Name, tables)
	f := docs.FormatMarkdown
	if format == "json" {
		f = docs.FormatJSON
	}
	return g.Generate(path, f)
}
