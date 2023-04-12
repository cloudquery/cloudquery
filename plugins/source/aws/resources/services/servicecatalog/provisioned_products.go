package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ProvisionedProducts() *schema.Table {
	tableName := "aws_servicecatalog_provisioned_products"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProvisionedProductAttribute.html`,
		Resolver:    fetchServicecatalogProvisionedProducts,
		Transform:   transformers.TransformWithStruct(&types.ProvisionedProductAttribute{}),
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

func fetchServicecatalogProvisionedProducts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Servicecatalog

	listInput := new(servicecatalog.SearchProvisionedProductsInput)
	paginator := servicecatalog.NewSearchProvisionedProductsPaginator(svc, listInput)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.ProvisionedProducts
	}

	return nil
}

func resolveProvisionedProductTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.ProvisionedProductAttribute)
	return resource.Set(c.Name, client.TagsToMap(p.Tags))
}
