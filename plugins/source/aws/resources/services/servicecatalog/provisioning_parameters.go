package servicecatalog

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func provisioningParameters() *schema.Table {
	tableName := "aws_servicecatalog_provisioning_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeProvisioningParameters.html`,
		Resolver:    fetchProvisioningParameters,
		Transform:   transformers.TransformWithStruct(&servicecatalog.DescribeProvisioningParametersOutput{}, transformers.WithSkipFields("ResultMetadata", "ProvisioningArtifactOutputs")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "provisioned_product_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("provisioned_product_arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "product_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            grandParentColumnResolver("product_id"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "provisioning_artifact_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            grandParentColumnResolver("provisioning_artifact_id"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "path_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("id"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchProvisioningParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceServicecatalog).Servicecatalog
	p := parent.Parent.Item.(types.ProvisionedProductAttribute)
	launchPathSummary := parent.Item.(types.LaunchPathSummary)
	input := servicecatalog.DescribeProvisioningParametersInput{
		ProductId:              p.ProductId,
		ProvisioningArtifactId: p.ProvisioningArtifactId,
		PathId:                 launchPathSummary.Id,
	}

	resp, err := svc.DescribeProvisioningParameters(ctx, &input, func(o *servicecatalog.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- resp
	return nil
}

func grandParentColumnResolver(name string) func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		return r.Set(c.Name, r.Parent.Get(name))
	}
}
