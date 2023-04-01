package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func nodes() *schema.Table {
	tableName := "aws_kafka_nodes"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-nodes.html#ListNodes`,
		Resolver:    fetchKafkaNodes,
		Transform:   transformers.TransformWithStruct(&types.NodeInfo{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "kafka"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "cluster_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchKafkaNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
