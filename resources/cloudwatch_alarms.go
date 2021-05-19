package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CloudwatchAlarms() *schema.Table {
	return &schema.Table{
		Name:         "aws_cloudwatch_alarms",
		Resolver:     fetchCloudwatchAlarms,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name: "actions_enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "alarm_actions",
				Type: schema.TypeStringArray,
			},
			{
				Name: "alarm_arn",
				Type: schema.TypeString,
			},
			{
				Name: "alarm_configuration_updated_timestamp",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "alarm_description",
				Type: schema.TypeString,
			},
			{
				Name: "alarm_name",
				Type: schema.TypeString,
			},
			{
				Name: "comparison_operator",
				Type: schema.TypeString,
			},
			{
				Name: "datapoints_to_alarm",
				Type: schema.TypeInt,
			},
			{
				Name:     "dimensions",
				Type:     schema.TypeJSON,
				Resolver: resolveCloudwatchAlarmDimensions,
			},
			{
				Name: "evaluate_low_sample_count_percentile",
				Type: schema.TypeString,
			},
			{
				Name: "evaluation_periods",
				Type: schema.TypeInt,
			},
			{
				Name: "extended_statistic",
				Type: schema.TypeString,
			},
			{
				Name: "insufficient_data_actions",
				Type: schema.TypeStringArray,
			},
			{
				Name: "metric_name",
				Type: schema.TypeString,
			},
			{
				Name: "namespace",
				Type: schema.TypeString,
			},
			{
				Name:     "ok_actions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("OKActions"),
			},
			{
				Name: "period",
				Type: schema.TypeInt,
			},
			{
				Name: "state_reason",
				Type: schema.TypeString,
			},
			{
				Name: "state_reason_data",
				Type: schema.TypeString,
			},
			{
				Name: "state_updated_timestamp",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "state_value",
				Type: schema.TypeString,
			},
			{
				Name: "statistic",
				Type: schema.TypeString,
			},
			{
				Name: "threshold",
				Type: schema.TypeFloat,
			},
			{
				Name: "threshold_metric_id",
				Type: schema.TypeString,
			},
			{
				Name: "treat_missing_data",
				Type: schema.TypeString,
			},
			{
				Name: "unit",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_cloudwatch_alarm_metrics",
				Resolver: fetchCloudwatchAlarmMetrics,
				Columns: []schema.Column{
					{
						Name:     "alarm_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "metric_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name: "expression",
						Type: schema.TypeString,
					},
					{
						Name: "label",
						Type: schema.TypeString,
					},
					{
						Name:     "metric_stat_metric_dimensions",
						Type:     schema.TypeJSON,
						Resolver: resolveCloudwatchAlarmMetricMetricStatMetricDimensions,
					},
					{
						Name:     "metric_stat_metric_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("MetricStat.Metric.MetricName"),
					},
					{
						Name:     "metric_stat_metric_namespace",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("MetricStat.Metric.Namespace"),
					},
					{
						Name:     "metric_stat_period",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("MetricStat.Period"),
					},
					{
						Name:     "metric_stat",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("MetricStat.Stat"),
					},
					{
						Name:     "metric_stat_unit",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("MetricStat.Unit"),
					},
					{
						Name: "period",
						Type: schema.TypeInt,
					},
					{
						Name: "return_data",
						Type: schema.TypeBool,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCloudwatchAlarms(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	var config cloudwatch.DescribeAlarmsInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudwatch
	for {
		response, err := svc.DescribeAlarms(ctx, &config, func(o *cloudwatch.Options) {
			o.Region = c.Region
		})

		if err != nil {
			return err
		}
		res <- response.MetricAlarms
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveCloudwatchAlarmDimensions(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	alarm := resource.Item.(types.MetricAlarm)
	dimensions := make(map[string]*string)
	for _, d := range alarm.Dimensions {
		dimensions[*d.Name] = d.Value
	}
	return resource.Set("dimensions", dimensions)
}
func fetchCloudwatchAlarmMetrics(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	alarm := parent.Item.(types.MetricAlarm)
	res <- alarm.Metrics
	return nil
}
func resolveCloudwatchAlarmMetricMetricStatMetricDimensions(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	metric := resource.Item.(types.MetricDataQuery)
	if metric.MetricStat == nil || metric.MetricStat.Metric == nil {
		return nil
	}
	dimensions := make(map[string]*string)
	for _, d := range metric.MetricStat.Metric.Dimensions {
		dimensions[*d.Name] = d.Value
	}
	return resource.Set("metric_stat_metric_dimensions", dimensions)
}
