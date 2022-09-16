package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ECRResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "repositories",
			Struct:     &types.Repository{},
			SkipFields: []string{"RepositoryArn"},
			Multiplex:  `client.ServiceAccountRegionMultiplexer("api.ecr")`,
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
						Type:     schema.TypeString,
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
		r.Service = "ecr"
	}
	return resources
}
