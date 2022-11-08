package kafka

import (
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/plugin-sdk/schema"
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
