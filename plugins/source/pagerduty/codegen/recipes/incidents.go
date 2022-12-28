package recipes

import "github.com/PagerDuty/go-pagerduty"

func IncidentResources() []*Resource {
	return []*Resource{
		{
			SubService:  "incidents",
			Description: "https://developer.pagerduty.com/api-reference/9d0b4b12e36f9-list-incidents",
			Struct:      pagerduty.Incident{},
			PKColumns:   []string{"id"},
			Relations: []string{
				"IncidentAlerts()",
				"IncidentNotes()",
				"IncidentLogEntries()",
			},
			Template: "basic",

			SkipMockGeneration: true,
		},
		{
			SubService:  "incident_alerts",
			Struct:      pagerduty.IncidentAlert{},
			Description: "https://developer.pagerduty.com/api-reference/4bc42e7ac0c59-list-alerts-for-an-incident",
			PKColumns:   []string{"id"},
			Template:    "nested",

			ResponseStructOverride: "ListAlertsResponse",
			ResponseFieldOverride:  "Alerts",
			RestPathOverride:       "/alerts",
		},
		{
			SubService:  "incident_notes",
			Struct:      pagerduty.IncidentNote{},
			Description: "https://developer.pagerduty.com/api-reference/a1ac30885eb7a-list-notes-for-an-incident",
			PKColumns:   []string{"id"},
			Template:    "nested_paginated_no_options",

			ListFunctionNameOverride: "ListIncidentNotesWithContext",
		},
		// Incident-log-entries is manually-generated
		// {
		// 	SubService:                       "incident_log_entries",
		// 	Struct:                           pagerduty.LogEntry{},
		// 	Description:                      "https://developer.pagerduty.com/api-reference/367602cbc1c28-list-log-entries-for-an-incident",
		// 	PKColumns:                        []string{"id"},
		// 	UnwrapEmbeddedStructsRecursively: true,

		// 	Template:                      "nested",
		// 	ListFunctionNameOverride:      "ListIncidentLogEntriesWithContext",
		// 	ListOptionsStructNameOverride: "ListIncidentLogEntriesOptions",
		// 	ResponseFieldOverride:         "LogEntries",
		// },
	}
}
