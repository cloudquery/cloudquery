package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func StepFunctionResources() []*Resource {
	mx := `client.ServiceAccountRegionMultiplexer("states")`
	resources := []*Resource{
		{
			SubService:          "state_machines",
			Struct:              new(sfn.DescribeStateMachineOutput),
			Description:         "https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeStateMachine.html",
			Multiplex:           mx,
			PKColumns:           []string{"arn"},
			SkipFields:          []string{"ResultMetadata", "StateMachineArn"},
			PreResourceResolver: "getStepFunction",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("StateMachineArn")`,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveStepFunctionTags`,
					},
				}...),
		},
	}
	for _, r := range resources {
		r.Service = "stepfunctions"
	}
	return resources
}
