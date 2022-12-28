package recipes

import "github.com/PagerDuty/go-pagerduty"

func RulesetResources() []*Resource {
	return []*Resource{
		{
			SubService:  "rulesets",
			PKColumns:   []string{"id"},
			Struct:      pagerduty.Ruleset{},
			Relations:   []string{"RulesetRules()"},
			Description: "https://developer.pagerduty.com/api-reference/633f1ecb6c03b-list-rulesets",

			Template: "paginated_no_options",
		},
		{
			SubService:  "ruleset_rules",
			PKColumns:   []string{"id"},
			Struct:      pagerduty.RulesetRule{},
			Description: "https://developer.pagerduty.com/api-reference/c39605f86c5b7-list-event-rules",

			Template: "nested_paginated_no_options",

			ResponseFieldOverride: "Rules",
			RestPathOverride:      "/rules",
			ParentIsPointer:       true,
		},
	}
}
