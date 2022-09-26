package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CodePipelineResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "webhooks",
			Struct:     &types.ListWebhookItem{},
			Multiplex:  `client.ServiceAccountRegionMultiplexer("codepipeline")`,
			SkipFields: []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:          "pipelines",
			Struct:              &codepipeline.GetPipelineOutput{},
			Multiplex:           `client.ServiceAccountRegionMultiplexer("codepipeline")`,
			PreResourceResolver: "getPipeline",
			SkipFields:          []string{"ResultMetadata"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolvePipelineArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeString,
						Resolver: `resolvePipelineTags`,
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "codepipeline"
	}
	return resources
}
