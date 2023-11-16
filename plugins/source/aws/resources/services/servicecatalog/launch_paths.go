package servicecatalog

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func launchPaths() *schema.Table {
	tableName := "aws_servicecatalog_launch_paths"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_LaunchPathSummary.html`,
		Resolver:    listLaunchPaths,
		Transform:   transformers.TransformWithStruct(&types.LaunchPathSummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "provisioned_product_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:       "product_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("product_id"),
				PrimaryKey: true,
			},
			{
				Name:       "provisioning_artifact_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("provisioning_artifact_id"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
		Relations: schema.Tables{
			provisioningParameters(),
		},
	}
}
func listLaunchPaths(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceServicecatalog).Servicecatalog
	p := parent.Item.(types.ProvisionedProductAttribute)

	input := servicecatalog.ListLaunchPathsInput{
		ProductId: p.ProductId,
	}

	pager := servicecatalog.NewListLaunchPathsPaginator(svc, &input)

	for pager.HasMorePages() {
		page, err := pager.NextPage(ctx, func(o *servicecatalog.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.LaunchPathSummaries
	}
	return nil
}
