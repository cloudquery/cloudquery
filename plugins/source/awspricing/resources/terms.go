package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func terms() *schema.Table {
	return &schema.Table{
		Name:        "awspricing_service_terms",
		Title:       "Service Terms from the AWS Price List API",
		Description: "https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/price-changes.html",
		Resolver:    fetchTerms,
		Transform:   transformers.TransformWithStruct(&Term{}, transformers.WithPrimaryKeys("OfferTermCode", "Sku")),
	}
}
func fetchTerms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	pricingFile := parent.Item.(PricingFile)
	res <- pricingFile.Terms
	return nil
}
