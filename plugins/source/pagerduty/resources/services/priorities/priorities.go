package priorities

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Priorities() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_priorities",
		Description: `https://developer.pagerduty.com/api-reference/0fa9ad52bf2d2-list-priorities`,
		Resolver:    fetchPriorities,
		Transform:   transformers.TransformWithStruct(&pagerduty.Priority{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL")),
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
