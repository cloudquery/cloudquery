package keyvault

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Keyvault() *schema.Table {
	return &schema.Table{
		Name:      "azure_keyvault_keyvault",
		Resolver:  fetchKeyvault,
		Multiplex: client.SubscriptionMultiplex,
		Transform: transformers.TransformWithStruct(&armkeyvault.Vault{}),
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

		Relations: []*schema.Table{
			keyvault_keys(),
			keyvault_secrets(),
		},
	}
}

func fetchKeyvault(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armkeyvault.NewVaultsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}

		for _, r := range p.Value {
			group, err := client.ParseResourceGroup(*r.ID)
			if err != nil {
				return err
			}
			resp, err := svc.Get(ctx, *r.Name, group, nil)
			if err != nil {
				return err
			}
			res <- resp.Vault
		}
	}
	return nil
}
