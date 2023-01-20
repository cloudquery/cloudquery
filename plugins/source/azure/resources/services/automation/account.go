package automation

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Account() *schema.Table {
	return &schema.Table{
		Name:        "azure_automation_account",
		Resolver:    fetchAccount,
		Description: "https://learn.microsoft.com/en-us/rest/api/automation/automation-account/list?tabs=HTTP#automationaccount",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_automation_account", client.Namespacemicrosoft_automation),
		Transform:   transformers.TransformWithStruct(&armautomation.Account{}),
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

func fetchAccount(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armautomation.NewAccountClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
