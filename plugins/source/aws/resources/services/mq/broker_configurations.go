package mq

import (
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BrokerConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "aws_mq_broker_configurations",
		Description: `https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id.html`,
		Resolver:    fetchMqBrokerConfigurations,
		Transform:   transformers.TransformWithStruct(&mq.DescribeConfigurationOutput{}, transformers.WithSkipFields("ResultMetadata"), transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("mq"),
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
