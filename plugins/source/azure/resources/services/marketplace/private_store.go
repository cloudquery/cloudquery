package marketplace

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PrivateStore() *schema.Table {
	return &schema.Table{
		Name:                 "azure_marketplace_private_store",
		Resolver:             fetchPrivateStore,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace@v1.0.0#PrivateStore",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_marketplace_private_store", client.Namespacemicrosoft_marketplace),
		Transform:            transformers.TransformWithStruct(&armmarketplace.PrivateStore{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchPrivateStore(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armmarketplace.NewPrivateStoreClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
