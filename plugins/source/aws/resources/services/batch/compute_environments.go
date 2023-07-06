package batch

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ComputeEnvironments() *schema.Table {
	tableName := "aws_batch_compute_environments"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeComputeEnvironments.html`,
		Resolver:    fetchBatchComputeEnvironments,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "batch"),
		Transform:   transformers.TransformWithStruct(&types.ComputeEnvironmentDetail{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveBatchComputeEnvironmentTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ComputeEnvironmentArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchBatchComputeEnvironments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := batch.DescribeComputeEnvironmentsInput{
		MaxResults: aws.Int32(100),
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Batch
	p := batch.NewDescribeComputeEnvironmentsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *batch.Options) {
			options.Region = cl.Region
		})
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
	output, err := svc.ListTagsForResource(ctx, &input, func(options *batch.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, output.Tags)
}
