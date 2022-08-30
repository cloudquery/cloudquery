package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

var ApplicationautoscalingResources = []*Resource{
	{
		DefaultColumns:             []codegen.ColumnDefinition{AccountIdColumn, RegionColumn, NamespaceColumn},
		AWSStruct:                  &types.ScalingPolicy{},
		AWSService:                 "ApplicationAutoscaling",
		AWSSubService:              "ScalingPolicies",
		MultiplexerServiceOverride: "application-autoscaling",
		CQSubserviceOverride:       "policies",
		Template:                   "resource_get",
		Verb:                       "Describe",
		ResponseItemsName:          "ScalingPolicies",
		TrimPrefix:                 "policy_",
		CustomInputs: []string{
			"\tServiceNamespace: types.ServiceNamespace(cl.AutoscalingNamespace),",
		},
		AddTypesImport: true,
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
	},
}
