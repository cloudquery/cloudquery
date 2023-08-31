package checkout

import (
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AbandonedCheckouts() *schema.Table {
	return &schema.Table{
		Name:      "shopify_abandoned_checkouts",
		Resolver:  fetchAbandonedCheckouts,
		Transform: transformers.TransformWithStruct(&shopify.Checkout{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:           "updated_at",
				Type:           arrow.FixedWidthTypes.Timestamp_us,
				Resolver:       schema.PathResolver("UpdatedAt"),
				IncrementalKey: true,
			},
		},
		IsIncremental: true,
	}
}
