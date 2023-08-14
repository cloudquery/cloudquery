package services

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func products() *schema.Table {
	return &schema.Table{
		Name:        "awspricing_service_products",
		Title:       "Service Products from the AWS Price List API",
		Description: "https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/price-changes.html",
		Resolver:    fetchProducts,
		Transform:   transformers.TransformWithStruct(&Product{}, transformers.WithPrimaryKeys("Sku")),
	}
}

func fetchProducts(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	pricingFile := parent.Item.(PricingFile)
	res <- pricingFile.Products
	return nil
}
