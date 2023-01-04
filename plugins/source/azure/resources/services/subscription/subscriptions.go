package subscription

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Subscriptions() *schema.Table {
	return &schema.Table{
		Name:      "azure_subscription_subscriptions",
		Resolver:  fetchSubscriptions,
		Multiplex: client.SingleSubscriptionMultiplex,
		Transform: transformers.TransformWithStruct(&armsubscription.Subscription{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
