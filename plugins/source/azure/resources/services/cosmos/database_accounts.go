package cosmos

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DatabaseAccounts() *schema.Table {
	return &schema.Table{
		Name:      "azure_cosmos_database_accounts",
		Resolver:  fetchDatabaseAccounts,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_cosmos_database_accounts", client.Namespacemicrosoft_documentdb),
		Transform: transformers.TransformWithStruct(&armcosmos.DatabaseAccountGetResults{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

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
