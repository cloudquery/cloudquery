package databases

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func firewallRules() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_database_firewall_rules",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/databases_list_firewall_rules",
		Resolver:    fetchDatabasesFirewallRules,
		Transform:   transformers.TransformWithStruct(&godo.DatabaseFirewallRule{}),
		Columns:     []schema.Column{},
	}
}
