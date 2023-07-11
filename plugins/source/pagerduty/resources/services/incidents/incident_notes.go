package incidents

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func IncidentNotes() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_incident_notes",
		Description: `https://developer.pagerduty.com/api-reference/a1ac30885eb7a-list-notes-for-an-incident`,
		Resolver:    fetchIncidentNotes,
		Transform:   transformers.TransformWithStruct(&pagerduty.IncidentNote{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "created_at",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: schema.PathResolver("CreatedAt"),
			},
		},
	}
}
