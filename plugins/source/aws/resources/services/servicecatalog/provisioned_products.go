package servicecatalog

import (
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ProvisionedProducts() *schema.Table {
	tableName := "aws_servicecatalog_provisioned_products"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProvisionedProductAttribute.html`,
		Resolver:    fetchServicecatalogProvisionedProducts,
		Transform:   client.TransformWithStruct(&types.ProvisionedProductAttribute{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "servicecatalog"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveProvisionedProductTags,
			},
		},
	}
}
