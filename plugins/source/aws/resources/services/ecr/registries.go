package ecr

import (
	"context"

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
		Transform: transformers.TransformWithStruct(&ecr.DescribeRegistryOutput{},
			transformers.WithPrimaryKeyComponents("RegistryId"),
			transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchEcrRegistries(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEcr).Ecr
	output, err := svc.DescribeRegistry(ctx, &ecr.DescribeRegistryInput{}, func(options *ecr.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output
	return nil
}
