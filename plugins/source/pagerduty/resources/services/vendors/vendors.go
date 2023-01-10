package vendors

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Vendors() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_vendors",
		Description: `https://developer.pagerduty.com/api-reference/d2aa663abec79-list-vendors`,
		Resolver:    fetchVendors,
		Transform:   transformers.TransformWithStruct(&pagerduty.Vendor{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL", "LogoURL", "WebsiteURL", "ThumbnailURL", "IsPDCEF")),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "logo_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogoURL"),
			},
			{
				Name:     "website_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebsiteURL"),
			},
			{
				Name:     "thumbnail_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThumbnailURL"),
			},
			{
				Name:     "is_pd_cef",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsPDCEF"),
			},
		},
	}
}
