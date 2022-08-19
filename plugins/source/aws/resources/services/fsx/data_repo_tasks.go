package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource data_repo_tasks --config data_repo_tasks.hcl --output .
func DataRepoTasks() *schema.Table {
	return &schema.Table{
		Name:         "aws_fsx_data_repo_tasks",
		Description:  "A description of the data repository task",
		Resolver:     fetchFsxDataRepoTasks,
		Multiplex:    client.ServiceAccountRegionMultiplexer("fsx"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "creation_time",
				Description: "The time that the resource was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "file_system_id",
				Description: "The globally unique ID of the file system, assigned by Amazon FSx",
				Type:        schema.TypeString,
			},
			{
				Name:        "lifecycle",
				Description: "The lifecycle status of the data repository task, as follows:  * PENDING - Amazon FSx has not started the task",
				Type:        schema.TypeString,
			},
			{
				Name:        "task_id",
				Description: "The system-generated, unique 17-digit ID of the data repository task",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of data repository task",
				Type:        schema.TypeString,
			},
			{
				Name:        "end_time",
				Description: "The time that Amazon FSx completed processing the task, populated after the task is complete",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "failure_details_message",
				Description: "A detailed error message",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FailureDetails.Message"),
			},
			{
				Name:        "paths",
				Description: "An array of paths on the Amazon FSx for Lustre file system that specify the data for the data repository task to process",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "report_enabled",
				Description: "Set Enabled to True to generate a CompletionReport when the task completes",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Report.Enabled"),
			},
			{
				Name:        "report_format",
				Description: "Required if Enabled is set to true",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Report.Format"),
			},
			{
				Name:        "report_path",
				Description: "Required if Enabled is set to true",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Report.Path"),
			},
			{
				Name:        "report_scope",
				Description: "Required if Enabled is set to true",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Report.Scope"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for a given resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceARN"),
			},
			{
				Name:        "start_time",
				Description: "The time that Amazon FSx began processing the task",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "status_failed_count",
				Description: "A running total of the number of files that the task failed to process",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Status.FailedCount"),
			},
			{
				Name:        "status_last_updated_time",
				Description: "The time at which the task status was last updated",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Status.LastUpdatedTime"),
			},
			{
				Name:        "status_succeeded_count",
				Description: "A running total of the number of files that the task has successfully processed",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Status.SucceededCount"),
			},
			{
				Name:        "status_total_count",
				Description: "The total number of files that the task will process",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Status.TotalCount"),
			},
			{
				Name:        "tags",
				Description: "A list of Tag values, with a maximum of 50 elements",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchFsxDataRepoTasks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().FSX
	input := fsx.DescribeDataRepositoryTasksInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeDataRepositoryTasksPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- result.DataRepositoryTasks
	}
	return nil
}
