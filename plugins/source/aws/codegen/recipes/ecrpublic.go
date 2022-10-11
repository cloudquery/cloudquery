package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ECRPublicResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "repositories",
			Struct:     &types.Repository{},
			SkipFields: []string{"RepositoryArn"},
			Multiplex:  `client.ServiceAccountRegionMultiplexer("api.ecr-public")`,
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
			SubService: "repository_images",
			Struct:     &types.ImageDetail{},
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
