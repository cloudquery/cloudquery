package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/autoscaling"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AutoscalingResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "launch_configurations",
			Struct:     &types.LaunchConfiguration{},
			SkipFields: []string{"LaunchConfigurationARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("LaunchConfigurationARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "groups",
			Struct:     &autoscaling.AutoScalingGroupWrapper{},
			SkipFields: []string{"AutoScalingGroupARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "load_balancers",
						Type:     schema.TypeJSON,
						Resolver: `resolveAutoscalingGroupLoadBalancers`,
					},
					{
						Name:     "load_balancer_target_groups",
						Type:     schema.TypeJSON,
						Resolver: `resolveAutoscalingGroupLoadBalancerTargetGroups`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("AutoScalingGroupARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"GroupScalingPolicies()",
				"GroupLifecycleHooks()",
			},
		},
		{
			SubService: "group_scaling_policies",
			Struct:     &types.ScalingPolicy{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "group_lifecycle_hooks",
			Struct:     &types.LifecycleHook{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "scheduled_actions",
			Struct:     &types.ScheduledUpdateGroupAction{},
			SkipFields: []string{"ScheduledActionARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ScheduledActionARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "autoscaling"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("autoscaling")`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
