package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Workflows() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_workflows",
		Description: "A workflow is a collection of multiple dependent Glue jobs and crawlers that are run to complete a complex ETL task",
		Resolver:    fetchGlueWorkflows,
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
				Resolver:        resolveGlueWorkflowArn,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveGlueWorkflowTags,
			},
			{
				Name:          "blueprint_details",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("BlueprintDetails"),
				IgnoreInTests: true,
			},
			{
				Name:        "created_on",
				Description: "The date and time when the workflow was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:          "default_run_properties",
				Description:   "A collection of properties to be used as part of each execution of the workflow. The run properties are made available to each job in the workflow",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "description",
				Description: "A description of the workflow.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_modified_on",
				Description: "The date and time when the workflow was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:     "last_run",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastRun"),
			},
			{
				Name:        "max_concurrent_runs",
				Description: "You can use this parameter to prevent unwanted multiple updates to data, to control costs, or in some cases, to prevent exceeding the maximum number of concurrent runs of any of the component jobs",
				Type:        schema.TypeInt,
			},
			{
				Name:        "name",
				Description: "The name of the workflow.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGlueWorkflows(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.ListWorkflowsInput{MaxResults: aws.Int32(25)}
	for {
		result, err := svc.ListWorkflows(ctx, &input)
		if err != nil {
			return err
		}
		for _, name := range result.Workflows {
			w, err := svc.GetWorkflow(ctx, &glue.GetWorkflowInput{Name: aws.String(name)})
			if err != nil {
				if cl.IsNotFoundError(err) {
					continue
				}
				return err
			}
			if w.Workflow != nil {
				res <- *w.Workflow
			}
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveGlueWorkflowArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	arn := aws.String(workflowARN(cl, aws.ToString(resource.Item.(types.Workflow).Name)))
	return resource.Set(c.Name, arn)
}
func resolveGlueWorkflowTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(workflowARN(cl, aws.ToString(resource.Item.(types.Workflow).Name))),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func workflowARN(cl *client.Client, name string) string {
	return cl.ARN(client.GlueService, "workflow", name)
}
