package mq

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func brokerUsers() *schema.Table {
	tableName := "aws_mq_broker_users"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazon-mq/latest/api-reference/brokers-broker-id-users-username.html`,
		Resolver:    fetchMqBrokerUsers,
		Transform:   transformers.TransformWithStruct(&mq.DescribeUserOutput{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "mq"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "broker_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchMqBrokerUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	broker := parent.Item.(*mq.DescribeBrokerOutput)
	c := meta.(*client.Client)
	svc := c.Services().Mq
	for _, us := range broker.Users {
		input := mq.DescribeUserInput{
			BrokerId: broker.BrokerId,
			Username: us.Username,
		}
		output, err := svc.DescribeUser(ctx, &input)
		if err != nil {
			return err
		}
		res <- output
	}
	return nil
}
