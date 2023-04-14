package mq

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Brokers() *schema.Table {
	tableName := "aws_mq_brokers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/amazon-mq/latest/api-reference/brokers.html`,
		Resolver:            fetchMqBrokers,
		PreResourceResolver: getMqBroker,
		Transform:           transformers.TransformWithStruct(&mq.DescribeBrokerOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "mq"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BrokerArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			brokerConfigurations(),
			brokerUsers(),
		},
	}
}

func fetchMqBrokers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config mq.ListBrokersInput
	c := meta.(*client.Client)
	svc := c.Services().Mq
	paginator := mq.NewListBrokersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.BrokerSummaries
	}
	return nil
}

func getMqBroker(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Mq
	bs := resource.Item.(types.BrokerSummary)

	output, err := svc.DescribeBroker(ctx, &mq.DescribeBrokerInput{BrokerId: bs.BrokerId})
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}
