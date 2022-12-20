package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchServicecatalogProvisionedProducts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Servicecatalog

	listInput := new(servicecatalog.SearchProvisionedProductsInput)
	for {
		output, err := svc.SearchProvisionedProducts(ctx, listInput)
		if err != nil {
			return err
		}

		res <- output.ProvisionedProducts

		if aws.ToString(output.NextPageToken) == "" {
			break
		}
		listInput.PageToken = output.NextPageToken
	}

	return nil
}

func resolveProvisionedProductTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.ProvisionedProductAttribute)
	return resource.Set(c.Name, client.TagsToMap(p.Tags))
}
