package extension_schemas

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ExtensionSchemas() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_extension_schemas",
		Description: `https://developer.pagerduty.com/api-reference/6eef27c5b452f-list-extension-schemas`,
		Resolver:    fetchExtensionSchemas,
		Transform:   transformers.TransformWithStruct(&pagerduty.ExtensionSchema{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL", "URL", "IconURL", "LogoURL")),
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
				Name:     "icon_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("IconURL"),
			},
			{
				Name:     "logo_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("LogoURL"),
			},
			{
				Name:     "guide_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("GuideURL"),
			},
			{
				Name:     "url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("URL"),
			},
		},
	}
}
