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
			//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		},
		&Resource{
			AWSStruct:            &types.ScalingPolicy{},
			AWSSubService:        "Policies",
			CQSubserviceOverride: "scaling_policies",
			Template:             "resource_get",
			ParentFieldName:      "AutoScalingGroupName",
			Verb:                 "Describe",
			ResponseItemsName:    "ScalingPolicies",
			Imports: []string{
				`resolvers "github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/resolvers/autoscaling"`,
			},
			CustomErrorBlock: `
			if resolvers.IsGroupNotExistsError(err) {
				return nil
			}`,
		},
	),
)
