package kafka

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v3/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func clusterOperations() *schema.Table {
	tableName := "aws_kafka_cluster_operations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-operations.html`,
		Resolver:    fetchKafkaClusterOperations,
		Transform:   transformers.TransformWithStruct(&types.ClusterOperationInfo{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "kafka"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("OperationArn"),
				PrimaryKey: true,
			},
			{
				Name:     "cluster_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveKafkaTags("OperationArn"),
			},
		},
	}
}

func fetchKafkaClusterOperations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	if parent.Item.(*types.Cluster).ClusterType == types.ClusterTypeServerless {
		// serverless clusters do not support cluster operations
		return nil
	}

	var input = getListClusterOperationsInput(parent)
	cl := meta.(*client.Client)
	svc := cl.Services().Kafka
	paginator := kafka.NewListClusterOperationsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *kafka.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ClusterOperationInfoList
	}
	return nil
}
