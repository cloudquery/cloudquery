package incidents

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Incidents() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_incidents",
		Description: `https://developer.pagerduty.com/api-reference/9d0b4b12e36f9-list-incidents`,
		Resolver:    fetchIncidents,
		Transform:   transformers.TransformWithStruct(&pagerduty.Incident{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL")),
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
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "last_status_change_at",
				Type:     schema.TypeTimestamp,
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
