package cosmos

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DatabaseAccounts() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cosmos_database_accounts",
		Resolver:             fetchDatabaseAccounts,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2022-05-15/database-accounts/list?tabs=HTTP#databaseaccountgetresults",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_cosmos_database_accounts", client.Namespacemicrosoft_documentdb),
		Transform:            transformers.TransformWithStruct(&armcosmos.DatabaseAccountGetResults{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},

		Relations: []*schema.Table{
			mongo_db_databases(),
			sql_databases(),
		},
	}
}

func fetchDatabaseAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcosmos.NewDatabaseAccountsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
