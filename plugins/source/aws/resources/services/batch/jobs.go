package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

var allJobStatuses = []types.JobStatus{
	types.JobStatusSubmitted,
	types.JobStatusPending,
	types.JobStatusRunnable,
	types.JobStatusStarting,
	types.JobStatusRunning,
	types.JobStatusSucceeded,
	types.JobStatusFailed,
}

func jobs() *schema.Table {
	tableName := "aws_batch_jobs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeJobs.html`,
		Resolver:    fetchBatchJobs,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "batch"),
		Transform:   transformers.TransformWithStruct(&types.JobDetail{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveBatchJobTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JobArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchBatchJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	for _, status := range allJobStatuses {
		config := batch.ListJobsInput{
			MaxResults: aws.Int32(100),
			JobQueue:   parent.Item.(types.JobQueueDetail).JobQueueArn,
			JobStatus:  status,
		}
		c := meta.(*client.Client)
		svc := c.Services().Batch
		p := batch.NewListJobsPaginator(svc, &config)
		for p.HasMorePages() {
			response, err := p.NextPage(ctx)
			if err != nil {
				return err
			}
			if len(response.JobSummaryList) == 0 {
				continue
			}

			// fetch details for each job
			ids := make([]string, len(response.JobSummaryList))
			for i, v := range response.JobSummaryList {
				ids[i] = *v.JobId
			}
			details, err := svc.DescribeJobs(ctx, &batch.DescribeJobsInput{
				Jobs: ids,
			})
			if err != nil {
				return err
			}

			res <- details.Jobs
		}
	}
	return nil
}

func resolveBatchJobTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Batch
	summary := resource.Item.(types.JobDetail)

	input := batch.ListTagsForResourceInput{
		ResourceArn: summary.JobArn,
	}
	output, err := svc.ListTagsForResource(ctx, &input)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, output.Tags)
}
