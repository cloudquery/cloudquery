package recipes

import "github.com/PagerDuty/go-pagerduty"

func EscalationPolicyResources() []*Resource {
	return []*Resource{
		{
			SubService:  "escalation_policies",
			Description: "https://developer.pagerduty.com/api-reference/51b21014a4f5a-list-escalation-policies",
			Struct:      pagerduty.EscalationPolicy{},
			PKColumns:   []string{"id"},

			ListFunctionNameOverride:      "ListEscalationPoliciesWithContext",
			ListOptionsStructNameOverride: "ListEscalationPoliciesOptions",
			ResponseFieldOverride:         "EscalationPolicies",
			ResponseStructOverride:        "ListEscalationPoliciesResponse",
			RestPathOverride:              "/escalation_policies",
		},
	}
}
