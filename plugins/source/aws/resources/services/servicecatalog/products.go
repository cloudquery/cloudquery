package servicecatalog

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Products() *schema.Table {
	tableName := "aws_servicecatalog_products"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProductViewDetail.html`,
		Resolver:    fetchServicecatalogProducts,
		Transform:   transformers.TransformWithStruct(&types.ProductViewDetail{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "servicecatalog"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ProductARN"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveProductTags,
			},
		},
	}
}

func fetchServicecatalogProducts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Servicecatalog

	listInput := new(servicecatalog.SearchProductsAsAdminInput)
	paginator := servicecatalog.NewSearchProductsAsAdminPaginator(svc, listInput)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *servicecatalog.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ProductViewDetails
	}

	return nil
}

func resolveProductTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.ProductViewDetail)

	cl := meta.(*client.Client)
	svc := cl.Services().Servicecatalogappregistry
	response, err := svc.ListTagsForResource(ctx, &servicecatalogappregistry.ListTagsForResourceInput{
		ResourceArn: p.ProductARN,
	}, func(o *servicecatalogappregistry.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, response.Tags)
}
