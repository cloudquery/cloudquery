package vpcs

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/digitalocean/godo"
)

func Vpcs() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_vpcs",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/VPCs",
		Resolver:    fetchVpcsVpcs,
		Transform:   transformers.TransformWithStruct(&godo.VPC{}),
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
