package rulesets

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func RulesetRules() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_ruleset_rules",
		Description: `https://developer.pagerduty.com/api-reference/c39605f86c5b7-list-event-rules`,
		Resolver:    fetchRulesetRules,
		Transform:   transformers.TransformWithStruct(&pagerduty.RulesetRule{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
