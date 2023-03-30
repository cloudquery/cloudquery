package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ComputeEnvironments() *schema.Table {
	tableName := "aws_batch_compute_environments"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeComputeEnvironments.html`,
		Resolver:    fetchBatchComputeEnvironments,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "batch"),
		Transform:   client.TransformWithStruct(&types.ComputeEnvironmentDetail{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveBatchComputeEnvironmentTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ComputeEnvironmentArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchBatchComputeEnvironments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := batch.DescribeComputeEnvironmentsInput{
		MaxResults: aws.Int32(100),
	}
	c := meta.(*client.Client)
	svc := c.Services().Batch
	p := batch.NewDescribeComputeEnvironmentsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.ComputeEnvironments
	}
	return nil
}

func resolveBatchComputeEnvironmentTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Batch
	summary := resource.Item.(types.ComputeEnvironmentDetail)

	input := batch.ListTagsForResourceInput{
		ResourceArn: summary.ComputeEnvironmentArn,
	}
	output, err := svc.ListTagsForResource(ctx, &input)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, output.Tags)
}
