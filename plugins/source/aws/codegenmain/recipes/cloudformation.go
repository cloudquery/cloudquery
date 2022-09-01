package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func init() {
	add(parentize(&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:      &types.Stack{},
		AWSService:     "Cloudformation",
		Template:       "resource_get",
		ItemsStruct:    &cloudformation.DescribeStacksOutput{},
		ColumnOverrides: map[string]codegen.ColumnDefinition{
			"stack_id": {
				Name: "arn",
			},
		},
	},
		&Resource{
			AWSStruct:       &types.StackResourceSummary{},
			Template:        "resource_get",
			ItemsStruct:     &cloudformation.ListStackResourcesOutput{},
			ParentFieldName: "StackName",
			CustomErrorBlock: `
			if client.IsErrorRegex(err, "ValidationError", resolvers.ValidStackNotFoundRegex) {
				meta.Logger().Debug("received ValidationError on ListStackResources, stack does not exist", "region", cl.Region, "err", err)
				return nil
			}
`,
			Imports: []string{
				`resolvers "github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/resolvers/cloudformation"`,
			},
		},
	)...)
}
