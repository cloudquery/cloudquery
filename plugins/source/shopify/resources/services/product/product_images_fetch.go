package product

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchProductImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(shopify.Product)

	res <- p.Images
	return nil
}
