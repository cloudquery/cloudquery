package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Portfolios() *schema.Table {
	tableName := "aws_servicecatalog_portfolios"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_PortfolioDetail.html`,
		Resolver:    fetchServicecatalogPortfolios,
		Transform:   transformers.TransformWithStruct(&types.PortfolioDetail{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "servicecatalog"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolvePortfolioTags,
			},
		},
	}
}

func fetchServicecatalogPortfolios(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Servicecatalog
	pagintor := servicecatalog.NewListPortfoliosPaginator(svc, &servicecatalog.ListPortfoliosInput{})
	for pagintor.HasMorePages() {
		page, err := pagintor.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.PortfolioDetails
	}

	return nil
}

func resolvePortfolioTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	port := resource.Item.(types.PortfolioDetail)

	cl := meta.(*client.Client)
	svc := cl.Services().Servicecatalogappregistry
	response, err := svc.ListTagsForResource(ctx, &servicecatalogappregistry.ListTagsForResourceInput{
		ResourceArn: port.ARN,
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, response.Tags)
}
