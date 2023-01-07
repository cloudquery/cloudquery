package addons

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Addons() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_addons",
		Description: `https://developer.pagerduty.com/api-reference/e58b140202a57-list-installed-add-ons`,
		Resolver:    fetchAddons,
		Transform:   transformers.TransformWithStruct(&pagerduty.Addon{}, transformers.WithUnwrapAllEmbeddedStructs()),
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
