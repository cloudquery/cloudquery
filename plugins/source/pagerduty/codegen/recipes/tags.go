package recipes

import "github.com/PagerDuty/go-pagerduty"

func TagsResources() []*Resource {
	return []*Resource{
		{
			SubService:  "tags",
			PKColumns:   []string{"id"},
			Struct:      pagerduty.Tag{},
			Description: "https://developer.pagerduty.com/api-reference/e44b160c69bf3-list-tags",

			Template: "paginated",

			ResponseStructOverride: "ListTagResponse",
		},
	}
}
