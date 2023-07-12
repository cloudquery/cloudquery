package rulesets

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Rulesets() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_rulesets",
		Description: `https://developer.pagerduty.com/api-reference/633f1ecb6c03b-list-rulesets`,
		Resolver:    fetchRulesets,
		Transform:   transformers.TransformWithStruct(&pagerduty.Ruleset{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "created_at",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: schema.PathResolver("CreatedAt"),
			},
		},

		Relations: []*schema.Table{
			RulesetRules(),
		},
	}
}
