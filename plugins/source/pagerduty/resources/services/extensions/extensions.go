package extensions

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Extensions() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_extensions",
		Description: `https://developer.pagerduty.com/api-reference/26b46f0092a55-list-extensions`,
		Resolver:    fetchExtensions,
		Transform:   transformers.TransformWithStruct(&pagerduty.Extension{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL", "EndpointURL")),
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
				Name:     "endpoint_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EndpointURL"),
			},
		},
	}
}
