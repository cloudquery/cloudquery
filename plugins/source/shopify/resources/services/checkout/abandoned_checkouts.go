package checkout

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AbandonedCheckouts() *schema.Table {
	return &schema.Table{
		Name:      "shopify_abandoned_checkouts",
		Resolver:  fetchAbandonedCheckouts,
		Transform: transformers.TransformWithStruct(&shopify.Checkout{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
