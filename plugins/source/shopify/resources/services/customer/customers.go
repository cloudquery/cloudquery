package customer

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func Customers() *schema.Table {
	return &schema.Table{
		Name:      "shopify_customers",
		Resolver:  fetchCustomers,
		Transform: client.TransformWithStruct(&shopify.Customer{}),
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
