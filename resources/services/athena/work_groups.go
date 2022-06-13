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

//go:generate cq-gen --resource work_groups --config gen.hcl --output .
func WorkGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_athena_work_groups",
		Description:  "A workgroup, which contains a name, description, creation time, state, and other configuration, listed under WorkGroup$Configuration",
		Resolver:     fetchAthenaWorkGroups,
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
				Name:        "arn",
				Description: "ARN of the resource.",
				Type:        schema.TypeString,
				Resolver:    ResolveAthenaWorkGroupArn,
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
				Resolver: ResolveAthenaWorkGroupTags,
			},
			{
				Name:        "name",
				Description: "The workgroup name",
				Type:        schema.TypeString,
			},
			{
				Name:          "bytes_scanned_cutoff_per_query",
				Description:   "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan",
				Type:          schema.TypeBigInt,
				Resolver:      schema.PathResolver("Configuration.BytesScannedCutoffPerQuery"),
				IgnoreInTests: true,
			},
			{
				Name:        "enforce_work_group_configuration",
				Description: "If set to \"true\", the settings for the workgroup override client-side settings If set to \"false\", client-side settings are used",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Configuration.EnforceWorkGroupConfiguration"),
			},
			{
				Name:        "effective_engine_version",
				Description: "The engine version on which the query runs If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.EngineVersion.EffectiveEngineVersion"),
			},
			{
				Name:        "selected_engine_version",
				Description: "The engine version requested by the user",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.EngineVersion.SelectedEngineVersion"),
			},
			{
				Name:        "publish_cloud_watch_metrics_enabled",
				Description: "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Configuration.PublishCloudWatchMetricsEnabled"),
			},
			{
				Name:        "requester_pays_enabled",
				Description: "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Configuration.RequesterPaysEnabled"),
			},
			{
				Name:        "acl_configuration_s3_acl_option",
				Description: "The Amazon S3 canned ACL that Athena should specify when storing query results Currently the only supported canned ACL is BUCKET_OWNER_FULL_CONTROL",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.ResultConfiguration.AclConfiguration.S3AclOption"),
			},
			{
				Name:        "encryption_configuration_encryption_option",
				Description: "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE_S3), server-side encryption with KMS-managed keys (SSE_KMS), or client-side encryption with KMS-managed keys (CSE_KMS) is used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption"),
			},
			{
				Name:          "encryption_configuration_kms_key",
				Description:   "For SSE_KMS and CSE_KMS, this is the KMS key ARN or ID",
				Type:          schema.TypeString,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("Configuration.ResultConfiguration.EncryptionConfiguration.KmsKey"),
			},
			{
				Name:          "expected_bucket_owner",
				Description:   "The Amazon Web Services account ID that you expect to be the owner of the Amazon S3 bucket specified by ResultConfiguration$OutputLocation",
				Type:          schema.TypeString,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("Configuration.ResultConfiguration.ExpectedBucketOwner"),
			},
			{
				Name:          "output_location",
				Description:   "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/",
				Type:          schema.TypeString,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("Configuration.ResultConfiguration.OutputLocation"),
			},
			{
				Name:        "creation_time",
				Description: "The date and time the workgroup was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "The workgroup description",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The state of the workgroup: ENABLED or DISABLED",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_athena_work_group_prepared_statements",
				Description:   "A prepared SQL statement for use with Athena",
				Resolver:      fetchAthenaWorkGroupPreparedStatements,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "work_group_cq_id",
						Description: "Unique CloudQuery ID of aws_athena_work_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "description",
						Description: "The description of the prepared statement",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_modified_time",
						Description: "The last modified time of the prepared statement",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "query_statement",
						Description: "The query string for the prepared statement",
						Type:        schema.TypeString,
					},
					{
						Name:        "statement_name",
						Description: "The name of the prepared statement",
						Type:        schema.TypeString,
					},
					{
						Name:        "work_group_name",
						Description: "The name of the workgroup to which the prepared statement belongs",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_athena_work_group_query_executions",
				Description:   "Information about a single instance of a query execution",
				Resolver:      fetchAthenaWorkGroupQueryExecutions,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "work_group_cq_id",
						Description: "Unique CloudQuery ID of aws_athena_work_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "effective_engine_version",
						Description: "The engine version on which the query runs If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EngineVersion.EffectiveEngineVersion"),
					},
					{
						Name:        "selected_engine_version",
						Description: "The engine version requested by the user",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("EngineVersion.SelectedEngineVersion"),
					},
					{
						Name:        "query",
						Description: "The SQL query statements which the query execution ran",
						Type:        schema.TypeString,
					},
					{
						Name:        "catalog",
						Description: "The name of the data catalog used in the query execution",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("QueryExecutionContext.Catalog"),
					},
					{
						Name:        "database",
						Description: "The name of the database used in the query execution",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("QueryExecutionContext.Database"),
					},
					{
						Name:        "id",
						Description: "The unique identifier for each query execution",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("QueryExecutionId"),
					},
					{
						Name:        "acl_configuration_s3_acl_option",
						Description: "The Amazon S3 canned ACL that Athena should specify when storing query results Currently the only supported canned ACL is BUCKET_OWNER_FULL_CONTROL",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResultConfiguration.AclConfiguration.S3AclOption"),
					},
					{
						Name:        "encryption_configuration_encryption_option",
						Description: "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE_S3), server-side encryption with KMS-managed keys (SSE_KMS), or client-side encryption with KMS-managed keys (CSE_KMS) is used",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResultConfiguration.EncryptionConfiguration.EncryptionOption"),
					},
					{
						Name:        "encryption_configuration_kms_key",
						Description: "For SSE_KMS and CSE_KMS, this is the KMS key ARN or ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResultConfiguration.EncryptionConfiguration.KmsKey"),
					},
					{
						Name:        "expected_bucket_owner",
						Description: "The Amazon Web Services account ID that you expect to be the owner of the Amazon S3 bucket specified by ResultConfiguration$OutputLocation",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResultConfiguration.ExpectedBucketOwner"),
					},
					{
						Name:        "output_location",
						Description: "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResultConfiguration.OutputLocation"),
					},
					{
						Name:        "statement_type",
						Description: "The type of query statement that was run",
						Type:        schema.TypeString,
					},
					{
						Name:        "data_manifest_location",
						Description: "The location and file name of a data manifest file",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Statistics.DataManifestLocation"),
					},
					{
						Name:        "data_scanned_in_bytes",
						Description: "The number of bytes in the data that was queried",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Statistics.DataScannedInBytes"),
					},
					{
						Name:        "engine_execution_time_in_millis",
						Description: "The number of milliseconds that the query took to execute",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Statistics.EngineExecutionTimeInMillis"),
					},
					{
						Name:        "query_planning_time_in_millis",
						Description: "The number of milliseconds that Athena took to plan the query processing flow This includes the time spent retrieving table partitions from the data source Note that because the query engine performs the query planning, query planning time is a subset of engine processing time",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Statistics.QueryPlanningTimeInMillis"),
					},
					{
						Name:        "query_queue_time_in_millis",
						Description: "The number of milliseconds that the query was in your query queue waiting for resources",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Statistics.QueryQueueTimeInMillis"),
					},
					{
						Name:        "service_processing_time_in_millis",
						Description: "The number of milliseconds that Athena took to finalize and publish the query results after the query engine finished running the query",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Statistics.ServiceProcessingTimeInMillis"),
					},
					{
						Name:        "total_execution_time_in_millis",
						Description: "The number of milliseconds that Athena took to run the query",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Statistics.TotalExecutionTimeInMillis"),
					},
					{
						Name:        "athena_error_error_category",
						Description: "An integer value that specifies the category of a query failure error",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Status.AthenaError.ErrorCategory"),
					},
					{
						Name:        "athena_error_error_message",
						Description: "Contains a short description of the error that occurred",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Status.AthenaError.ErrorMessage"),
					},
					{
						Name:        "athena_error_error_type",
						Description: "An integer value that provides specific information about an Athena query error For the meaning of specific values, see the Error Type Reference (https://docsawsamazoncom/athena/latest/ug/error-referencehtml#error-reference-error-type-reference) in the Amazon Athena User Guide",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Status.AthenaError.ErrorType"),
					},
					{
						Name:        "athena_error_retryable",
						Description: "True if the query might succeed if resubmitted",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Status.AthenaError.Retryable"),
					},
					{
						Name:        "completion_date_time",
						Description: "The date and time that the query completed",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("Status.CompletionDateTime"),
					},
					{
						Name:        "state",
						Description: "The state of query execution",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Status.State"),
					},
					{
						Name:        "state_change_reason",
						Description: "Further detail about the status of the query",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Status.StateChangeReason"),
					},
					{
						Name:        "submission_date_time",
						Description: "The date and time that the query was submitted",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("Status.SubmissionDateTime"),
					},
					{
						Name:        "work_group",
						Description: "The name of the workgroup in which the query ran",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_athena_work_group_named_queries",
				Description: "A query, where QueryString contains the SQL statements that make up the query",
				Resolver:    fetchAthenaWorkGroupNamedQueries,
				Columns: []schema.Column{
					{
						Name:        "work_group_cq_id",
						Description: "Unique CloudQuery ID of aws_athena_work_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "database",
						Description: "The database to which the query belongs",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The query name",
						Type:        schema.TypeString,
					},
					{
						Name:        "query_string",
						Description: "The SQL statements that make up the query",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "The query description",
						Type:        schema.TypeString,
					},
					{
						Name:        "named_query_id",
						Description: "The unique identifier of the query",
						Type:        schema.TypeString,
					},
					{
						Name:        "work_group",
						Description: "The name of the workgroup that contains the named query",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAthenaWorkGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	input := athena.ListWorkGroupsInput{}
	var sem = semaphore.NewWeighted(int64(MAX_GOROUTINES))
	for {
		response, err := svc.ListWorkGroups(ctx, &input, func(options *athena.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		errs, ctx := errgroup.WithContext(ctx)
		for _, d := range response.WorkGroups {
			if err := sem.Acquire(ctx, 1); err != nil {
				return diag.WrapError(err)
			}
			func(summary types.WorkGroupSummary) {
				errs.Go(func() error {
					defer sem.Release(1)
					return fetchWorkGroup(ctx, res, svc, c.Region, summary)
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
func ResolveAthenaWorkGroupArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	dc := resource.Item.(types.WorkGroup)
	return diag.WrapError(resource.Set(c.Name, createWorkGroupArn(cl, *dc.Name)))
}
func ResolveAthenaWorkGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	wg := resource.Item.(types.WorkGroup)
	arn := createWorkGroupArn(cl, *wg.Name)
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
func fetchAthenaWorkGroupPreparedStatements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	wg := parent.Item.(types.WorkGroup)
	input := athena.ListPreparedStatementsInput{WorkGroup: wg.Name}
	for {
		response, err := svc.ListPreparedStatements(ctx, &input, func(options *athena.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, d := range response.PreparedStatements {
			dc, err := svc.GetPreparedStatement(ctx, &athena.GetPreparedStatementInput{
				WorkGroup:     wg.Name,
				StatementName: d.StatementName,
			}, func(options *athena.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}
			res <- *dc.PreparedStatement
			return nil
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func fetchAthenaWorkGroupQueryExecutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	wg := parent.Item.(types.WorkGroup)
	input := athena.ListQueryExecutionsInput{WorkGroup: wg.Name}
	for {
		response, err := svc.ListQueryExecutions(ctx, &input, func(options *athena.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, d := range response.QueryExecutionIds {
			dc, err := svc.GetQueryExecution(ctx, &athena.GetQueryExecutionInput{

				QueryExecutionId: aws.String(d),
			}, func(options *athena.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}
			res <- *dc.QueryExecution
			return nil
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func fetchAthenaWorkGroupNamedQueries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	wg := parent.Item.(types.WorkGroup)
	input := athena.ListNamedQueriesInput{WorkGroup: wg.Name}
	for {
		response, err := svc.ListNamedQueries(ctx, &input, func(options *athena.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, d := range response.NamedQueryIds {
			dc, err := svc.GetNamedQuery(ctx, &athena.GetNamedQueryInput{
				NamedQueryId: aws.String(d),
			}, func(options *athena.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}
			res <- *dc.NamedQuery
			return nil
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func fetchWorkGroup(ctx context.Context, res chan<- interface{}, svc client.AthenaClient, region string, groupSummary types.WorkGroupSummary) error {
	dc, err := svc.GetWorkGroup(ctx, &athena.GetWorkGroupInput{
		WorkGroup: groupSummary.Name,
	}, func(options *athena.Options) {
		options.Region = region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	res <- *dc.WorkGroup
	return nil
}

func createWorkGroupArn(cl *client.Client, groupName string) string {
	return cl.ARN(client.Athena, "workgroup", groupName)
}
