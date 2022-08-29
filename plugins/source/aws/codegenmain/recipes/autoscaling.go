package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

var AutoscalingResources = combine(&Resource{
	DefaultColumns:    []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
	AWSStruct:         &types.LaunchConfiguration{},
	AWSService:        "Autoscaling",
	AWSSubService:     "LaunchConfigurations",
	Template:          "resource_get",
	Verb:              "Describe",
	ResponseItemsName: "LaunchConfigurations",
	Imports:           nil,
	MockImports:       nil,
	MockListStruct:    "",
	SkipFields:        nil,
	TrimPrefix:        "launch_configuration_",
	//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
},
	&Resource{
		DefaultColumns:    []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:         &types.ScheduledUpdateGroupAction{},
		AWSService:        "Autoscaling",
		AWSSubService:     "ScheduledActions",
		Template:          "resource_get",
		Verb:              "Describe",
		ResponseItemsName: "ScheduledUpdateGroupActions",
		Imports:           nil,
		MockImports:       nil,
		MockListStruct:    "",
		SkipFields:        nil,
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
	},
	parentize(
		&Resource{
			DefaultColumns:       []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
			AWSStruct:            &types.AutoScalingGroup{},
			AWSService:           "Autoscaling",
			AWSSubService:        "AutoScalingGroups",
			CQSubserviceOverride: "groups",
			Template:             "resource_get",
			Verb:                 "Describe",
			ResponseItemsName:    "AutoScalingGroups",
			Imports:              []string{"errors", "regexp", "github.com/aws/smithy-go"},
			MockImports:          nil,
			MockListStruct:       "",
			SkipFields:           nil,
			//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
			CustomInit: []string{
				"var autoscalingGroupNotFoundRegex = regexp.MustCompile(`AutoScalingGroup name not found|Group .* not found`)",
			},
			CustomResolvers: []string{`
func isAutoScalingGroupNotExistsError(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "ValidationError" && autoscalingGroupNotFoundRegex.MatchString(ae.ErrorMessage()) {
			return true
		}
	}
	return false
}
`,
			},
		},
		&Resource{
			AWSStruct:            &types.ScalingPolicy{},
			AWSSubService:        "Policies",
			CQSubserviceOverride: "scaling_policies",
			Template:             "resource_get",
			ParentFieldName:      "AutoScalingGroupName",
			Verb:                 "Describe",
			ResponseItemsName:    "ScalingPolicies",
			CustomErrorBlock: `
			if isAutoScalingGroupNotExistsError(err) {
				return nil
			}`,
		},
	),
)
