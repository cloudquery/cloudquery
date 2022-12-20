// This file can't be called "maintenance_windows", because golang thinks it should only be built on windows. Yuck!
package recipes

import "github.com/PagerDuty/go-pagerduty"

func MaintenanceWindowResources() []*Resource {
	return []*Resource{
		{
			SubService:  "maintenance_windows",
			PKColumns:   []string{"id"},
			Struct:      pagerduty.MaintenanceWindow{},
			Description: "https://developer.pagerduty.com/api-reference/4c0936c241cbb-list-maintenance-windows",

			ResponseFieldOverride: "MaintenanceWindows",
		},
	}
}
