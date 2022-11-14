package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchKafkaClusterOperations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	if parent.Item.(*types.Cluster).ClusterType == types.ClusterTypeServerless {
		// serverless clusters do not support cluster operations
		return nil
	}

	var input = getListClusterOperationsInput(parent)
	c := meta.(*client.Client)
	svc := c.Services().Kafka
	for {
		response, err := svc.ListClusterOperations(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.ClusterOperationInfoList
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
