package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func QuickSightResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "analyses",
			Struct:              &types.Analysis{},
			SkipFields:          []string{"Arn"},
			PreResourceResolver: "getAnalysis",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveTags()`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "dashboards",
			Struct:     types.DashboardSummary{},
			SkipFields: []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveTags()`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "data_sets",
			Struct:     types.DataSetSummary{},
			SkipFields: []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveTags()`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"Ingestions()",
			},
		},
		{
			SubService: "data_sources",
			Struct:     types.DataSource{},
			SkipFields: []string{"Arn", "AlternateDataSourceParameters"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveTags()`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					//{
					//	Name:          "alternate_data_source_parameters",
					//	Type:          schema.TypeJSON,
					//	Resolver:      `schema.PathResolver("AlternateDataSourceParameters")`,
					//	IgnoreInTests: true,
					//},
				}...),
		},
		{
			SubService:          "folders",
			Struct:              types.Folder{},
			SkipFields:          []string{"Arn"},
			PreResourceResolver: "getFolder",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveTags()`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "groups",
			Struct:     types.Group{},
			SkipFields: []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveTags()`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"GroupMembers()",
			},
		},
		{
			SubService: "group_members",
			Struct:     types.GroupMember{},
			SkipFields: []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "user_arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "ingestions",
			Struct:     types.Ingestion{},
			SkipFields: []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveTags()`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "data_set_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "templates",
			Struct:     types.TemplateSummary{},
			SkipFields: []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveTags()`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "users",
			Struct:     types.User{},
			SkipFields: []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveTags()`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "quicksight"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("quicksight")`
	}
	return resources
}
