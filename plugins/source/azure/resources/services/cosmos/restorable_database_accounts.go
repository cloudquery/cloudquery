package cosmos

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RestorableDatabaseAccounts() *schema.Table {
	return &schema.Table{
		Name:        "azure_cosmos_restorable_database_accounts",
		Resolver:    fetchRestorableDatabaseAccounts,
		Description: "https://learn.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2022-05-15/restorable-database-accounts/list?tabs=HTTP#restorabledatabaseaccountgetresult",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_cosmos_restorable_database_accounts", client.Namespacemicrosoft_documentdb),
		Transform:   transformers.TransformWithStruct(&armcosmos.RestorableDatabaseAccountGetResult{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
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
