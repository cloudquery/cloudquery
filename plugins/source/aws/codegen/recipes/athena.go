package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AthenaResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "data_catalogs",
			Description:         "https://docs.aws.amazon.com/athena/latest/APIReference/API_DataCatalog.html",
			Struct:              &types.DataCatalog{},
			SkipFields:          []string{},
			PreResourceResolver: "getDataCatalog",
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
			SubService:  "data_catalog_databases",
			Struct:      &types.Database{},
			Description: "https://docs.aws.amazon.com/athena/latest/APIReference/API_Database.html",
			SkipFields:  []string{"Name"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "data_catalog_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
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
			SubService:  "data_catalog_database_tables",
			Struct:      &types.TableMetadata{},
			Description: "https://docs.aws.amazon.com/athena/latest/APIReference/API_TableMetadata.html",
			SkipFields:  []string{"Name"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "data_catalog_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("data_catalog_arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "data_catalog_database_name",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("name")`,
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
			SubService:          "work_groups",
			Description:         "https://docs.aws.amazon.com/athena/latest/APIReference/API_WorkGroup.html",
			Struct:              &types.WorkGroup{},
			SkipFields:          []string{},
			PreResourceResolver: "getWorkGroup",
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
			SubService:          "work_group_prepared_statements",
			Struct:              &types.PreparedStatement{},
			Description:         "https://docs.aws.amazon.com/athena/latest/APIReference/API_PreparedStatement.html",
			PreResourceResolver: "getWorkGroupPreparedStatement",
			SkipFields:          []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "work_group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:          "work_group_query_executions",
			Struct:              &types.QueryExecution{},
			Description:         "https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryExecution.html",
			PreResourceResolver: "getWorkGroupQueryExecution",
			SkipFields:          []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "work_group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:          "work_group_named_queries",
			Struct:              &types.NamedQuery{},
			Description:         "https://docs.aws.amazon.com/athena/latest/APIReference/API_NamedQuery.html",
			PreResourceResolver: "getWorkGroupNamedQuery",
			SkipFields:          []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "work_group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
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
