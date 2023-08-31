package datalakestore

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:                 "azure_datalakestore_accounts",
		Resolver:             fetchAccounts,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/datalakestore/accounts/list?tabs=HTTP#datalakestoreaccountbasic",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_datalakestore_accounts", client.Namespacemicrosoft_datalakestore),
		Transform:            transformers.TransformWithStruct(&armdatalakestore.AccountBasic{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdatalakestore.NewAccountsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
