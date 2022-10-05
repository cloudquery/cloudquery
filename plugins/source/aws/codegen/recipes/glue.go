package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GlueResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "classifiers",
			Struct:     &types.Classifier{},
			SkipFields: []string{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSRegion`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "name",
					Type:     schema.TypeString,
					Resolver: `resolveGlueClassifierName`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService: "connections",
			Struct:     &types.Connection{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveGlueConnectionArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "crawlers",
			Struct:     &types.Crawler{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveGlueCrawlerArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveGlueCrawlerTags`,
					},
				}...),
		},
		{
			SubService: "databases",
			Struct:     &types.Database{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveGlueDatabaseArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveGlueDatabaseTags`,
					},
				}...),
			Relations: []string{
				"DatabaseTables()",
			},
		},
		{
			SubService: "database_tables",
			Struct:     &types.Table{},
			SkipFields: []string{"Name"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "database_arn",
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
				"DatabaseTableIndexes()",
			},
		},
		{
			SubService: "database_table_indexes",
			Struct:     &types.PartitionIndexDescriptor{},
			SkipFields: []string{"IndexName"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "database_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("database_arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "database_table_name",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("name")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "index_name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("IndexName")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "datacatalog_encryption_settings",
			Struct:     &types.DataCatalogEncryptionSettings{},
			SkipFields: []string{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSRegion",
				},
			},
		},
		{
			SubService: "dev_endpoints",
			Struct:     &types.DevEndpoint{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveGlueDevEndpointArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveGlueDevEndpointTags`,
					},
				}...),
		},
		{
			SubService: "jobs",
			Struct:     &types.Job{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveGlueJobArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveGlueJobTags`,
					},
				}...),
			Relations: []string{
				"JobRuns()",
			},
		},
		{
			SubService: "job_runs",
			Struct:     &types.JobRun{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "job_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "ml_transforms",
			Struct:     &types.MLTransform{},
			SkipFields: []string{"Schema"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveGlueMlTransformArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveGlueMlTransformTags`,
					},
					{
						Name:     "schema",
						Type:     schema.TypeJSON,
						Resolver: `resolveMlTransformsSchema`,
					},
				}...),
			Relations: []string{
				"MlTransformTaskRuns()",
			},
		},
		{
			SubService: "ml_transform_task_runs",
			Struct:     &types.TaskRun{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "ml_transform_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "registries",
			Struct:     &types.RegistryListItem{},
			SkipFields: []string{"RegistryArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveGlueRegistryTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("RegistryArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"RegistrySchemas()",
			},
		},
		{
			SubService:          "registry_schemas",
			Struct:              &glue.GetSchemaOutput{},
			SkipFields:          []string{"SchemaArn"},
			PreResourceResolver: "getRegistrySchema",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("SchemaArn")`,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveGlueRegistrySchemaTags`,
					},
				}...),
			Relations: []string{
				"RegistrySchemaVersions()",
			},
		},
		{
			SubService:          "registry_schema_versions",
			Struct:              &glue.GetSchemaVersionOutput{},
			SkipFields:          []string{},
			PreResourceResolver: "getRegistrySchemaVersion",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "registry_schema_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "metadata",
						Type:     schema.TypeJSON,
						Resolver: `resolveGlueRegistrySchemaVersionMetadata`,
					},
				}...),
		},
		{
			SubService: "security_configurations",
			Struct:     &types.SecurityConfiguration{},
			SkipFields: []string{"Name"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSRegion`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:    "name",
					Type:    schema.TypeString,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService:          "triggers",
			Struct:              &types.Trigger{},
			SkipFields:          []string{},
			PreResourceResolver: "getTrigger",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveGlueTriggerArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveGlueTriggerTags`,
					},
				}...),
		},
		{
			SubService:          "workflows",
			Struct:              &types.Workflow{},
			SkipFields:          []string{},
			PreResourceResolver: "getWorkflow",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveGlueWorkflowArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveGlueWorkflowTags`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "glue"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("glue")`
	}
	return resources
}
