package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

var ApplicationautoscalingResources = []*Resource{
	{
		DefaultColumns:             []codegen.ColumnDefinition{AccountIdColumn, RegionColumn, NamespaceColumn},
		AWSStruct:                  &types.ScalingPolicy{},
		AWSService:                 "ApplicationAutoscaling",
		MultiplexerServiceOverride: "application-autoscaling",
		CQSubserviceOverride:       "policies",
		Template:                   "resource_get",
		ItemsStruct:                &applicationautoscaling.DescribeScalingPoliciesOutput{},
		TrimPrefix:                 "policy_",
		CustomInputs: []string{
			"\tServiceNamespace: types.ServiceNamespace(cl.AutoscalingNamespace),",
		},
		AddTypesImport: true,
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
	},
}
