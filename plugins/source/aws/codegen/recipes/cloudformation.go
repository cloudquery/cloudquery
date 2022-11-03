package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CloudformationResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "stacks",
			Struct:      &types.Stack{},
			Description: "https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Stack.html",
			SkipFields:  []string{"StackId"},
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
				}...),
			Relations: []string{
				"StackResources()",
			},
		},
		{
			SubService:   "stack_resources",
			Struct:       &types.StackResourceSummary{},
			Description:  "https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackResourceSummary.html",
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
