package schedules

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Schedules() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_schedules",
		Description: `https://developer.pagerduty.com/api-reference/846ecf84402bb-list-schedules`,
		Resolver:    fetchSchedules,
		Transform:   transformers.TransformWithStruct(&pagerduty.Schedule{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL")),
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
