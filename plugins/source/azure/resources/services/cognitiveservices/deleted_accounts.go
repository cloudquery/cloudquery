package cognitiveservices

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DeletedAccounts() *schema.Table {
	return &schema.Table{
		Name:        "azure_cognitiveservices_deleted_accounts",
		Resolver:    fetchDeletedAccounts,
		Description: "https://learn.microsoft.com/en-us/rest/api/cognitiveservices/accountmanagement/deleted-accounts/list?tabs=HTTP#account",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_cognitiveservices_deleted_accounts", client.Namespacemicrosoft_cognitiveservices),
		Transform:   transformers.TransformWithStruct(&armcognitiveservices.Account{}),
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

func fetchDeletedAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcognitiveservices.NewDeletedAccountsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
