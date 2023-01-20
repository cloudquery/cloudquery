package security

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Pricings() *schema.Table {
	return &schema.Table{
		Name:        "azure_security_pricings",
		Resolver:    fetchPricings,
		Description: "https://learn.microsoft.com/en-us/rest/api/defenderforcloud/pricings/list?tabs=HTTP#pricing",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_security_pricings", client.Namespacemicrosoft_security),
		Transform:   transformers.TransformWithStruct(&armsecurity.Pricing{}),
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
