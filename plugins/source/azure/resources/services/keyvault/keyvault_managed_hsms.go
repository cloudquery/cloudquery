package keyvault

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func KeyvaultManagedHsms() *schema.Table {
	return &schema.Table{
		Name:                 "azure_keyvault_keyvault_managed_hsms",
		Resolver:             fetchKeyvaultManagedHsms,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/keyvault/managedhsm/managed-hsms/list-by-subscription?tabs=HTTP#managedhsm",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_keyvault_keyvault_managed_hsms", client.Namespacemicrosoft_keyvault),
		Transform:            transformers.TransformWithStruct(&armkeyvault.ManagedHsm{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
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
