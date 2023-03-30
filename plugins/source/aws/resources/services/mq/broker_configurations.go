package mq

import (
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BrokerConfigurations() *schema.Table {
	tableName := "aws_mq_broker_configurations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id.html`,
		Resolver:    fetchMqBrokerConfigurations,
		Transform: client.TransformWithStruct(&mq.DescribeConfigurationOutput{},
			transformers.WithSkipFields("ResultMetadata"),
			transformers.WithPrimaryKeys("Arn"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "mq"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "broker_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},

		Relations: []*schema.Table{
			BrokerConfigurationRevisions(),
		},
	}
}
