package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ECRResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "registries",
			Struct:     &ecr.DescribeRegistryOutput{},
			SkipFields: []string{"RegistryId", "ResultMetadata"},
			Multiplex:  `client.ServiceAccountRegionMultiplexer("api.ecr")`,
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "registry_id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("RegistryId")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "registry_policies",
			Struct:     &ecr.GetRegistryPolicyOutput{},
			SkipFields: []string{"RegistryId", "PolicyText", "ResultMetadata"},
			Multiplex:  `client.ServiceAccountRegionMultiplexer("api.ecr")`,
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "registry_id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("RegistryId")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "policy_text",
						Type:     schema.TypeJSON,
						Resolver: `client.MarshaledJsonResolver("PolicyText")`,
					},
				}...),
		},
		{
			SubService:  "repositories",
			Struct:      &types.Repository{},
			Description: "https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_Repository.html",
			SkipFields:  []string{"RepositoryArn"},
			Multiplex:   `client.ServiceAccountRegionMultiplexer("api.ecr")`,
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
			Description: "https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageDetail.html",
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
