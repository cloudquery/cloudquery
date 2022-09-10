package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Databases() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_databases",
		Description: "The Database object represents a logical grouping of tables that might reside in a Hive metastore or an RDBMS",
		Resolver:    fetchGlueDatabases,
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the workflow.",
				Type:            schema.TypeString,
				Resolver:        resolveGlueDatabaseArn,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "tags",
				Description: "The collection of tags associated with the database",
				Type:        schema.TypeJSON,
				Resolver:    resolveGlueDatabaseTags,
			},
			{
				Name:        "name",
				Description: "The name of the database",
				Type:        schema.TypeString,
			},
			{
				Name:        "catalog_id",
				Description: "The ID of the Data Catalog in which the database resides",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_table_default_permissions",
				Description: "Creates a set of default permissions on the table for principals",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "create_time",
				Description: "The time at which the metadata database was created in the catalog",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "A description of the database",
				Type:        schema.TypeString,
			},
			{
				Name:        "location_uri",
				Description: "The location of the database (for example, an HDFS path)",
				Type:        schema.TypeString,
			},
			{
				Name:        "parameters",
				Description: "These key-value pairs define parameters and properties of the database",
				Type:        schema.TypeJSON,
			},
			{
				Name:     "target_database",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TargetDatabase"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_glue_database_tables",
				Description: "Represents a collection of related data organized in columns and rows",
				Resolver:    fetchGlueDatabaseTables,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "parameters",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Parameters"),
					},
					{
						Name:     "storage_descriptor",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("StorageDescriptor"),
					},
					{
						Name:        "name",
						Description: "The table name",
						Type:        schema.TypeString,
					},
					{
						Name:        "catalog_id",
						Description: "The ID of the Data Catalog in which the table resides",
						Type:        schema.TypeString,
					},
					{
						Name:        "create_time",
						Description: "The time when the table definition was created in the Data Catalog",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "created_by",
						Description: "The person or entity who created the table",
						Type:        schema.TypeString,
					},
					{
						Name:        "database_name",
						Description: "The name of the database where the table metadata resides",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "A description of the table",
						Type:        schema.TypeString,
					},
					{
						Name:        "is_registered_with_lake_formation",
						Description: "Indicates whether the table has been registered with Lake Formation",
						Type:        schema.TypeBool,
					},
					{
						Name:        "last_access_time",
						Description: "The last time that the table was accessed",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "last_analyzed_time",
						Description: "The last time that column statistics were computed for this table",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "owner",
						Description: "The owner of the table",
						Type:        schema.TypeString,
					},
					{
						Name:        "retention",
						Description: "The retention time for this table",
						Type:        schema.TypeInt,
					},
					{
						Name:        "table_type",
						Description: "The type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc)",
						Type:        schema.TypeString,
					},
					{
						Name:     "target_table",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("TargetTable"),
					},
					{
						Name:        "update_time",
						Description: "The last time that the table was updated",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "version_id",
						Description: "The ID of the table version",
						Type:        schema.TypeString,
					},
					{
						Name:        "view_expanded_text",
						Description: "If the table is a view, the expanded text of the view; otherwise null",
						Type:        schema.TypeString,
					},
					{
						Name:        "view_original_text",
						Description: "If the table is a view, the original text of the view; otherwise null",
						Type:        schema.TypeString,
					},
					{
						Name:        "partition_keys",
						Description: "Partition keys",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("PartitionKeys"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_glue_database_table_indexes",
						Description: "A descriptor for a partition index in a table",
						Resolver:    fetchGlueDatabaseTableIndexes,
						Columns: []schema.Column{
							{
								Name:        "database_table_cq_id",
								Description: "Unique CloudQuery ID of aws_glue_database_tables table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "index_name",
								Description: "The name of the partition index",
								Type:        schema.TypeString,
							},
							{
								Name:        "index_status",
								Description: "The status of the partition index",
								Type:        schema.TypeString,
							},
							{
								Name:        "keys",
								Description: "A list of one or more keys, as KeySchemaElement structures, for the partition index",
								Type:        schema.TypeJSON,
								Resolver:    client.SliceJsonResolver("Keys", "Name", "Type"),
							},
							{
								Name:        "backfill_errors",
								Description: "A list of errors that can occur when registering partition indexes for an existing table",
								Type:        schema.TypeJSON,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGlueDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetDatabasesInput{}
	for {
		result, err := svc.GetDatabases(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.DatabaseList
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveGlueDatabaseArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	arn := aws.String(databaseARN(cl, aws.ToString(resource.Item.(types.Database).Name)))
	return resource.Set(c.Name, arn)
}
func resolveGlueDatabaseTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetTagsInput{
		ResourceArn: aws.String(databaseARN(cl, aws.ToString(resource.Item.(types.Database).Name))),
	}

	response, err := svc.GetTags(ctx, &input)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Tags)
}
func fetchGlueDatabaseTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Database)
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetTablesInput{
		DatabaseName: r.Name,
	}
	for {
		result, err := svc.GetTables(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.TableList
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func fetchGlueDatabaseTableIndexes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	d := parent.Parent.Item.(types.Database)
	t := parent.Item.(types.Table)
	input := glue.GetPartitionIndexesInput{DatabaseName: d.Name, CatalogId: d.CatalogId, TableName: t.Name}
	for {
		result, err := svc.GetPartitionIndexes(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.PartitionIndexDescriptorList
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func databaseARN(cl *client.Client, name string) string {
	return cl.ARN(client.GlueService, "database", name)
}
