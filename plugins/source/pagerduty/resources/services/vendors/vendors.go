package vendors

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Vendors() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_vendors",
		Description: `https://developer.pagerduty.com/api-reference/d2aa663abec79-list-vendors`,
		Resolver:    fetchVendors,
		Transform:   transformers.TransformWithStruct(&pagerduty.Vendor{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL", "LogoURL", "WebsiteURL", "ThumbnailURL", "IsPDCEF")),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "html_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "logo_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("LogoURL"),
			},
			{
				Name:     "website_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("WebsiteURL"),
			},
			{
				Name:     "thumbnail_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("ThumbnailURL"),
			},
			{
				Name:     "is_pd_cef",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("IsPDCEF"),
			},
		},
	}
}
