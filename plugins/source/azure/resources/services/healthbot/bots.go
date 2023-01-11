package healthbot

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Bots() *schema.Table {
	return &schema.Table{
		Name:      "azure_healthbot_bots",
		Resolver:  fetchBots,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_healthbot_bots", client.Namespacemicrosoft_healthbot),
		Transform: transformers.TransformWithStruct(&armhealthbot.HealthBot{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
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
