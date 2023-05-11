package firewalls

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/digitalocean/godo"
)

func Firewalls() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_firewalls",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/Firewalls",
		Resolver:    fetchFirewallsFirewalls,
		Transform:   transformers.TransformWithStruct(&godo.Firewall{}),
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
