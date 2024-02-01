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

func Portfolios() *schema.Table {
	tableName := "aws_servicecatalog_portfolios"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribePortfolio.html`,
		Resolver:            fetchServicecatalogPortfolios,
		PreResourceResolver: getPortfolio,
		Transform:           transformers.TransformWithStruct(&servicecatalog.DescribePortfolioOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "servicecatalog"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("PortfolioDetail.ARN"),
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

func fetchServicecatalogPortfolios(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceServicecatalog).Servicecatalog
	paginator := servicecatalog.NewListPortfoliosPaginator(svc, &servicecatalog.ListPortfoliosInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *servicecatalog.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.PortfolioDetails
	}

	return nil
}

func getPortfolio(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceServicecatalog).Servicecatalog
	response, err := svc.DescribePortfolio(ctx, &servicecatalog.DescribePortfolioInput{
		Id: resource.Item.(types.PortfolioDetail).Id,
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
