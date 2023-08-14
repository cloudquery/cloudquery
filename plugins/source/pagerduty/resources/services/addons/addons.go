package addons

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Addons() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_addons",
		Description: `https://developer.pagerduty.com/api-reference/e58b140202a57-list-installed-add-ons`,
		Resolver:    fetchAddons,
		Transform:   transformers.TransformWithStruct(&pagerduty.Addon{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
