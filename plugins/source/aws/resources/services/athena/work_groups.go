package athena

import (
	"context"
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_athena_work_groups",
		Description: "A workgroup, which contains a name, description, creation time, state, and other configuration, listed under WorkGroup$Configuration",
		Resolver:    fetchAthenaWorkGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("athena"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:            "arn",
				Description:     "ARN of the resource.",
				Type:            schema.TypeString,
				Resolver:        resolveAthenaWorkGroupArn,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "tags",
				Description: "Tags associated with the Athena work group.",
				Type:        schema.TypeJSON,
				Resolver:    resolveAthenaWorkGroupTags,
			},
			{
				Name:        "name",
				Description: "The workgroup name",
				Type:        schema.TypeString,
			},
			{
				Name:          "bytes_scanned_cutoff_per_query",
				Description:   "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan",
				Type:          schema.TypeInt,
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
				Name:     "configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Configuration"),
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
						Name:     "engine_version",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("EngineVersion"),
					},
					{
						Name:        "execution_parameters",
						Description: "A list of values for the parameters in a query",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "query",
						Description: "The SQL query statements which the query execution ran",
						Type:        schema.TypeString,
					},
					{
						Name:     "query_execution_context",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("QueryExecutionContext"),
					},
					{
						Name:        "id",
						Description: "The unique identifier for each query execution",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("QueryExecutionId"),
					},
					{
						Name:     "result_configuration",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("ResultConfiguration"),
					},
					{
						Name:        "statement_type",
						Description: "The type of query statement that was run",
						Type:        schema.TypeString,
					},
					{
						Name:     "statistics",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Statistics"),
					},
					{
						Name:     "status",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Status"),
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
	return client.ListAndDetailResolver(ctx, meta, res, listWorkGroups, workGroupDetail)
}
func resolveAthenaWorkGroupArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	dc := resource.Item.(types.WorkGroup)
	return resource.Set(c.Name, createWorkGroupArn(cl, *dc.Name))
}
func resolveAthenaWorkGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
			return err
		}
		client.TagsIntoMap(result.Tags, tags)
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return resource.Set(c.Name, tags)
}
func fetchAthenaWorkGroupPreparedStatements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	wg := parent.Item.(types.WorkGroup)
	input := athena.ListPreparedStatementsInput{WorkGroup: wg.Name}
	for {
		response, err := svc.ListPreparedStatements(ctx, &input)
		if err != nil {
			return err
		}
		for _, d := range response.PreparedStatements {
			dc, err := svc.GetPreparedStatement(ctx, &athena.GetPreparedStatementInput{
				WorkGroup:     wg.Name,
				StatementName: d.StatementName,
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
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
		response, err := svc.ListQueryExecutions(ctx, &input)
		if err != nil {
			return err
		}
		for _, d := range response.QueryExecutionIds {
			dc, err := svc.GetQueryExecution(ctx, &athena.GetQueryExecutionInput{
				QueryExecutionId: aws.String(d),
			})
			if err != nil {
				if c.IsNotFoundError(err) || isQueryExecutionNotFound(err) {
					continue
				}
				return err
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
		response, err := svc.ListNamedQueries(ctx, &input)
		if err != nil {
			return err
		}
		for _, d := range response.NamedQueryIds {
			dc, err := svc.GetNamedQuery(ctx, &athena.GetNamedQueryInput{
				NamedQueryId: aws.String(d),
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
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

func listWorkGroups(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	input := athena.ListWorkGroupsInput{}
	for {
		response, err := svc.ListWorkGroups(ctx, &input, func(options *athena.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, item := range response.WorkGroups {
			detailChan <- item
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
func workGroupDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, summary interface{}) {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	wg := summary.(types.WorkGroupSummary)
	dc, err := svc.GetWorkGroup(ctx, &athena.GetWorkGroupInput{
		WorkGroup: wg.Name,
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return
		}
		errorChan <- err
		return
	}
	resultsChan <- *dc.WorkGroup
}
func createWorkGroupArn(cl *client.Client, groupName string) string {
	return cl.ARN(client.Athena, "workgroup", groupName)
}
func isQueryExecutionNotFound(err error) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	return ae.ErrorCode() == "InvalidRequestException" && strings.Contains(ae.ErrorMessage(), "was not found")
}
