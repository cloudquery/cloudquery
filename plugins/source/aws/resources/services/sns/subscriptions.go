package sns

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sns/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Subscriptions() *schema.Table {
	return &schema.Table{
		Name:                "aws_sns_subscriptions",
		Description:         `https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html`,
		Resolver:            fetchSnsSubscriptions,
		PreResourceResolver: getSnsSubscription,
		Transform:           transformers.TransformWithStruct(&models.Subscription{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("sns"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubscriptionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "delivery_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeliveryPolicy"),
			},
			{
				Name:     "effective_delivery_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EffectiveDeliveryPolicy"),
			},
			{
				Name:     "filter_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FilterPolicy"),
			},
			{
				Name:     "redrive_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RedrivePolicy"),
			},
		},
	}
}
