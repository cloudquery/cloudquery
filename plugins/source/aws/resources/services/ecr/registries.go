package ecr

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Registries() *schema.Table {
	tableName := "aws_ecr_registries"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_DescribeRegistry.html`,
		Resolver:    fetchEcrRegistries,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "api.ecr"),
		Transform:   transformers.TransformWithStruct(&ecr.DescribeRegistryOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "registry_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("RegistryId"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchEcrRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecr
	output, err := svc.DescribeRegistry(ctx, &ecr.DescribeRegistryInput{}, func(options *ecr.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output
	return nil
}
