package billing

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:                 "azure_billing_accounts",
		Resolver:             fetchAccounts,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/billing/2020-05-01/billing-accounts/list?tabs=HTTP#billingaccount",
		Transform:            transformers.TransformWithStruct(&armbilling.Account{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	// we already fetch all accounts during initialization, so no need to fetch them again
	res <- meta.(*client.Client).BillingAccounts
	return nil
}
