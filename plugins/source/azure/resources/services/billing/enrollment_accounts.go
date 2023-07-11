package billing

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func EnrollmentAccounts() *schema.Table {
	return &schema.Table{
		Name:                 "azure_billing_enrollment_accounts",
		Resolver:             fetchEnrollmentAccounts,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling@v0.5.0#EnrollmentAccountSummary",
		Transform:            transformers.TransformWithStruct(&armbilling.EnrollmentAccountSummary{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchEnrollmentAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armbilling.NewEnrollmentAccountsClient(cl.Creds, cl.Options)
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
