package sql

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func virtualNetworkRules() *schema.Table {
	return &schema.Table{
		Name:      "azure_sql_server_virtual_network_rules",
		Resolver:  fetchVirtualNetworkRules,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_sql_virtual_network_rules", client.Namespacemicrosoft_sql),
		Transform: transformers.TransformWithStruct(&armsql.VirtualNetworkRule{}),
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
	}
}

func fetchVirtualNetworkRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armsql.Server)
	cl := meta.(*client.Client)
	svc, err := armsql.NewVirtualNetworkRulesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*p.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListByServerPager(group, *p.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
