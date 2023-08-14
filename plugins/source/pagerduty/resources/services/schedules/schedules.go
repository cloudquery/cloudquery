package schedules

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Schedules() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_schedules",
		Description: `https://developer.pagerduty.com/api-reference/846ecf84402bb-list-schedules`,
		Resolver:    fetchSchedules,
		Transform:   transformers.TransformWithStruct(&pagerduty.Schedule{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL")),
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
