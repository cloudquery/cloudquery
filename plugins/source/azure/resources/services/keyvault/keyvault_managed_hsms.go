package keyvault

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func KeyvaultManagedHsms() *schema.Table {
	return &schema.Table{
		Name:      "azure_keyvault_keyvault_managed_hsms",
		Resolver:  fetchKeyvaultManagedHsms,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_keyvault_keyvault_managed_hsms", client.Namespacemicrosoft_keyvault),
		Transform: transformers.TransformWithStruct(&armkeyvault.ManagedHsm{}),
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

func fetchKeyvaultManagedHsms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armkeyvault.NewManagedHsmsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListBySubscriptionPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
