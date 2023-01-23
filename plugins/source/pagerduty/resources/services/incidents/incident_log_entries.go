package incidents

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func IncidentLogEntries() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_incident_log_entries",
		Description: `https://developer.pagerduty.com/api-reference/367602cbc1c28-list-log-entries-for-an-incident`,
		Resolver:    fetchIncidentLogEntries,
		Transform:   transformers.TransformWithStruct(&pagerduty.LogEntry{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL", "APIObject")),
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
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "summary",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Summary"),
			},
			{
				Name:     "self",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Self"),
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
		},
	}
}
