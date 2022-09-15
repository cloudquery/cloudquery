package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CloudformationResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "stacks",
			Struct:     &types.Stack{},
			SkipFields: []string{"StackId", "Tags"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("StackId")`,
					},
					{
						// duplicated column "arn" because "StackId" actually contains an ARN,
						// so we store it once as "id" and once as "arn" for consistency with other
						// AWS resources
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("StackId")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
				}...),
			Relations: []string{
				"StackResources()",
			},
		},
		{
			SubService:   "stack_resources",
			Struct:       &types.StackResourceSummary{},
			SkipFields:   []string{},
			ExtraColumns: defaultRegionalColumns,
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "cloudformation"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("cloudformation")`
	}
	return resources
}
