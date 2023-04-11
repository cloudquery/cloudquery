package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Clusters() *schema.Table {
	tableName := "aws_kafka_clusters"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/MSK/2.0/APIReference/v2-clusters-clusterarn.html#v2-clusters-clusterarn-properties`,
		Resolver:            fetchKafkaClusters,
		PreResourceResolver: getCluster,
		Transform:           transformers.TransformWithStruct(&types.Cluster{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "kafka"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
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
			nodes(),
			clusterOperations(),
		},
	}
}

func fetchKafkaClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input kafka.ListClustersV2Input
	c := meta.(*client.Client)
	svc := c.Services().Kafka
	paginator := kafka.NewListClustersV2Paginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.ClusterInfoList
	}
	return nil
}

func getCluster(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Kafka
	var input kafka.DescribeClusterV2Input = describeClustersInput(resource)
	output, err := svc.DescribeClusterV2(ctx, &input)
	if err != nil {
		return err
	}
	resource.Item = output.ClusterInfo
	return nil
}
