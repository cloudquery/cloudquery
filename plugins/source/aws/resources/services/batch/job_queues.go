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

func JobQueues() *schema.Table {
	tableName := "aws_batch_job_queues"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeJobQueues.html`,
		Resolver:    fetchBatchJobQueues,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "batch"),
		Transform:   transformers.TransformWithStruct(&types.JobQueueDetail{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveBatchJobQueueTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JobQueueArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			jobs(),
		},
	}
}

func fetchBatchJobQueues(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := batch.DescribeJobQueuesInput{
		MaxResults: aws.Int32(100),
	}
	c := meta.(*client.Client)
	svc := c.Services().Batch
	p := batch.NewDescribeJobQueuesPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.JobQueues
	}
	return nil
}

func resolveBatchJobQueueTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Batch
	summary := resource.Item.(types.JobQueueDetail)

	input := batch.ListTagsForResourceInput{
		ResourceArn: summary.JobQueueArn,
	}
	output, err := svc.ListTagsForResource(ctx, &input)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, output.Tags)
}
