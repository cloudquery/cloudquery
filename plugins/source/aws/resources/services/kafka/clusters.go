package kafka

import (
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:                "aws_kafka_clusters",
		Description:         `https://docs.aws.amazon.com/MSK/2.0/APIReference/v2-clusters-clusterarn.html#v2-clusters-clusterarn-properties`,
		Resolver:            fetchKafkaClusters,
		PreResourceResolver: getCluster,
		Transform:           transformers.TransformWithStruct(&types.Cluster{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("kafka"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			Nodes(),
			ClusterOperations(),
		},
	}
}
