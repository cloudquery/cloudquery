package autoscaling

import (
	"context"
	"errors"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

type autoscalingGroupWrapper struct {
	types.AutoScalingGroup
	NotificationConfigurations []types.NotificationConfiguration
}

var groupNotFoundRegex = regexp.MustCompile(`AutoScalingGroup name not found|Group .* not found`)

func AutoscalingGroups() *schema.Table {
	return &schema.Table{
		Name:         "aws_autoscaling_groups",
		Description:  "Describes an Auto Scaling group.",
		Resolver:     fetchAutoscalingGroups,
		Multiplex:    client.ServiceAccountRegionMultiplexer("autoscaling"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:     "load_balancers",
				Type:     schema.TypeJSON,
				Resolver: resolveAutoscalingGroupLoadBalancers,
			},
			{
				Name:     "load_balancer_target_groups",
				Type:     schema.TypeJSON,
				Resolver: resolveAutoscalingGroupLoadBalancerTargetGroups,
			},
			{
				Name:     "notifications_configurations",
				Type:     schema.TypeJSON,
				Resolver: resolveAutoscalingGroupNotificationsConfigurations,
			},
			{
				Name:        "name",
				Description: "The name of the Auto Scaling group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AutoScalingGroupName"),
			},
			{
				Name:        "availability_zones",
				Description: "One or more Availability Zones for the group.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "created_time",
				Description: "The date and time the group was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "default_cooldown",
				Description: "The duration of the default cooldown period, in seconds.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "desired_capacity",
				Description: "The desired size of the group.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "health_check_type",
				Description: "The service to use for the health checks",
				Type:        schema.TypeString,
			},
			{
				Name:        "max_size",
				Description: "The maximum size of the group.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "min_size",
				Description: "The minimum size of the group.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the Auto Scaling group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AutoScalingGroupARN"),
			},
			{
				Name:          "capacity_rebalance",
				Description:   "Indicates whether Capacity Rebalancing is enabled.",
				Type:          schema.TypeBool,
				IgnoreInTests: true,
			},
			{
				Name:        "enabled_metrics",
				Description: "The metrics enabled for the group.",
				Type:        schema.TypeJSON,
				Resolver:    resolveAutoscalingGroupsEnabledMetrics,
			},
			{
				Name:        "health_check_grace_period",
				Description: "The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "launch_configuration_name",
				Description: "The name of the associated launch configuration.",
				Type:        schema.TypeString,
			},
			{
				Name:          "launch_template_id",
				Description:   "The ID of the launch template",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("LaunchTemplate.LaunchTemplateId"),
				IgnoreInTests: true,
			},
			{
				Name:          "launch_template_name",
				Description:   "The name of the launch template",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("LaunchTemplate.LaunchTemplateName"),
				IgnoreInTests: true,
			},
			{
				Name:          "launch_template_version",
				Description:   "The version number, $Latest, or $Default",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("LaunchTemplate.Version"),
				IgnoreInTests: true,
			},
			{
				Name:        "load_balancer_names",
				Description: "One or more load balancers associated with the group.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:          "max_instance_lifetime",
				Description:   "The maximum amount of time, in seconds, that an instance can be in service. Valid Range: Minimum value of 0.",
				Type:          schema.TypeInt,
				IgnoreInTests: true,
			},
			{
				Name:          "mixed_instances_policy",
				Description:   "The mixed instances policy for the group.",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "new_instances_protected_from_scale_in",
				Description: "Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.",
				Type:        schema.TypeBool,
			},
			{
				Name:          "placement_group",
				Description:   "The name of the placement group into which to launch your instances, if any.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "service_linked_role_arn",
				Description: "The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServiceLinkedRoleARN"),
			},
			{
				Name:          "status",
				Description:   "The current state of the group when the DeleteAutoScalingGroup operation is in progress.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "suspended_processes",
				Description: "The suspended processes associated with the group.",
				Type:        schema.TypeJSON,
				Resolver:    resolveAutoscalingGroupsSuspendedProcesses,
			},
			{
				Name:        "target_group_arns",
				Description: "The Amazon Resource Names (ARN) of the target groups for your load balancer.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("TargetGroupARNs"),
			},
			{
				Name:        "termination_policies",
				Description: "The termination policies for the group.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "vpc_zone_identifier",
				Description: "One or more subnet IDs, if applicable, separated by commas.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VPCZoneIdentifier"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_autoscaling_group_instances",
				Description: "Describes an EC2 instance.",
				Resolver:    schema.PathTableResolver("Instances"),
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of aws_autoscaling_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "availability_zone",
						Description: "The Availability Zone in which the instance is running.",
						Type:        schema.TypeString,
					},
					{
						Name:        "health_status",
						Description: "The last reported health status of the instance",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The ID of the instance.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InstanceId"),
					},
					{
						Name:        "lifecycle_state",
						Description: "A description of the current lifecycle state",
						Type:        schema.TypeString,
					},
					{
						Name:        "protected_from_scale_in",
						Description: "Indicates whether the instance is protected from termination by Amazon EC2 Auto Scaling when scaling in.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "type",
						Description: "The instance type of the EC2 instance.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InstanceType"),
					},
					{
						Name:        "launch_configuration_name",
						Description: "The launch configuration associated with the instance.",
						Type:        schema.TypeString,
					},
					{
						Name:          "launch_template_id",
						Description:   "The ID of the launch template",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("LaunchTemplate.LaunchTemplateId"),
						IgnoreInTests: true,
					},
					{
						Name:          "launch_template_name",
						Description:   "The name of the launch template",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("LaunchTemplate.LaunchTemplateName"),
						IgnoreInTests: true,
					},
					{
						Name:          "launch_template_version",
						Description:   "The version number, $Latest, or $Default",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("LaunchTemplate.Version"),
						IgnoreInTests: true,
					},
					{
						Name:          "weighted_capacity",
						Description:   "The number of capacity units contributed by the instance based on its instance type",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:        "aws_autoscaling_group_tags",
				Description: "Describes a tag for an Auto Scaling group.",
				Resolver:    schema.PathTableResolver("Tags"),
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of aws_autoscaling_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "key",
						Description: "The tag key.",
						Type:        schema.TypeString,
					},
					{
						Name:        "propagate_at_launch",
						Description: "Determines whether the tag is added to new instances as they are launched in the group.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "resource_id",
						Description: "The name of the group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_type",
						Description: "The type of resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "value",
						Description: "The tag value.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_autoscaling_group_scaling_policies",
				Description:   "Describes a scaling policy.",
				Resolver:      fetchAutoscalingGroupScalingPolicies,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of aws_autoscaling_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "adjustment_type",
						Description: "Specifies how the scaling adjustment is interpreted (for example, an absolute number or a percentage)",
						Type:        schema.TypeString,
					},
					{
						Name:        "alarms",
						Description: "The CloudWatch alarms related to the policy.",
						Type:        schema.TypeJSON,
						Resolver:    resolveAutoscalingGroupScalingPoliciesAlarms,
					},
					{
						Name:        "auto_scaling_group_name",
						Description: "The name of the Auto Scaling group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "cooldown",
						Description: "The duration of the policy's cooldown period, in seconds.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "enabled",
						Description: "Indicates whether the policy is enabled (true) or disabled (false).",
						Type:        schema.TypeBool,
					},
					{
						Name:        "estimated_instance_warmup",
						Description: "The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "metric_aggregation_type",
						Description: "The aggregation type for the CloudWatch metrics",
						Type:        schema.TypeString,
					},
					{
						Name:        "min_adjustment_magnitude",
						Description: "The minimum value to scale by when the adjustment type is PercentChangeInCapacity.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "min_adjustment_step",
						Description: "Available for backward compatibility",
						Type:        schema.TypeInt,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the policy.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PolicyARN"),
					},
					{
						Name:        "name",
						Description: "The name of the scaling policy.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PolicyName"),
					},
					{
						Name:        "type",
						Description: "One of the following policy types:  * TargetTrackingScaling  * StepScaling  * SimpleScaling (default)  For more information, see Target tracking scaling policies (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-target-tracking.html) and Step and simple scaling policies (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html) in the Amazon EC2 Auto Scaling User Guide.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PolicyType"),
					},
					{
						Name:        "scaling_adjustment",
						Description: "The amount by which to scale, based on the specified adjustment type",
						Type:        schema.TypeInt,
					},
					{
						Name:        "step_adjustments",
						Description: "A set of adjustments that enable you to scale based on the size of the alarm breach.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("StepAdjustments"),
					},
					{
						Name:        "target_tracking_configuration_target_value",
						Description: "The target value for the metric.",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("TargetTrackingConfiguration.TargetValue"),
					},
					{
						Name:        "target_tracking_configuration_customized_metric_name",
						Description: "The name of the metric.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TargetTrackingConfiguration.CustomizedMetricSpecification.MetricName"),
					},
					{
						Name:        "target_tracking_configuration_customized_metric_namespace",
						Description: "The namespace of the metric.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TargetTrackingConfiguration.CustomizedMetricSpecification.Namespace"),
					},
					{
						Name:        "target_tracking_configuration_customized_metric_statistic",
						Description: "The statistic of the metric.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TargetTrackingConfiguration.CustomizedMetricSpecification.Statistic"),
					},
					{
						Name:        "target_tracking_configuration_customized_metric_dimensions",
						Description: "The dimensions of the metric",
						Type:        schema.TypeJSON,
						Resolver:    resolveAutoscalingGroupScalingPoliciesTargetTrackingConfigurationCustomizedMetricDimensions,
					},
					{
						Name:        "target_tracking_configuration_customized_metric_unit",
						Description: "The unit of the metric.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TargetTrackingConfiguration.CustomizedMetricSpecification.Unit"),
					},
					{
						Name:        "target_tracking_configuration_disable_scale_in",
						Description: "Indicates whether scaling in by the target tracking scaling policy is disabled. If scaling in is disabled, the target tracking scaling policy doesn't remove instances from the Auto Scaling group",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("TargetTrackingConfiguration.DisableScaleIn"),
					},
					{
						Name:        "target_tracking_configuration_predefined_metric_type",
						Description: "The metric type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TargetTrackingConfiguration.PredefinedMetricSpecification.PredefinedMetricType"),
					},
					{
						Name:        "target_tracking_configuration_predefined_metric_resource_label",
						Description: "Identifies the resource associated with the metric type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TargetTrackingConfiguration.PredefinedMetricSpecification.ResourceLabel"),
					},
				},
			},
			{
				Name:          "aws_autoscaling_group_lifecycle_hooks",
				Description:   "Describes a lifecycle hook, which tells Amazon EC2 Auto Scaling that you want to perform an action whenever it launches instances or terminates instances.",
				Resolver:      fetchAutoscalingGroupLifecycleHooks,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "group_cq_id",
						Description: "Unique CloudQuery ID of aws_autoscaling_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "auto_scaling_group_name",
						Description: "The name of the Auto Scaling group for the lifecycle hook.",
						Type:        schema.TypeString,
					},
					{
						Name:        "default_result",
						Description: "Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs",
						Type:        schema.TypeString,
					},
					{
						Name:        "global_timeout",
						Description: "The maximum time, in seconds, that an instance can remain in a Pending:Wait or Terminating:Wait state",
						Type:        schema.TypeInt,
					},
					{
						Name:        "heartbeat_timeout",
						Description: "The maximum time, in seconds, that can elapse before the lifecycle hook times out",
						Type:        schema.TypeInt,
					},
					{
						Name:        "lifecycle_hook_name",
						Description: "The name of the lifecycle hook.",
						Type:        schema.TypeString,
					},
					{
						Name:        "lifecycle_transition",
						Description: "The state of the EC2 instance to which to attach the lifecycle hook",
						Type:        schema.TypeString,
					},
					{
						Name:        "notification_metadata",
						Description: "Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.",
						Type:        schema.TypeString,
					},
					{
						Name:        "notification_target_arn",
						Description: "The ARN of the target that Amazon EC2 Auto Scaling sends notifications to when an instance is in the transition state for the lifecycle hook",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("NotificationTargetARN"),
					},
					{
						Name:        "role_arn",
						Description: "The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoleARN"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAutoscalingGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Autoscaling
	processGroupsBundle := func(groups []types.AutoScalingGroup) error {
		input := autoscaling.DescribeNotificationConfigurationsInput{
			MaxRecords: aws.Int32(100),
		}
		for _, h := range groups {
			input.AutoScalingGroupNames = append(input.AutoScalingGroupNames, *h.AutoScalingGroupName)
		}
		var configurations []types.NotificationConfiguration
		for {
			output, err := svc.DescribeNotificationConfigurations(ctx, &input, func(o *autoscaling.Options) {
				o.Region = c.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}
			configurations = append(configurations, output.NotificationConfigurations...)
			if aws.ToString(output.NextToken) == "" {
				break
			}
			input.NextToken = output.NextToken
		}
		for _, gr := range groups {
			wrapper := autoscalingGroupWrapper{
				AutoScalingGroup:           gr,
				NotificationConfigurations: getNotificationConfigurationByGroupName(*gr.AutoScalingGroupName, configurations),
			}
			res <- wrapper
		}
		return nil
	}

	config := autoscaling.DescribeAutoScalingGroupsInput{}
	for {
		output, err := svc.DescribeAutoScalingGroups(ctx, &config, func(o *autoscaling.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		groups := output.AutoScalingGroups
		for i := 0; i < len(groups); i += 255 {
			end := i + 255

			if end > len(groups) {
				end = len(groups)
			}
			t := groups[i:end]
			err := processGroupsBundle(t)
			if err != nil {
				return diag.WrapError(err)
			}
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveAutoscalingGroupLoadBalancers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(autoscalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	config := autoscaling.DescribeLoadBalancersInput{AutoScalingGroupName: p.AutoScalingGroupName}
	j := map[string]interface{}{}
	for {
		output, err := svc.DescribeLoadBalancers(ctx, &config, func(o *autoscaling.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			if isAutoScalingGroupNotExistsError(err) {
				return nil
			}
			return diag.WrapError(err)
		}
		for _, lb := range output.LoadBalancers {
			j[*lb.LoadBalancerName] = *lb.State
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveAutoscalingGroupLoadBalancerTargetGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(autoscalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	config := autoscaling.DescribeLoadBalancerTargetGroupsInput{AutoScalingGroupName: p.AutoScalingGroupName}
	j := map[string]interface{}{}
	for {
		output, err := svc.DescribeLoadBalancerTargetGroups(ctx, &config, func(o *autoscaling.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			if isAutoScalingGroupNotExistsError(err) {
				return nil
			}
			return diag.WrapError(err)
		}
		for _, lb := range output.LoadBalancerTargetGroups {
			j[*lb.LoadBalancerTargetGroupARN] = *lb.State
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveAutoscalingGroupNotificationsConfigurations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(autoscalingGroupWrapper)
	j := map[string]interface{}{}
	for _, n := range p.NotificationConfigurations {
		j[*n.NotificationType] = *n.TopicARN
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveAutoscalingGroupsEnabledMetrics(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(autoscalingGroupWrapper)
	j := map[string]interface{}{}
	for _, em := range p.EnabledMetrics {
		j[*em.Metric] = *em.Granularity
	}

	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveAutoscalingGroupsSuspendedProcesses(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(autoscalingGroupWrapper)
	j := map[string]interface{}{}
	for _, sp := range p.SuspendedProcesses {
		j[*sp.ProcessName] = *sp.SuspensionReason
	}

	return diag.WrapError(resource.Set(c.Name, j))
}
func fetchAutoscalingGroupScalingPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(autoscalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	config := autoscaling.DescribePoliciesInput{AutoScalingGroupName: p.AutoScalingGroupName}

	for {
		output, err := svc.DescribePolicies(ctx, &config, func(o *autoscaling.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			if isAutoScalingGroupNotExistsError(err) {
				return nil
			}
			return diag.WrapError(err)
		}
		res <- output.ScalingPolicies

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveAutoscalingGroupScalingPoliciesAlarms(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.ScalingPolicy)
	j := map[string]interface{}{}
	for _, a := range p.Alarms {
		j[*a.AlarmName] = *a.AlarmARN
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func resolveAutoscalingGroupScalingPoliciesTargetTrackingConfigurationCustomizedMetricDimensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.ScalingPolicy)
	if p.TargetTrackingConfiguration == nil || p.TargetTrackingConfiguration.CustomizedMetricSpecification == nil {
		return nil
	}
	j := map[string]interface{}{}
	for _, d := range p.TargetTrackingConfiguration.CustomizedMetricSpecification.Dimensions {
		j[*d.Name] = *d.Value
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func fetchAutoscalingGroupLifecycleHooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(autoscalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	config := autoscaling.DescribeLifecycleHooksInput{AutoScalingGroupName: p.AutoScalingGroupName}

	output, err := svc.DescribeLifecycleHooks(ctx, &config, func(o *autoscaling.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		if isAutoScalingGroupNotExistsError(err) {
			return nil
		}
		return diag.WrapError(err)
	}
	res <- output.LifecycleHooks
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func getNotificationConfigurationByGroupName(name string, set []types.NotificationConfiguration) []types.NotificationConfiguration {
	var response []types.NotificationConfiguration
	for _, s := range set {
		if *s.AutoScalingGroupName == name {
			response = append(response, s)
		}
	}
	return response
}

func isAutoScalingGroupNotExistsError(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "ValidationError" && groupNotFoundRegex.MatchString(ae.ErrorMessage()) {
			return true
		}
	}
	return false
}
