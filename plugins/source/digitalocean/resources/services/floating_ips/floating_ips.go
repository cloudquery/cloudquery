package floating_ips

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func FloatingIps() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_floating_ips",
		Resolver:  fetchFloatingIpsFloatingIps,
		Transform: transformers.TransformWithStruct(&godo.FloatingIP{}),
		Columns: []schema.Column{
			{
				Name:     "ip",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IP"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
