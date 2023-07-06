package floating_ips

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "ip",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("IP"),
				PrimaryKey: true,
			},
		},
	}
}
