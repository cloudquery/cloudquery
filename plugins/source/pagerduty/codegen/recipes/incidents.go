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
			},
			Template: "basic",
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
	}
}
