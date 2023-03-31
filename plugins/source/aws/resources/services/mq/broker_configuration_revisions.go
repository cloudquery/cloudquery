package mq

import (
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func brokerConfigurationRevisions() *schema.Table {
	tableName := "aws_mq_broker_configuration_revisions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id-revisions.html`,
		Resolver:            fetchMqBrokerConfigurationRevisions,
		PreResourceResolver: getMqBrokerConfigurationRevision,
		Transform:           transformers.TransformWithStruct(&mq.DescribeConfigurationRevisionOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "mq"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "broker_configuration_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "data",
				Type:     schema.TypeJSON,
				Resolver: resolveBrokerConfigurationRevisionsData,
			},
		},
	}
}
