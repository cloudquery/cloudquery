package incidents

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func IncidentLogEntries() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_incident_log_entries",
		Description: `https://developer.pagerduty.com/api-reference/367602cbc1c28-list-log-entries-for-an-incident`,
		Resolver:    fetchIncidentLogEntries,
		Transform:   transformers.TransformWithStruct(&pagerduty.LogEntry{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL", "APIObject")),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "type",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "summary",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Summary"),
			},
			{
				Name:     "self",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Self"),
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
		},
	}
}
