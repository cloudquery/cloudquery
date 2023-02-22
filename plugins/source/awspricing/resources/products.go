package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func products() *schema.Table {
	return &schema.Table{
		Name:      "awspricing_service_products",
		Resolver:  fetchProducts,
		Transform: transformers.TransformWithStruct(&Product{}, transformers.WithPrimaryKeys("Sku")),
	}
}

func fetchProducts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	pricingFile := parent.Item.(PricingFile)
	res <- pricingFile.Products
	return nil
}
