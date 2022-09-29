// Code generated by codegen; DO NOT EDIT.

package autoscaling

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GroupScalingPolicies() *schema.Table {
	return &schema.Table{
		Name:      "aws_autoscaling_group_scaling_policies",
		Resolver:  fetchAutoscalingGroupScalingPolicies,
		Multiplex: client.ServiceAccountRegionMultiplexer("autoscaling"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "adjustment_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AdjustmentType"),
			},
			{
				Name:     "alarms",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Alarms"),
			},
			{
				Name:     "auto_scaling_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingGroupName"),
			},
			{
				Name:     "cooldown",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Cooldown"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
			{
				Name:     "estimated_instance_warmup",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("EstimatedInstanceWarmup"),
			},
			{
				Name:     "metric_aggregation_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricAggregationType"),
			},
			{
				Name:     "min_adjustment_magnitude",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinAdjustmentMagnitude"),
			},
			{
				Name:     "min_adjustment_step",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinAdjustmentStep"),
			},
			{
				Name:     "policy_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyName"),
			},
			{
				Name:     "policy_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyType"),
			},
			{
				Name:     "predictive_scaling_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PredictiveScalingConfiguration"),
			},
			{
				Name:     "scaling_adjustment",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ScalingAdjustment"),
			},
			{
				Name:     "step_adjustments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StepAdjustments"),
			},
			{
				Name:     "target_tracking_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TargetTrackingConfiguration"),
			},
		},
	}
}
