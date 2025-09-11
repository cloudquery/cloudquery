package aionboarding

import (
	"context"
	"strings"

	"github.com/cloudquery/cloudquery/cli/v6/aionboarding/openaiapi"
)

func (a *AIOnboarding) buildTools() map[string]openaiapi.Tool {
	return map[string]openaiapi.Tool{
		"list_all_valid_sources_and_destinations": {
			Name:        "list_all_valid_sources_and_destinations",
			Description: "List valid source and destination plugin names that can be used to create a spec file. Only these strings can be used as plugin names by other functions. Start by calling this function.",
			Parameters:  map[string]any{},
			Fn: func(ctx context.Context, args map[string]any) (string, error) {
				plugins, err := a.CloudAPI.ListAllPlugins(ctx)
				if err != nil {
					return "", err
				}
				return strings.Join(plugins, ", "), nil
			},
		},
		"get_source_or_destination_data": {
			Name:        "get_source_or_destination_data",
			Description: "Get the docs and latest version for a given source or destination plugin. Use this to know how to authenticate, configure and know the latest version of a given plugin. After this, call get_source_plugin_tables to know which tables are available to sync. Docs teach you what this plugin does, how it works and how to create the spec section of this plugin.",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"plugin_name": map[string]any{
						"type":        "string",
						"description": "The name of the source or destination plugin to get data for (use list_all_valid_sources_and_destinations function to get valid names)",
					},
					"plugin_kind": map[string]any{
						"type":        "string",
						"description": "The kind of plugin ('source' or 'destination')",
					},
				},
				"required": []string{"plugin_name", "plugin_kind"},
			},
			Fn: func(ctx context.Context, args map[string]any) (string, error) {
				pluginData, pluginLatestVersion, err := a.CloudAPI.GetPluginData(ctx, args["plugin_name"].(string), args["plugin_kind"].(string))
				if err != nil {
					return "", err
				}
				return pluginData + "\n\nPlugin latest version: " + pluginLatestVersion, nil
			},
		},
		"get_source_plugin_tables": {
			Name:        "get_source_plugin_tables",
			Description: "Get the tables available for a given plugin. Use this to know which tables are available to sync. There can be 1000s, so you can filter by RE2 regex. Use this to know what to put in the `tables: []` section of the source plugin spec. Note that you can use glob patterns in there, like `*` (or prefixing/suffixing with `*`) to include many/all tables.",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"plugin_name": map[string]any{
						"type":        "string",
						"description": "The name of the plugin to get tables for (use list_all_valid_sources_and_destinations function to get valid plugin names)",
					},
					"version": map[string]any{
						"type":        "string",
						"description": "The version of the plugin (use get_source_or_destination_data function to get latest version)",
					},
					"tables_regex": map[string]any{
						"type":        "string",
						"description": "RE2 regex pattern to filter tables; use * to include all tables",
					},
				},
				"required": []string{"plugin_name", "version"},
			},
			Fn: func(ctx context.Context, args map[string]any) (string, error) {
				tables, err := a.CloudAPI.GetSourcePluginTables(ctx, args["plugin_name"].(string), args["version"].(string), args["tables_regex"].(string))
				if err != nil {
					return "", err
				}
				return strings.Join(tables, ", "), nil
			},
		},
		"create_spec_file": {
			Name:        "create_spec_file",
			Description: "Create a spec file with the given filename and content. This function will write the provided content to the specified filename. This should be the last function called in the conversation.",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"filename_without_extension": map[string]any{
						"type":        "string",
						"description": "The filename without extension (i.e. no .yaml) for the spec file. Will overwrite existing. Will be created in the current directory.",
					},
					"content": map[string]any{
						"type":        "string",
						"description": "The content to write to the spec file",
					},
				},
				"required": []string{"path", "content"},
			},
			Fn: nil, // external
		},
	}
}
