package extensions

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Extensions() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_extensions",
		Description: `https://developer.pagerduty.com/api-reference/26b46f0092a55-list-extensions`,
		Resolver:    fetchExtensions,
		Transform:   transformers.TransformWithStruct(&pagerduty.Extension{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL", "EndpointURL")),
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
				Name:     "endpoint_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("EndpointURL"),
			},
		},
	}
}
