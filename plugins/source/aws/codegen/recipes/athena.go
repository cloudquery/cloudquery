package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AthenaResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "data_catalogs",
			Struct:     &types.DataCatalog{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveAthenaDataCatalogArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveAthenaDataCatalogTags`,
					},
				}...),
			Relations: []string{
				"DataCatalogDatabases()",
			},
		},
		{
			SubService: "data_catalog_databases",
			Struct:     &types.Database{},
			SkipFields: []string{"Name"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "data_catalog_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Name")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"DataCatalogDatabaseTables()",
			},
		},
		{
			SubService: "data_catalog_database_tables",
			Struct:     &types.TableMetadata{},
			SkipFields: []string{"Name"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "data_catalog_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("data_catalog_arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "data_catalog_database_name",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("name")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Name")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "work_groups",
			Struct:     &types.WorkGroup{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveAthenaWorkGroupArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveAthenaWorkGroupTags`,
					},
				}...),
			Relations: []string{
				"WorkGroupPreparedStatements()",
				"WorkGroupQueryExecutions()",
				"WorkGroupNamedQueries()",
			},
		},
		{
			SubService: "work_group_prepared_statements",
			Struct:     &types.PreparedStatement{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "work_group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "work_group_query_executions",
			Struct:     &types.QueryExecution{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "work_group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "work_group_named_queries",
			Struct:     &types.NamedQuery{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "work_group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "athena"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("athena")`
	}
	return resources
}
