package priorities

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Priorities() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_priorities",
		Description: `https://developer.pagerduty.com/api-reference/0fa9ad52bf2d2-list-priorities`,
		Resolver:    fetchPriorities,
		Transform:   transformers.TransformWithStruct(&pagerduty.Priority{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL")),
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
		},
	}
}
