// Code generated by codegen; DO NOT EDIT.

package kafka

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Nodes() *schema.Table {
	return &schema.Table{
		Name:        "aws_kafka_nodes",
		Description: `https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-nodes.html#ListNodes`,
		Resolver:    fetchKafkaNodes,
		Multiplex:   client.ServiceAccountRegionMultiplexer("kafka"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
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
			{
				Name:     "added_to_cluster_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AddedToClusterTime"),
			},
			{
				Name:     "broker_node_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BrokerNodeInfo"),
			},
			{
				Name:     "instance_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceType"),
			},
			{
				Name:     "node_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeType"),
			},
			{
				Name:     "zookeeper_node_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ZookeeperNodeInfo"),
			},
		},
	}
}
