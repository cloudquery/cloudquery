package servicecatalog

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func launchPaths() *schema.Table {
	tableName := "aws_servicecatalog_launch_paths"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeProvisioningParameters.html`,
		Resolver:    listLaunchPaths,
		Transform:   transformers.TransformWithStruct(&servicecatalog.DescribeProvisioningParametersOutput{}, transformers.WithSkipFields("ResultMetadata", "ProvisioningArtifactOutputs")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "servicecatalog"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "product_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   parentPathResolver("ProductId"),
				PrimaryKey: true,
			},
			{
				Name:       "provisioning_artifact_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   parentPathResolver("ProvisioningArtifactId"),
				PrimaryKey: true,
			},
		},
		Relations: schema.Tables{
			provisioningParameters(),
		},
	}
}
func listLaunchPaths(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Servicecatalog
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
