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
		Imports:                    nil,
		MockImports:                nil,
		MockListStruct:             "",
		SkipFields:                 nil,
		TrimPrefix:                 "policy_",
		CustomInputs: []string{
			"\tServiceNamespace: types.ServiceNamespace(cl.AutoscalingNamespace),",
		},
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
	},
}
