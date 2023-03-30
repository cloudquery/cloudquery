package shield

import (
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Subscriptions() *schema.Table {
	tableName := "aws_shield_subscriptions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_Subscription.html`,
		Resolver:    fetchShieldSubscriptions,
		Transform:   client.TransformWithStruct(&types.Subscription{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "shield"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
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
