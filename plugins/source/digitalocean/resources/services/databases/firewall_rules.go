package databases

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func firewallRules() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_database_firewall_rules",
		Description: "https://pkg.go.dev/github.com/digitalocean/godo#DatabaseFirewallRule",
		Resolver:    fetchDatabasesFirewallRules,
		Transform:   transformers.TransformWithStruct(&godo.DatabaseFirewallRule{}),
		Columns:     []schema.Column{},
	}
}
