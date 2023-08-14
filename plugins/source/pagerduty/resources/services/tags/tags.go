package tags

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Tags() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_tags",
		Description: `https://developer.pagerduty.com/api-reference/e44b160c69bf3-list-tags`,
		Resolver:    fetchTags,
		Transform:   transformers.TransformWithStruct(&pagerduty.Tag{}, transformers.WithSkipFields("HTMLURL"), transformers.WithUnwrapAllEmbeddedStructs()),
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
		},
	}
}
