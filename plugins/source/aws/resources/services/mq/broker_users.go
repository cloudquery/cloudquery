package mq

import (
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func BrokerUsers() *schema.Table {
	tableName := "aws_mq_broker_users"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazon-mq/latest/api-reference/brokers-broker-id-users-username.html`,
		Resolver:    fetchMqBrokerUsers,
		Transform:   client.TransformWithStruct(&mq.DescribeUserOutput{}),
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
