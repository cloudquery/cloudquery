package vpcs

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func Vpcs() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_vpcs",
		Resolver:  fetchVpcsVpcs,
		Transform: transformers.TransformWithStruct(&godo.VPC{}),
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

		Relations: []*schema.Table{
			members(),
		},
	}
}
