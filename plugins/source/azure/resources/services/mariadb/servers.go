package mariadb

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Servers() *schema.Table {
	return &schema.Table{
		Name:      "azure_mariadb_servers",
		Resolver:  fetchServers,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_mariadb_servers", client.Namespacemicrosoft_dbformariadb),
		Transform: transformers.TransformWithStruct(&armmariadb.Server{}),
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

		Relations: []*schema.Table{
			server_configurations(),
		},
	}
}

func fetchServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armmariadb.NewServersClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
