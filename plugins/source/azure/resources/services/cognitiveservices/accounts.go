package cognitiveservices

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:        "azure_cognitiveservices_accounts",
		Resolver:    fetchAccounts,
		Description: "https://learn.microsoft.com/en-us/rest/api/cognitiveservices/accountmanagement/accounts/list?tabs=HTTP#account",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_cognitiveservices_accounts", client.Namespacemicrosoft_cognitiveservices),
		Transform:   transformers.TransformWithStruct(&armcognitiveservices.Account{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcognitiveservices.NewAccountsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
