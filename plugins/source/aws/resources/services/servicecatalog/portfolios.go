package servicecatalog

import (
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Portfolios() *schema.Table {
	return &schema.Table{
		Name:        "aws_servicecatalog_portfolios",
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_PortfolioDetail.html`,
		Resolver:    fetchServicecatalogPortfolios,
		Transform:   transformers.TransformWithStruct(&types.PortfolioDetail{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("servicecatalog"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
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
