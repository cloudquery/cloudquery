package postgresql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func firewall_rules() *schema.Table {
	return &schema.Table{
		Name:        "azure_postgresql_server_firewall_rules",
		Resolver:    fetchFirewallRules,
		Description: "https://learn.microsoft.com/en-us/rest/api/postgresql/singleserver/firewall-rules/list-by-server?tabs=HTTP#firewallrule",
		Transform:   transformers.TransformWithStruct(&armpostgresql.FirewallRule{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{},
	}
}
