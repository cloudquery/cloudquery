package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func virtualNetworkRules() *schema.Table {
	return &schema.Table{
		Name:                 "azure_sql_server_virtual_network_rules",
		Resolver:             fetchVirtualNetworkRules,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/virtual-network-rules/list-by-server?tabs=HTTP#virtualnetworkrule",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_sql_virtual_network_rules", client.Namespacemicrosoft_sql),
		Transform:            transformers.TransformWithStruct(&armsql.VirtualNetworkRule{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
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
