package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource crawlers --config crawlers.hcl --output .
func Crawlers() *schema.Table {
	return &schema.Table{
		Name:         "aws_glue_crawlers",
		Description:  "Specifies a crawler program that examines a data source and uses classifiers to try to determine its schema",
		Resolver:     fetchGlueCrawlers,
		Multiplex:    client.ServiceAccountRegionMultiplexer("glue"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "arn",
				Description: "ARN of the resource.",
				Type:        schema.TypeString,
				Resolver:    resolveGlueCrawlerArn,
			},
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueCrawlerTags,
			},
			{
				Name:        "classifiers",
				Description: "A list of UTF-8 strings that specify the custom classifiers that are associated with the crawler",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "configuration",
				Description: "Crawler configuration information",
				Type:        schema.TypeString,
			},
			{
				Name:        "crawl_elapsed_time",
				Description: "If the crawler is running, contains the total time elapsed since the last crawl began",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "crawler_security_configuration",
				Description: "The name of the SecurityConfiguration structure to be used by this crawler",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_time",
				Description: "The time that the crawler was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "database_name",
				Description: "The name of the database in which the crawler's output is stored",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "A description of the crawler",
				Type:        schema.TypeString,
			},
			{
				Name:        "lake_formation_configuration_account_id",
				Description: "Required for cross account crawls",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LakeFormationConfiguration.AccountId"),
			},
			{
				Name:        "lake_formation_configuration_use_lake_formation_credentials",
				Description: "Specifies whether to use Lake Formation credentials for the crawler instead of the IAM role credentials",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("LakeFormationConfiguration.UseLakeFormationCredentials"),
			},
			{
				Name:        "last_crawl_error_message",
				Description: "If an error occurred, the error information about the last crawl",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LastCrawl.ErrorMessage"),
			},
			{
				Name:        "last_crawl_log_group",
				Description: "The log group for the last crawl",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LastCrawl.LogGroup"),
			},
			{
				Name:        "last_crawl_log_stream",
				Description: "The log stream for the last crawl",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LastCrawl.LogStream"),
			},
			{
				Name:        "last_crawl_message_prefix",
				Description: "The prefix for a message about this crawl",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LastCrawl.MessagePrefix"),
			},
			{
				Name:        "last_crawl_start_time",
				Description: "The time at which the crawl started",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("LastCrawl.StartTime"),
			},
			{
				Name:        "last_crawl_status",
				Description: "Status of the last crawl",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LastCrawl.Status"),
			},
			{
				Name:        "last_updated",
				Description: "The time that the crawler was last updated",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "lineage_configuration_crawler_lineage_settings",
				Description: "Specifies whether data lineage is enabled for the crawler",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LineageConfiguration.CrawlerLineageSettings"),
			},
			{
				Name:        "name",
				Description: "The name of the crawler",
				Type:        schema.TypeString,
			},
			{
				Name:        "recrawl_behavior",
				Description: "Specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RecrawlPolicy.RecrawlBehavior"),
			},
			{
				Name:        "role",
				Description: "The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data",
				Type:        schema.TypeString,
			},
			{
				Name:        "schedule_expression",
				Description: "A cron expression used to specify the schedule (see Time-Based Schedules for Jobs and Crawlers (https://docsawsamazoncom/glue/latest/dg/monitor-data-warehouse-schedulehtml) For example, to run something every day at 12:15 UTC, you would specify: cron(15 12 * * ? *)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Schedule.ScheduleExpression"),
			},
			{
				Name:        "schedule_state",
				Description: "The state of the schedule",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Schedule.State"),
			},
			{
				Name:        "schema_change_policy_delete_behavior",
				Description: "The deletion behavior when the crawler finds a deleted object",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SchemaChangePolicy.DeleteBehavior"),
			},
			{
				Name:        "schema_change_policy_update_behavior",
				Description: "The update behavior when the crawler finds a changed schema",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SchemaChangePolicy.UpdateBehavior"),
			},
			{
				Name:        "state",
				Description: "Indicates whether the crawler is running, or whether a run is pending",
				Type:        schema.TypeString,
			},
			{
				Name:        "table_prefix",
				Description: "The prefix added to the names of tables that are created",
				Type:        schema.TypeString,
			},
			{
				Name:        "version",
				Description: "The version of the crawler",
				Type:        schema.TypeBigInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_glue_crawler_targets_catalog_targets",
				Description: "Specifies an Glue Data Catalog target",
				Resolver:    schema.PathTableResolver("Targets.CatalogTargets"),
				Columns: []schema.Column{
					{
						Name:        "crawler_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_crawlers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "database_name",
						Description: "The name of the database to be synchronized",
						Type:        schema.TypeString,
					},
					{
						Name:        "tables",
						Description: "A list of the tables to be synchronized",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "connection_name",
						Description: "The name of the connection for an Amazon S3-backed Data Catalog table to be a target of the crawl when using a Catalog connection type paired with a NETWORK Connection type",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_glue_crawler_targets_delta_targets",
				Description: "Specifies a Delta data store to crawl one or more Delta tables",
				Resolver:    schema.PathTableResolver("Targets.DeltaTargets"),
				Columns: []schema.Column{
					{
						Name:        "crawler_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_crawlers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "connection_name",
						Description: "The name of the connection to use to connect to the Delta table target",
						Type:        schema.TypeString,
					},
					{
						Name:        "delta_tables",
						Description: "A list of the Amazon S3 paths to the Delta tables",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "write_manifest",
						Description: "Specifies whether to write the manifest files to the Delta table path",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "aws_glue_crawler_targets_dynamo_db_targets",
				Description: "Specifies an Amazon DynamoDB table to crawl",
				Resolver:    schema.PathTableResolver("Targets.DynamoDBTargets"),
				Columns: []schema.Column{
					{
						Name:        "crawler_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_crawlers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "path",
						Description: "The name of the DynamoDB table to crawl",
						Type:        schema.TypeString,
					},
					{
						Name:        "scan_all",
						Description: "Indicates whether to scan all the records, or to sample rows from the table Scanning all the records can take a long time when the table is not a high throughput table",
						Type:        schema.TypeBool,
					},
					{
						Name:        "scan_rate",
						Description: "The percentage of the configured read capacity units to use by the Glue crawler Read capacity units is a term defined by DynamoDB, and is a numeric value that acts as rate limiter for the number of reads that can be performed on that table per second",
						Type:        schema.TypeFloat,
					},
				},
			},
			{
				Name:        "aws_glue_crawler_targets_jdbc_targets",
				Description: "Specifies a JDBC data store to crawl",
				Resolver:    schema.PathTableResolver("Targets.JdbcTargets"),
				Columns: []schema.Column{
					{
						Name:        "crawler_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_crawlers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "connection_name",
						Description: "The name of the connection to use to connect to the JDBC target",
						Type:        schema.TypeString,
					},
					{
						Name:        "exclusions",
						Description: "A list of glob patterns used to exclude from the crawl",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "path",
						Description: "The path of the JDBC target",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_glue_crawler_targets_mongo_db_targets",
				Description: "Specifies an Amazon DocumentDB or MongoDB data store to crawl",
				Resolver:    schema.PathTableResolver("Targets.MongoDBTargets"),
				Columns: []schema.Column{
					{
						Name:        "crawler_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_crawlers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "connection_name",
						Description: "The name of the connection to use to connect to the Amazon DocumentDB or MongoDB target",
						Type:        schema.TypeString,
					},
					{
						Name:        "path",
						Description: "The path of the Amazon DocumentDB or MongoDB target (database/collection)",
						Type:        schema.TypeString,
					},
					{
						Name:        "scan_all",
						Description: "Indicates whether to scan all the records, or to sample rows from the table Scanning all the records can take a long time when the table is not a high throughput table",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "aws_glue_crawler_targets_s3_targets",
				Description: "Specifies a data store in Amazon Simple Storage Service (Amazon S3)",
				Resolver:    schema.PathTableResolver("Targets.S3Targets"),
				Columns: []schema.Column{
					{
						Name:        "crawler_cq_id",
						Description: "Unique CloudQuery ID of aws_glue_crawlers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "connection_name",
						Description: "The name of a connection which allows a job or crawler to access data in Amazon S3 within an Amazon Virtual Private Cloud environment (Amazon VPC)",
						Type:        schema.TypeString,
					},
					{
						Name:        "dlq_event_queue_arn",
						Description: "A valid Amazon dead-letter SQS ARN",
						Type:        schema.TypeString,
					},
					{
						Name:        "event_queue_arn",
						Description: "A valid Amazon SQS ARN",
						Type:        schema.TypeString,
					},
					{
						Name:        "exclusions",
						Description: "A list of glob patterns used to exclude from the crawl",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "path",
						Description: "The path to the Amazon S3 target",
						Type:        schema.TypeString,
					},
					{
						Name:        "sample_size",
						Description: "Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset",
						Type:        schema.TypeBigInt,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGlueCrawlers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Glue
	input := glue.GetCrawlersInput{}
	for {
		output, err := svc.GetCrawlers(ctx, &input, func(o *glue.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.Crawlers

		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
func resolveGlueCrawlerArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	arn := aws.String(crawlerARN(cl, aws.ToString(resource.Item.(types.Crawler).Name)))
	return diag.WrapError(resource.Set(c.Name, arn))
}
func resolveGlueCrawlerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetTagsInput{
		ResourceArn: aws.String(crawlerARN(cl, aws.ToString(resource.Item.(types.Crawler).Name))),
	}

	response, err := svc.GetTags(ctx, &input, func(options *glue.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, response.Tags))
}

func crawlerARN(cl *client.Client, name string) string {
	return cl.ARN(client.GlueService, "crawler", name)
}
