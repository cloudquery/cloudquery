package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ServiceCatalogResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "portfolios",
			Struct:      &types.PortfolioDetail{},
			Description: "https://docs.aws.amazon.com/servicecatalog/latest/dg/API_PortfolioDetail.html",
			SkipFields:  []string{"ARN"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolvePortfolioTags`,
					},
				}...),
		},
		{
			SubService:  "products",
			Struct:      &types.ProductViewDetail{},
			Description: "https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProductViewDetail.html",
			SkipFields:  []string{"ProductARN"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ProductARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveProductTags`,
					},
				}...),
		}, {
			SubService:  "provisioned_products",
			Struct:      &types.ProvisionedProductAttribute{},
			Description: "https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProvisionedProductAttribute.html",
			SkipFields:  []string{"Arn", "Tags"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveProvisionedProductTags`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "servicecatalog"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("servicecatalog")`
	}
	return resources
}
