package firewalls

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func Firewalls() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_firewalls",
		Resolver:  fetchFirewallsFirewalls,
		Transform: transformers.TransformWithStruct(&godo.Firewall{}),
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
