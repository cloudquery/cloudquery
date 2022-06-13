package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

const MAX_GOROUTINES = 10

//go:generate cq-gen --resource data_catalogs --config gen.hcl --output .
func DataCatalogs() *schema.Table {
	return &schema.Table{
		Name:         "aws_athena_data_catalogs",
		Description:  "Contains information about a data catalog in an Amazon Web Services account",
		Resolver:     fetchAthenaDataCatalogs,
		Multiplex:    client.ServiceAccountRegionMultiplexer("athena"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "arn",
				Description: "ARN of the resource.",
				Type:        schema.TypeString,
				Resolver:    ResolveAthenaDataCatalogArn,
			},
			{
				Name:          "tags",
				Type:          schema.TypeJSON,
				Resolver:      ResolveAthenaDataCatalogTags,
				IgnoreInTests: true,
			},
			{
				Name:        "name",
				Description: "The name of the data catalog",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of data catalog to create: LAMBDA for a federated catalog, HIVE for an external hive metastore, or GLUE for an Glue Data Catalog",
				Type:        schema.TypeString,
			},
			{
				Name:          "description",
				Description:   "An optional description of the data catalog",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "parameters",
				Description:   "Specifies the Lambda function or functions to use for the data catalog",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_athena_data_catalog_databases",
				Description: "Contains metadata information for a database in a data catalog",
				Resolver:    fetchAthenaDataCatalogDatabases,
				Columns: []schema.Column{
					{
						Name:        "data_catalog_cq_id",
						Description: "Unique CloudQuery ID of aws_athena_data_catalogs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name of the database",
						Type:        schema.TypeString,
					},
					{
						Name:          "description",
						Description:   "An optional description of the database",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "parameters",
						Description:   "A set of custom key/value pairs",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_athena_data_catalog_database_tables",
						Description: "Contains metadata for a table",
						Resolver:    fetchAthenaDataCatalogDatabaseTables,
						Columns: []schema.Column{
							{
								Name:        "data_catalog_database_cq_id",
								Description: "Unique CloudQuery ID of aws_athena_data_catalog_databases table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The name of the table",
								Type:        schema.TypeString,
							},
							{
								Name:        "create_time",
								Description: "The time that the table was created",
								Type:        schema.TypeTimestamp,
							},
							{
								Name:          "last_access_time",
								Description:   "The last time the table was accessed",
								Type:          schema.TypeTimestamp,
								IgnoreInTests: true,
							},
							{
								Name:        "parameters",
								Description: "A set of custom key/value pairs for table properties",
								Type:        schema.TypeJSON,
							},
							{
								Name:          "table_type",
								Description:   "The type of table",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
						},
						Relations: []*schema.Table{
							{
								Name:          "aws_athena_data_catalog_database_table_columns",
								Description:   "Contains metadata for a column in a table",
								Resolver:      fetchAthenaDataCatalogDatabaseTableColumns,
								IgnoreInTests: true,
								Columns: []schema.Column{
									{
										Name:        "data_catalog_database_table_cq_id",
										Description: "Unique CloudQuery ID of aws_athena_data_catalog_database_tables table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "name",
										Description: "The name of the column",
										Type:        schema.TypeString,
									},
									{
										Name:        "comment",
										Description: "Optional information about the column",
										Type:        schema.TypeString,
									},
									{
										Name:        "type",
										Description: "The data type of the column",
										Type:        schema.TypeString,
									},
								},
							},
							{
								Name:          "aws_athena_data_catalog_database_table_partition_keys",
								Description:   "Contains metadata for a column in a table",
								Resolver:      fetchAthenaDataCatalogDatabaseTablePartitionKeys,
								IgnoreInTests: true,
								Columns: []schema.Column{
									{
										Name:        "data_catalog_database_table_cq_id",
										Description: "Unique CloudQuery ID of aws_athena_data_catalog_database_tables table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "name",
										Description: "The name of the column",
										Type:        schema.TypeString,
									},
									{
										Name:        "comment",
										Description: "Optional information about the column",
										Type:        schema.TypeString,
									},
									{
										Name:        "type",
										Description: "The data type of the column",
										Type:        schema.TypeString,
									},
								},
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

func fetchAthenaDataCatalogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	input := athena.ListDataCatalogsInput{}
	var sem = semaphore.NewWeighted(int64(MAX_GOROUTINES))
	for {
		response, err := svc.ListDataCatalogs(ctx, &input, func(options *athena.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		errs, ctx := errgroup.WithContext(ctx)
		for _, d := range response.DataCatalogsSummary {
			if err := sem.Acquire(ctx, 1); err != nil {
				return diag.WrapError(err)
			}
			func(summary types.DataCatalogSummary) {
				errs.Go(func() error {
					defer sem.Release(1)
					return fetchDataCatalog(ctx, res, c, summary)
				})
			}(d)
		}
		err = errs.Wait()
		if err != nil {
			return diag.WrapError(err)
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func ResolveAthenaDataCatalogArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	dc := resource.Item.(types.DataCatalog)
	return diag.WrapError(resource.Set(c.Name, createDataCatalogArn(cl, *dc.Name)))
}
func ResolveAthenaDataCatalogTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	dc := resource.Item.(types.DataCatalog)
	arn := createDataCatalogArn(cl, *dc.Name)
	params := athena.ListTagsForResourceInput{ResourceARN: &arn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return diag.WrapError(err)
		}
		client.TagsIntoMap(result.Tags, tags)
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}
func fetchAthenaDataCatalogDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	input := athena.ListDatabasesInput{
		CatalogName: parent.Item.(types.DataCatalog).Name,
	}
	for {
		response, err := svc.ListDatabases(ctx, &input, func(options *athena.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.DatabaseList

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func fetchAthenaDataCatalogDatabaseTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	input := athena.ListTableMetadataInput{
		CatalogName:  parent.Parent.Item.(types.DataCatalog).Name,
		DatabaseName: parent.Item.(types.Database).Name,
	}
	for {
		response, err := svc.ListTableMetadata(ctx, &input, func(options *athena.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.TableMetadataList

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func fetchAthenaDataCatalogDatabaseTableColumns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	res <- parent.Item.(types.TableMetadata).Columns
	return nil
}
func fetchAthenaDataCatalogDatabaseTablePartitionKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	res <- parent.Item.(types.TableMetadata).PartitionKeys
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func fetchDataCatalog(ctx context.Context, res chan<- interface{}, c *client.Client, catalogSummary types.DataCatalogSummary) error {
	svc := c.Services().Athena
	dc, err := svc.GetDataCatalog(ctx, &athena.GetDataCatalogInput{
		Name: catalogSummary.CatalogName,
	}, func(options *athena.Options) {
		options.Region = c.Region
	})
	if err != nil {
		// retrieving of default data catalog (AwsDataCatalog) returns "not found error" but it exists and its
		// relations can be fetched by its name
		if *catalogSummary.CatalogName == "AwsDataCatalog" {
			res <- types.DataCatalog{Name: catalogSummary.CatalogName, Type: catalogSummary.Type}
			return nil
		}
		if c.IsNotFoundError(err) {
			return nil
		}
		return diag.WrapError(err)
	}
	res <- *dc.DataCatalog
	return nil
}

func createDataCatalogArn(cl *client.Client, catalogName string) string {
	return cl.ARN(client.Athena, "datacatalog", catalogName)
}
