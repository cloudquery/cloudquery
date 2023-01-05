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
		Resolver:            fetchSnsSubscriptions,
		PreResourceResolver: getSnsSubscription,
		Transform:           transformers.TransformWithStruct(&models.Subscription{}),
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
				Resolver: schema.PathResolver("SubscriptionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
