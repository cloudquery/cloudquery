package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchKafkaNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Kafka
	cluster := parent.Item.(types.Cluster)
	input := &kafka.ListNodesInput{
		ClusterArn: cluster.ClusterArn,
	}
	paginator := kafka.NewListNodesPaginator(svc, input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.NodeInfoList
	}
	return nil
}
