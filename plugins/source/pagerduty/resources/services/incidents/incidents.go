package incidents

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Incidents() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_incidents",
		Description: `https://developer.pagerduty.com/api-reference/9d0b4b12e36f9-list-incidents`,
		Resolver:    fetchIncidents,
		Transform:   transformers.TransformWithStruct(&pagerduty.Incident{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL")),
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
				Name:     "created_at",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "last_status_change_at",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: schema.PathResolver("LastStatusChangeAt"),
			},
		},

		Relations: []*schema.Table{
			IncidentAlerts(),
			IncidentNotes(),
			IncidentLogEntries(),
		},
	}
}
