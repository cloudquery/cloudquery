package keyvault

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Keyvault() *schema.Table {
	return &schema.Table{
		Name:                 "azure_keyvault_keyvault",
		PreResourceResolver:  keyvaultGet,
		Resolver:             fetchKeyvault,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/keyvault/keyvault/vaults/get?tabs=HTTP#vault",
		Multiplex:            client.SubscriptionMultiplex,
		Transform:            transformers.TransformWithStruct(&armkeyvault.Vault{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},

		Relations: []*schema.Table{
			keyvault_keys(),
			keyvault_secrets(),
		},
	}
}

func keyvaultGet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	r := resource.Item.(*armkeyvault.Vault)
	cl := meta.(*client.Client)
	svc, err := armkeyvault.NewVaultsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*r.ID)
	if err != nil {
		return err
	}
	resp, err := svc.Get(ctx, group, *r.Name, nil)
	if err != nil {
		return err
	}
	resource.SetItem(resp.Vault)
	return nil
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
			res <- &armkeyvault.Vault{
				ID:   r.ID,
				Name: r.Name,
			}
		}
	}
	return nil
}
