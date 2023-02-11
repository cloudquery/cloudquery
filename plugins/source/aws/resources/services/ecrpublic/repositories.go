package ecrpublic

import (
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Repositories() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecrpublic_repositories",
		Description: `https://docs.aws.amazon.com/AmazonECRPublic/latest/APIReference/API_Repository.html`,
		Resolver:    fetchEcrpublicRepositories,
		Multiplex:   client.ServiceAccountRegionMultiplexer("api.ecr-public"),
		Transform:   transformers.TransformWithStruct(&types.Repository{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoryArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRepositoryTags,
			},
		},

		Relations: []*schema.Table{
			RepositoryImages(),
		},
	}
}
