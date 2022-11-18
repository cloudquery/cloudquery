package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchKafkaNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	if parent.Item.(*types.Cluster).ClusterType == types.ClusterTypeServerless {
		// serverless clusters do not have nodes
		return nil
	}
	if parent.Item.(*types.Cluster).State == types.ClusterStateCreating {
		// nodes for clusters in state "CREATING" cannot be listed
		return nil
	}

	var input = getListNodesInput(parent)
	c := meta.(*client.Client)
	svc := c.Services().Kafka
	for {
		response, err := svc.ListNodes(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.NodeInfoList
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
