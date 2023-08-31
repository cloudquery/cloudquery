package cosmos

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func RestorableDatabaseAccounts() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cosmos_restorable_database_accounts",
		Resolver:             fetchRestorableDatabaseAccounts,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2022-05-15/restorable-database-accounts/list?tabs=HTTP#restorabledatabaseaccountgetresult",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_cosmos_restorable_database_accounts", client.Namespacemicrosoft_documentdb),
		Transform:            transformers.TransformWithStruct(&armcosmos.RestorableDatabaseAccountGetResult{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchRestorableDatabaseAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcosmos.NewRestorableDatabaseAccountsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
