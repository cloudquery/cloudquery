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

//go:generate cq-gen --resource workflows --config workflows.hcl --output .
func Workflows() *schema.Table {
	return &schema.Table{
		Name:         "aws_glue_workflows",
		Description:  "A workflow is a collection of multiple dependent Glue jobs and crawlers that are run to complete a complex ETL task",
		Resolver:     fetchGlueWorkflows,
		Multiplex:    client.ServiceAccountRegionMultiplexer("glue"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
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
				Description: "The Amazon Resource Name (ARN) of the workflow.",
				Type:        schema.TypeString,
				Resolver:    resolveGlueWorkflowArn,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveGlueWorkflowTags,
			},
			{
				Name:          "blueprint_name",
				Description:   "The name of the blueprint.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("BlueprintDetails.BlueprintName"),
				IgnoreInTests: true,
			},
			{
				Name:          "blueprint_run_id",
				Description:   "The run ID for this blueprint.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("BlueprintDetails.RunId"),
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
				Name:          "last_run_completed_on",
				Description:   "The date and time when the workflow run completed.",
				Type:          schema.TypeTimestamp,
				Resolver:      schema.PathResolver("LastRun.CompletedOn"),
				IgnoreInTests: true,
			},
			{
				Name:          "last_run_error_message",
				Description:   "This error message describes any error that may have occurred in starting the workflow run",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("LastRun.ErrorMessage"),
				IgnoreInTests: true,
			},
			{
				Name:          "last_run_name",
				Description:   "Name of the workflow that was run.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("LastRun.Name"),
				IgnoreInTests: true,
			},
			{
				Name:          "last_run_previous_run_id",
				Description:   "The ID of the previous workflow run.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("LastRun.PreviousRunId"),
				IgnoreInTests: true,
			},
			{
				Name:          "last_run_started_on",
				Description:   "The date and time when the workflow run was started.",
				Type:          schema.TypeTimestamp,
				Resolver:      schema.PathResolver("LastRun.StartedOn"),
				IgnoreInTests: true,
			},
			{
				Name:          "last_run_starting_event_batch_condition_size",
				Description:   "Number of events in the batch.",
				Type:          schema.TypeBigInt,
				Resolver:      schema.PathResolver("LastRun.StartingEventBatchCondition.BatchSize"),
				IgnoreInTests: true,
			},
			{
				Name:          "last_run_starting_event_batch_condition_window",
				Description:   "Duration of the batch window in seconds.",
				Type:          schema.TypeBigInt,
				Resolver:      schema.PathResolver("LastRun.StartingEventBatchCondition.BatchWindow"),
				IgnoreInTests: true,
			},
			{
				Name:        "last_run_statistics_failed_actions",
				Description: "Total number of Actions that have failed.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("LastRun.Statistics.FailedActions"),
			},
			{
				Name:        "last_run_statistics_running_actions",
				Description: "Total number Actions in running state.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("LastRun.Statistics.RunningActions"),
			},
			{
				Name:        "last_run_statistics_stopped_actions",
				Description: "Total number of Actions that have stopped.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("LastRun.Statistics.StoppedActions"),
			},
			{
				Name:        "last_run_statistics_succeeded_actions",
				Description: "Total number of Actions that have succeeded.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("LastRun.Statistics.SucceededActions"),
			},
			{
				Name:        "last_run_statistics_timeout_actions",
				Description: "Total number of Actions that timed out.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("LastRun.Statistics.TimeoutActions"),
			},
			{
				Name:        "last_run_statistics_total_actions",
				Description: "Total number of Actions in the workflow run.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("LastRun.Statistics.TotalActions"),
			},
			{
				Name:        "last_run_status",
				Description: "The status of the workflow run.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LastRun.Status"),
			},
			{
				Name:          "last_run_workflow_run_id",
				Description:   "The ID of this workflow run.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("LastRun.WorkflowRunId"),
				IgnoreInTests: true,
			},
			{
				Name:          "last_run_workflow_run_properties",
				Description:   "The workflow run properties which were set during the run.",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("LastRun.WorkflowRunProperties"),
				IgnoreInTests: true,
			},
			{
				Name:        "max_concurrent_runs",
				Description: "You can use this parameter to prevent unwanted multiple updates to data, to control costs, or in some cases, to prevent exceeding the maximum number of concurrent runs of any of the component jobs",
				Type:        schema.TypeBigInt,
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
			return diag.WrapError(err)
		}
		for _, name := range result.Workflows {
			w, err := svc.GetWorkflow(ctx, &glue.GetWorkflowInput{Name: aws.String(name)})
			if err != nil {
				if cl.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
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
	return diag.WrapError(resource.Set(c.Name, arn))
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
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, result.Tags))
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func workflowARN(cl *client.Client, name string) string {
	return cl.ARN(client.GlueService, "workflow", name)
}
