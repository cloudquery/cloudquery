package ecrpublic

import (
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RepositoryImages() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecrpublic_repository_images",
		Description: `https://docs.aws.amazon.com/AmazonECRPublic/latest/APIReference/API_ImageDetail.html`,
		Resolver:    fetchEcrpublicRepositoryImages,
		Transform:   transformers.TransformWithStruct(&types.ImageDetail{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveImageArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
