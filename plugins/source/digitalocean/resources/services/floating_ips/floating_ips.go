package floating_ips

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/digitalocean/godo"
)

func FloatingIps() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_floating_ips",
		Description: "Deprecated. https://docs.digitalocean.com/reference/api/api-reference/#tag/Floating-IPs",
		Resolver:    fetchFloatingIpsFloatingIps,
		Transform:   transformers.TransformWithStruct(&godo.FloatingIP{}),
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
