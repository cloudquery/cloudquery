package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func serverFirewallRules() *schema.Table {
	return &schema.Table{
		Name:                 "azure_sql_server_firewall_rules",
		Resolver:             fetchFirewallRules,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/firewall-rules/list-by-server?tabs=HTTP#firewallrule",
		Transform:            transformers.TransformWithStruct(&armsql.FirewallRule{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchFirewallRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*armsql.Server)
	cl := meta.(*client.Client)
	svc, err := armsql.NewFirewallRulesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
