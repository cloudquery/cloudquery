package firewalls

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
