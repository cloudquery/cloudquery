package incidents

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func IncidentNotes() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_incident_notes",
		Description: `https://developer.pagerduty.com/api-reference/a1ac30885eb7a-list-notes-for-an-incident`,
		Resolver:    fetchIncidentNotes,
		Transform:   transformers.TransformWithStruct(&pagerduty.IncidentNote{}),
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
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
		},
	}
}
