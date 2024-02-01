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
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
		Relations: schema.Tables{
			provisioningArtifact(),
			launchPaths(),
		},
	}
}

func fetchServicecatalogProvisionedProducts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceServicecatalog).Servicecatalog

	listInput := new(servicecatalog.SearchProvisionedProductsInput)
	paginator := servicecatalog.NewSearchProvisionedProductsPaginator(svc, listInput)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *servicecatalog.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ProvisionedProducts
	}

	return nil
}
