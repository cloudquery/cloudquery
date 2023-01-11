package extension_schemas

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ExtensionSchemas() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_extension_schemas",
		Description: `https://developer.pagerduty.com/api-reference/6eef27c5b452f-list-extension-schemas`,
		Resolver:    fetchExtensionSchemas,
		Transform:   transformers.TransformWithStruct(&pagerduty.ExtensionSchema{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL", "URL", "IconURL", "LogoURL")),
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
				Name:     "icon_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IconURL"),
			},
			{
				Name:     "logo_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogoURL"),
			},
			{
				Name:     "guide_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GuideURL"),
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URL"),
			},
		},
	}
}
