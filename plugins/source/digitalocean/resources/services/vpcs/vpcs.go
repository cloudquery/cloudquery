package vpcs

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			members(),
		},
	}
}
