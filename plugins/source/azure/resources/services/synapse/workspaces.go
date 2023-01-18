package synapse

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Workspaces() *schema.Table {
	return &schema.Table{
		Name:        "azure_synapse_workspaces",
		Resolver:    fetchWorkspaces,
		Description: "https://learn.microsoft.com/en-us/rest/api/synapse/workspaces/list?tabs=HTTP#workspace",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_synapse_workspaces", client.Namespacemicrosoft_synapse),
		Transform:   transformers.TransformWithStruct(&armsynapse.Workspace{}),
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

func fetchWorkspaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsynapse.NewWorkspacesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
