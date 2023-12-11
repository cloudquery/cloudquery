package oncalls

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Oncalls() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_oncalls",
		Description: `https://developer.pagerduty.com/api-reference/3a6b910f11050-list-all-of-the-on-calls`,
		Resolver:    fetchOncalls,
		Transform:   transformers.TransformWithStruct(&pagerduty.OnCall{}, transformers.WithUnwrapAllEmbeddedStructs()),
	}
}
