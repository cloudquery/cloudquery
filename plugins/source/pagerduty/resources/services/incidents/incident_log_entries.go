package incidents

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func IncidentLogEntries() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_incident_log_entries",
		Description: `https://developer.pagerduty.com/api-reference/367602cbc1c28-list-log-entries-for-an-incident`,
		Resolver:    fetchIncidentLogEntries,
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
			{
				Name:     "agent",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Agent"),
			},
			{
				Name:     "channel",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Channel"),
			},
			{
				Name:     "teams",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Teams"),
			},
			{
				Name:     "contexts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Contexts"),
			},
			{
				Name:     "acknowledgement_timeout",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AcknowledgementTimeout"),
			},
			{
				Name:     "event_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EventDetails"),
			},
			{
				Name:     "assignees",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Assignees"),
			},
			{
				Name:     "incident",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Incident"),
			},
			{
				Name:     "service",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Service"),
			},
			{
				Name:     "user",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("User"),
			},
		},
	}
}
