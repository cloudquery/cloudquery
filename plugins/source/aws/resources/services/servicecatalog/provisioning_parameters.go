package servicecatalog

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/thoas/go-funk"
)

func provisioningParameters() *schema.Table {
	tableName := "aws_servicecatalog_provisioning_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeProvisioningParameters.html`,
		Resolver:    fetchProvisioningParameters,
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
			{
				Name:       "path_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolvePathID("ProvisioningArtifactId"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchProvisioningParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Servicecatalog
	p := parent.Parent.Parent.Item.(types.ProvisionedProductAttribute)
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

func parentPathResolver(path string) func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		return r.Set(c.Name, funk.Get(r.Parent.Item, path, funk.WithAllowZero()))
	}
}

func resolvePathID(path string) func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		return r.Set(c.Name, funk.Get(r.Parent.Item, path, funk.WithAllowZero()))
	}
}
