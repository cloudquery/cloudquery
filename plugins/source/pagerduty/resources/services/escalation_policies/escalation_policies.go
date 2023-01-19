package escalation_policies

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EscalationPolicies() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_escalation_policies",
		Description: `https://developer.pagerduty.com/api-reference/51b21014a4f5a-list-escalation-policies`,
		Resolver:    fetchEscalationPolicies,
		Transform:   transformers.TransformWithStruct(&pagerduty.EscalationPolicy{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL")),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
		},
	}
}
