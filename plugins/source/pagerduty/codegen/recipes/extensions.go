package recipes

import "github.com/PagerDuty/go-pagerduty"

func ExtensionResources() []*Resource {
	return []*Resource{
		{
			SubService:  "extensions",
			PKColumns:   []string{"id"},
			Struct:      pagerduty.Extension{},
			Description: "https://developer.pagerduty.com/api-reference/26b46f0092a55-list-extensions",

			ListOptionsStructNameOverride: "ListExtensionOptions",
			ResponseStructOverride:        "ListExtensionResponse",
		},
	}
}
