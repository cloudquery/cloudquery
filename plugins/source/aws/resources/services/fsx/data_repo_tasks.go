package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DataRepoTasks() *schema.Table {
	return &schema.Table{
		Name:        "aws_fsx_data_repo_tasks",
		Description: "A description of the data repository task",
		Resolver:    fetchFsxDataRepoTasks,
		Multiplex:   client.ServiceAccountRegionMultiplexer("fsx"),
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
				Name:     "failure_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FailureDetails"),
			},
			{
				Name:        "paths",
				Description: "An array of paths on the Amazon FSx for Lustre file system that specify the data for the data repository task to process",
				Type:        schema.TypeStringArray,
			},
			{
				Name:     "report",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Report"),
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) for a given resource",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ResourceARN"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "start_time",
				Description: "The time that Amazon FSx began processing the task",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
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
			return err
		}
		res <- result.DataRepositoryTasks
	}
	return nil
}
