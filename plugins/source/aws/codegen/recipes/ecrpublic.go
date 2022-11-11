package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ECRPublicResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "repositories",
			Struct:      &types.Repository{},
			Description: "https://docs.aws.amazon.com/AmazonECRPublic/latest/APIReference/API_Repository.html",
			SkipFields:  []string{"RepositoryArn"},
			Multiplex:   `client.ServiceAccountRegionMultiplexer("api.ecr-public")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("RepositoryArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveRepositoryTags`,
					},
				}...),
			Relations: []string{"RepositoryImages()"},
		},
		{
			SubService:  "repository_images",
			Struct:      &types.ImageDetail{},
			Description: "https://docs.aws.amazon.com/AmazonECRPublic/latest/APIReference/API_ImageDetail.html",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveImageArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "ecrpublic"
	}
	return resources
}
