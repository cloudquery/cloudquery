package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource autoscaling_scheduled_actions --config ./resources/services/autoscaling/gen.hcl --output .
func AutoscalingScheduledActions() *schema.Table {
	return &schema.Table{
		Name:         "aws_autoscaling_scheduled_actions",
		Description:  "Describes a scheduled scaling action.",
		Resolver:     fetchAutoscalingScheduledActions,
		Multiplex:    client.ServiceAccountRegionMultiplexer("autoscaling"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
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
				Name:        "auto_scaling_group_name",
				Description: "The name of the Auto Scaling group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "desired_capacity",
				Description: "The desired capacity is the initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "end_time",
				Description: "The date and time in UTC for the recurring schedule to end",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "max_size",
				Description: "The maximum size of the Auto Scaling group.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "min_size",
				Description: "The minimum size of the Auto Scaling group.",
				Type:        schema.TypeInt,
			},
			{
				Name:          "recurrence",
				Description:   "The recurring schedule for the action, in Unix cron syntax format",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the scheduled action.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ScheduledActionARN"),
			},
			{
				Name:        "name",
				Description: "The name of the scheduled action.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ScheduledActionName"),
			},
			{
				Name:        "start_time",
				Description: "The date and time in UTC for this action to start",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "time",
				Description: "This parameter is no longer used.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:          "time_zone",
				Description:   "The time zone for the cron expression.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAutoscalingScheduledActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Autoscaling
	params := &autoscaling.DescribeScheduledActionsInput{
		MaxRecords: aws.Int32(100),
	}
	for {
		output, err := svc.DescribeScheduledActions(ctx, params, func(options *autoscaling.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, scheduledUpdateGroupAction := range output.ScheduledUpdateGroupActions {
			res <- scheduledUpdateGroupAction
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return nil
}
