package billing

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:        "azure_billing_accounts",
		Resolver:    fetchAccounts,
		Description: "https://learn.microsoft.com/en-us/rest/api/billing/2020-05-01/billing-accounts/list?tabs=HTTP#billingaccount",
		Transform:   transformers.TransformWithStruct(&armbilling.Account{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armbilling.NewAccountsClient(cl.Creds, cl.Options)
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
