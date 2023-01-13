package tags

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Tags() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_tags",
		Description: `https://developer.pagerduty.com/api-reference/e44b160c69bf3-list-tags`,
		Resolver:    fetchTags,
		Transform:   transformers.TransformWithStruct(&pagerduty.Tag{}, transformers.WithSkipFields("HTMLURL"), transformers.WithUnwrapAllEmbeddedStructs()),
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
