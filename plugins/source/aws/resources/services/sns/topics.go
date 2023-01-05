package sns

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sns/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Topics() *schema.Table {
	return &schema.Table{
		Name:                "aws_sns_topics",
		Resolver:            fetchSnsTopics,
		PreResourceResolver: getTopic,
		Transform:           transformers.TransformWithStruct(&models.Topic{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("sns"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveSnsTopicTags,
			},
			{
				Name:     "delivery_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeliveryPolicy"),
			},
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Policy"),
			},
			{
				Name:     "effective_delivery_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EffectiveDeliveryPolicy"),
			},
		},
	}
}
