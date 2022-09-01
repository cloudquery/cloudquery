package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

var AutoscalingResources = combine(&Resource{
	DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
	AWSStruct:      &types.LaunchConfiguration{},
	AWSService:     "Autoscaling",
	Template:       "resource_get",
	ItemsStruct:    &autoscaling.DescribeLaunchConfigurationsOutput{},
	TrimPrefix:     "launch_configuration_",
	//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
},
	&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:      &types.ScheduledUpdateGroupAction{},
		AWSService:     "Autoscaling",
		Template:       "resource_get",
		ItemsStruct:    &autoscaling.DescribeScheduledActionsOutput{},
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
	},
	parentize(
		&Resource{
			DefaultColumns:       []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
			AWSStruct:            &types.AutoScalingGroup{},
			AWSService:           "Autoscaling",
			CQSubserviceOverride: "groups",
			Template:             "resource_get",
			ItemsStruct:          &autoscaling.DescribeAutoScalingGroupsOutput{},
			//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
			// TODO missing the `autoscalingGroupWrapper` for NotificationConfigurations
		},
		&Resource{
			AWSStruct:            &types.ScalingPolicy{},
			CQSubserviceOverride: "scaling_policies",
			Template:             "resource_get",
			ParentFieldName:      "AutoScalingGroupName",
			ItemsStruct:          &autoscaling.DescribePoliciesOutput{},
			Imports: []string{
				`resolvers "github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/resolvers/autoscaling"`,
			},
			CustomErrorBlock: `
			if resolvers.IsGroupNotExistsError(err) {
				return nil
			}`,
		},
		&Resource{
			AWSStruct:            &types.LifecycleHook{},
			CQSubserviceOverride: "lifecycle_hooks",
			Template:             "resource_get",
			ParentFieldName:      "AutoScalingGroupName",
			ItemsStruct:          &autoscaling.DescribeLifecycleHooksOutput{},
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
