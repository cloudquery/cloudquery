package databases

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func firewallRules() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_database_firewall_rules",
		Resolver:  fetchDatabasesFirewallRules,
		Transform: transformers.TransformWithStruct(&godo.DatabaseFirewallRule{}),
		Columns:   []schema.Column{},
	}
}
