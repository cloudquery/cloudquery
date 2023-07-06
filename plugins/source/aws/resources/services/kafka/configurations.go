package kafka

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Configurations() *schema.Table {
	tableName := "aws_kafka_configurations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-configuration.html`,
		Resolver:    fetchKafkaConfigurations,
		Transform:   transformers.TransformWithStruct(&types.Configuration{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "kafka"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchKafkaConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input kafka.ListConfigurationsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Kafka
	paginator := kafka.NewListConfigurationsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *kafka.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Configurations
	}
	return nil
}
