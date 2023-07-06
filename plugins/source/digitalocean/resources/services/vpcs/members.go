package vpcs

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func members() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_vpc_members",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/vpcs_list_members",
		Resolver:    fetchVpcsMembers,
		Transform:   transformers.TransformWithStruct(&godo.VPCMember{}),
		Columns: []schema.Column{
			{
				Name:       "urn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("URN"),
				PrimaryKey: true,
			},
		},
	}
}
