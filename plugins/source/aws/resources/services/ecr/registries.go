package ecr

import (
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Registries() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecr_registries",
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_DescribeRegistry.html`,
		Resolver:    fetchEcrRegistries,
		Multiplex:   client.ServiceAccountRegionMultiplexer("api.ecr"),
		Transform:   transformers.TransformWithStruct(&ecr.DescribeRegistryOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "registry_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegistryId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
