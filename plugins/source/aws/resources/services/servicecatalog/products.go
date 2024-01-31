package servicecatalog

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Products() *schema.Table {
	tableName := "aws_servicecatalog_products"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeProductAsAdmin.html`,
		Resolver:            fetchServicecatalogProducts,
		PreResourceResolver: getServicecatalogProduct,
		Transform:           transformers.TransformWithStruct(&servicecatalog.DescribeProductAsAdminOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "servicecatalog"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ProductViewDetail.ProductARN"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchServicecatalogProducts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceServicecatalog).Servicecatalog

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

func getServicecatalogProduct(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceServicecatalog).Servicecatalog

	response, err := svc.DescribeProductAsAdmin(ctx, &servicecatalog.DescribeProductAsAdminInput{
		Id: resource.Item.(types.ProductViewDetail).ProductViewSummary.ProductId,
	}, func(o *servicecatalog.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	resource.Item = response
	return nil
}
