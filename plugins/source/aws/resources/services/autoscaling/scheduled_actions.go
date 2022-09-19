// Code generated by codegen; DO NOT EDIT.

package autoscaling

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ScheduledActions() *schema.Table {
	return &schema.Table{
		Name:      "aws_autoscaling_scheduled_actions",
		Resolver:  fetchAutoscalingScheduledActions,
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScheduledActionARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "auto_scaling_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingGroupName"),
			},
			{
				Name:     "desired_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DesiredCapacity"),
			},
			{
				Name:     "end_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EndTime"),
			},
			{
				Name:     "max_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxSize"),
			},
			{
				Name:     "min_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinSize"),
			},
			{
				Name:     "recurrence",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Recurrence"),
			},
			{
				Name:     "scheduled_action_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScheduledActionName"),
			},
			{
				Name:     "start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartTime"),
			},
			{
				Name:     "time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Time"),
			},
			{
				Name:     "time_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TimeZone"),
			},
		},
	}
}
