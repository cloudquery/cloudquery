package recipes

import "github.com/PagerDuty/go-pagerduty"

func ScheduleResources() []*Resource {
	return []*Resource{
		{
			SubService:  "schedules",
			PKColumns:   []string{"id"},
			Description: "https://developer.pagerduty.com/api-reference/846ecf84402bb-list-schedules",
			Struct:      pagerduty.Schedule{},
		},
	}
}
