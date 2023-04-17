package rulesets

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Rulesets() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_rulesets",
		Description: `https://developer.pagerduty.com/api-reference/633f1ecb6c03b-list-rulesets`,
		Resolver:    fetchRulesets,
		Transform:   transformers.TransformWithStruct(&pagerduty.Ruleset{}),
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
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
		},

		Relations: []*schema.Table{
			RulesetRules(),
		},
	}
}
