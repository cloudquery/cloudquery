package servicecatalog

import (
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Products() *schema.Table {
	tableName := "aws_servicecatalog_products"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProductViewDetail.html`,
		Resolver:    fetchServicecatalogProducts,
		Transform:   client.TransformWithStruct(&types.ProductViewDetail{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "servicecatalog"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
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
