package servicecatalog

import (
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Products() *schema.Table {
	return &schema.Table{
		Name:        "aws_servicecatalog_products",
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProductViewDetail.html`,
		Resolver:    fetchServicecatalogProducts,
		Transform:   transformers.TransformWithStruct(&types.ProductViewDetail{}),
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
				Resolver: schema.PathResolver("ProductARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveProductTags,
			},
		},
	}
}
