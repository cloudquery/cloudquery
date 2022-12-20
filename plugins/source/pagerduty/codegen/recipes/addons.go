package recipes

import "github.com/PagerDuty/go-pagerduty"

func AddonResources() []*Resource {
	return []*Resource{
		{
			SubService:  "addons",
			Struct:      pagerduty.Addon{},
			PKColumns:   []string{"id"},
			Description: "https://developer.pagerduty.com/api-reference/e58b140202a57-list-installed-add-ons",

			ListOptionsStructNameOverride: "ListAddonOptions",
			ResponseStructOverride:        "ListAddonResponse",
			Template:                      "basic",
		},
	}
}
