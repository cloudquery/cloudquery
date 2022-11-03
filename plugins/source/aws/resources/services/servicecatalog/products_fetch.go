package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchServicecatalogProducts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Servicecatalog

	listInput := new(servicecatalog.SearchProductsAsAdminInput)
	for {
		output, err := svc.SearchProductsAsAdmin(ctx, listInput)
		if err != nil {
			return err
		}

		res <- output.ProductViewDetails

		if aws.ToString(output.NextPageToken) == "" {
			break
		}
		listInput.PageToken = output.NextPageToken
	}

	return nil
}

func resolveProductTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.ProductViewDetail)

	cl := meta.(*client.Client)
	svc := cl.Services().Servicecatalogappregistry
	response, err := svc.ListTagsForResource(ctx, &servicecatalogappregistry.ListTagsForResourceInput{
		ResourceArn: p.ProductARN,
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, response.Tags)
}
