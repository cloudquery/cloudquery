package waf_packages

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func WAFGroups() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_waf_groups",
		Resolver:  fetchWAFGroups,
		Transform: client.TransformWithStruct(&cloudflare.WAFGroup{}),
		Columns: []schema.Column{
			{
				Name:     "waf_package_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
