package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)



func CodeBuildResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "projects",
			Struct: &types.Project{},
			Multiplex: `client.ServiceAccountRegionMultiplexer("codebuild")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: `schema.PathResolver("Arn")`,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
		},
	}

	for _, r := range resources {
		r.Service = "codebuild"
	}
	return resources
}