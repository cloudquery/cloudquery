package vpcs

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func members() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_vpc_members",
		Resolver:  fetchVpcsMembers,
		Transform: transformers.TransformWithStruct(&godo.VPCMember{}),
		Columns: []schema.Column{
			{
				Name:     "urn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
