package recipes

import "github.com/PagerDuty/go-pagerduty"

func VendorResources() []*Resource {
	return []*Resource{
		{
			SubService:  "vendors",
			PKColumns:   []string{"id"},
			Struct:      pagerduty.Vendor{},
			Description: "https://developer.pagerduty.com/api-reference/d2aa663abec79-list-vendors",

			ListOptionsStructNameOverride: "ListVendorOptions",
			ResponseStructOverride:        "ListVendorResponse",
		},
	}
}
