package recipes

import "github.com/PagerDuty/go-pagerduty"

func PriorityResources() []*Resource {
	return []*Resource{
		{
			SubService:  "priorities",
			Description: "https://developer.pagerduty.com/api-reference/0fa9ad52bf2d2-list-priorities",
			PKColumns:   []string{"id"},
			Struct:      pagerduty.Priority{},

			// overrides because `Prioritys` is not a valid pluralization of `Priority`.
			ListFunctionNameOverride:      "ListPrioritiesWithContext",
			ListOptionsStructNameOverride: "ListPrioritiesOptions",
			ResponseStructOverride:        "ListPrioritiesResponse",
			ResponseFieldOverride:         "Priorities",
			RestPathOverride:              "/priorities",
		},
	}
}
