package datalakeanalytics

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:                 "azure_datalakeanalytics_accounts",
		Resolver:             fetchAccounts,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/datalakeanalytics/accounts/list?tabs=HTTP#datalakeanalyticsaccountbasic",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_datalakeanalytics_accounts", client.Namespacemicrosoft_datalakeanalytics),
		Transform:            transformers.TransformWithStruct(&armdatalakeanalytics.AccountBasic{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdatalakeanalytics.NewAccountsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
