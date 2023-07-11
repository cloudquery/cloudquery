package escalation_policies

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func EscalationPolicies() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_escalation_policies",
		Description: `https://developer.pagerduty.com/api-reference/51b21014a4f5a-list-escalation-policies`,
		Resolver:    fetchEscalationPolicies,
		Transform:   transformers.TransformWithStruct(&pagerduty.EscalationPolicy{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL")),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "html_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("HTMLURL"),
			},
		},
	}
}
