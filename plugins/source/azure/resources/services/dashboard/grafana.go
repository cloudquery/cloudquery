package dashboard

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Grafana() *schema.Table {
	return &schema.Table{
		Name:        "azure_dashboard_grafana",
		Resolver:    fetchGrafana,
		Description: "https://learn.microsoft.com/en-us/rest/api/managed-grafana/grafana/list-by-resource-group?tabs=HTTP#managedgrafana",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_dashboard_grafana", client.Namespacemicrosoft_dashboard),
		Transform:   transformers.TransformWithStruct(&armdashboard.ManagedGrafana{}),
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

func fetchGrafana(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdashboard.NewGrafanaClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
