package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/thoas/go-funk"
)

func getListNodesInput(parent *schema.Resource) kafka.ListNodesInput {
	return kafka.ListNodesInput{
		ClusterArn: parent.Item.(*types.Cluster).ClusterArn,
		MaxResults: 100,
	}
}

func getListClusterOperationsInput(parent *schema.Resource) kafka.ListClusterOperationsInput {
	return kafka.ListClusterOperationsInput{
		ClusterArn: parent.Item.(*types.Cluster).ClusterArn,
		MaxResults: 100,
	}
}

func describeClustersInput(resource *schema.Resource) kafka.DescribeClusterV2Input {
	return kafka.DescribeClusterV2Input{
		ClusterArn: resource.Item.(types.Cluster).ClusterArn,
	}
}

func resolveKafkaTags(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		arn := funk.Get(r.Item, path, funk.WithAllowZero()).(*string)
		cl := meta.(*client.Client)
		svc := cl.Services().Kafka
		params := kafka.ListTagsForResourceInput{ResourceArn: arn}

		output, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		return r.Set(c.Name, output.Tags)
	}
}
