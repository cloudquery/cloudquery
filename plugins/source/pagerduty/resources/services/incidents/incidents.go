// Code generated by codegen; DO NOT EDIT.

package incidents

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Incidents() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_incidents",
		Description: `https://developer.pagerduty.com/api-reference/9d0b4b12e36f9-list-incidents`,
		Resolver:    fetchIncidents,
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
				Name:     "incident_number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("IncidentNumber"),
			},
			{
				Name:     "title",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Title"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "pending_actions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingActions"),
			},
			{
				Name:     "incident_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IncidentKey"),
			},
			{
				Name:     "service",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Service"),
			},
			{
				Name:     "assignments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Assignments"),
			},
			{
				Name:     "acknowledgements",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Acknowledgements"),
			},
			{
				Name:     "last_status_change_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastStatusChangeAt"),
			},
			{
				Name:     "last_status_change_by",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastStatusChangeBy"),
			},
			{
				Name:     "first_trigger_log_entry",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FirstTriggerLogEntry"),
			},
			{
				Name:     "escalation_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EscalationPolicy"),
			},
			{
				Name:     "teams",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Teams"),
			},
			{
				Name:     "priority",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Priority"),
			},
			{
				Name:     "urgency",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Urgency"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "resolve_reason",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResolveReason"),
			},
			{
				Name:     "alert_counts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AlertCounts"),
			},
			{
				Name:     "body",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Body"),
			},
			{
				Name:     "is_mergeable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsMergeable"),
			},
			{
				Name:     "conference_bridge",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConferenceBridge"),
			},
			{
				Name:     "assigned_via",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssignedVia"),
			},
			{
				Name:     "occurrence",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Occurrence"),
			},
			{
				Name:     "incidents_responders",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IncidentResponders"),
			},
			{
				Name:     "responder_requests",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResponderRequests"),
			},
		},

		Relations: []*schema.Table{
			IncidentAlerts(),
			IncidentNotes(),
			IncidentLogEntries(),
		},
	}
}
