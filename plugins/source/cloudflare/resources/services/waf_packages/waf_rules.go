package waf_packages

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func WAFRules() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_waf_rules",
		Resolver:  fetchWAFRules,
		Transform: transformers.TransformWithStruct(&cloudflare.WAFRule{}),
		Columns: []schema.Column{
			{
				Name:     "waf_package_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
