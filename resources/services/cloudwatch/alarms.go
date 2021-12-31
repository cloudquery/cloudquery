package cloudwatch

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
		Description:  "The details about a metric alarm.",
		Resolver:     fetchCloudwatchAlarms,
		Multiplex:    client.ServiceAccountRegionMultiplexer("logs"),
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
				Name:        "actions_enabled",
				Description: "Indicates whether actions should be executed during any changes to the alarm state.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "actions",
				Description: "The actions to execute when this alarm transitions to the ALARM state from any other state.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("AlarmActions"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the alarm.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AlarmArn"),
			},
			{
				Name:        "configuration_updated_timestamp",
				Description: "The time stamp of the last update to the alarm configuration.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("AlarmConfigurationUpdatedTimestamp"),
			},
			{
				Name:        "description",
				Description: "The description of the alarm.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AlarmDescription"),
			},
			{
				Name:        "name",
				Description: "The name of the alarm.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AlarmName"),
			},
			{
				Name:        "comparison_operator",
				Description: "The arithmetic operation to use when comparing the specified statistic and threshold.",
				Type:        schema.TypeString,
			},
			{
				Name:        "datapoints_to_alarm",
				Description: "The number of data points that must be breaching to trigger the alarm.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "dimensions",
				Description: "The dimensions for the metric associated with the alarm.",
				Type:        schema.TypeJSON,
				Resolver:    resolveCloudwatchAlarmDimensions,
			},
			{
				Name:        "evaluate_low_sample_count_percentile",
				Description: "Used only for alarms based on percentiles.",
				Type:        schema.TypeString,
			},
			{
				Name:        "evaluation_periods",
				Description: "The number of periods over which data is compared to the specified threshold.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "extended_statistic",
				Description: "The percentile statistic for the metric associated with the alarm.",
				Type:        schema.TypeString,
			},
			{
				Name:        "insufficient_data_actions",
				Description: "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "metric_name",
				Description: "The name of the metric associated with the alarm, if this is an alarm based on a single metric.",
				Type:        schema.TypeString,
			},
			{
				Name:        "namespace",
				Description: "The namespace of the metric associated with the alarm.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ok_actions",
				Description: "The actions to execute when this alarm transitions to the OK state from any other state.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("OKActions"),
			},
			{
				Name:        "period",
				Description: "The period, in seconds, over which the statistic is applied.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "state_reason",
				Description: "An explanation for the alarm state, in text format.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state_reason_data",
				Description: "An explanation for the alarm state, in JSON format.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state_updated_timestamp",
				Description: "The time stamp of the last update to the alarm state.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "state_value",
				Description: "The state value for the alarm.",
				Type:        schema.TypeString,
			},
			{
				Name:        "statistic",
				Description: "The statistic for the metric associated with the alarm, other than percentile.",
				Type:        schema.TypeString,
			},
			{
				Name:        "threshold",
				Description: "The value to compare with the specified statistic.",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "threshold_metric_id",
				Description: "In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.",
				Type:        schema.TypeString,
			},
			{
				Name:        "treat_missing_data",
				Description: "Sets how this alarm is to handle missing data points.",
				Type:        schema.TypeString,
			},
			{
				Name:        "unit",
				Description: "The unit of the metric associated with the alarm.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_cloudwatch_alarm_metrics",
				Description: "This structure is used in both GetMetricData and PutMetricAlarm.",
				Resolver:    fetchCloudwatchAlarmMetrics,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"alarm_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "alarm_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudwatch_alarms table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "alarm_arn",
						Description: "The Amazon Resource Name (ARN) of the alarm.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("arn"),
					},
					{
						Name:        "alarm_name",
						Description: "The name of the alarm.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("name"),
					},
					{
						Name:        "id",
						Description: "A short name used to tie this object to the results in the response.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
					{
						Name:        "expression",
						Description: "The math expression to be performed on the returned data, if this object is performing a math expression.",
						Type:        schema.TypeString,
					},
					{
						Name:        "label",
						Description: "A human-readable label for this metric or expression.",
						Type:        schema.TypeString,
					},
					{
						Name:        "metric_stat_metric_dimensions",
						Description: "The dimensions for the metric.",
						Type:        schema.TypeJSON,
						Resolver:    resolveCloudwatchAlarmMetricMetricStatMetricDimensions,
					},
					{
						Name:        "metric_stat_metric_name",
						Description: "The name of the metric.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("MetricStat.Metric.MetricName"),
					},
					{
						Name:        "metric_stat_metric_namespace",
						Description: "The namespace of the metric.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("MetricStat.Metric.Namespace"),
					},
					{
						Name:        "metric_stat_period",
						Description: "The granularity, in seconds, of the returned data points.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("MetricStat.Period"),
					},
					{
						Name:        "metric_stat",
						Description: "The statistic to return.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("MetricStat.Stat"),
					},
					{
						Name:        "metric_stat_unit",
						Description: "When you are using a Put operation, this defines what unit you want to use when storing the metric.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("MetricStat.Unit"),
					},
					{
						Name:        "period",
						Description: "The granularity, in seconds, of the returned data points.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "return_data",
						Description: "When used in GetMetricData, this option indicates whether to return the timestamps and raw data values of this metric.",
						Type:        schema.TypeBool,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCloudwatchAlarms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func resolveCloudwatchAlarmDimensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	alarm := resource.Item.(types.MetricAlarm)
	dimensions := make(map[string]*string)
	for _, d := range alarm.Dimensions {
		dimensions[*d.Name] = d.Value
	}
	return resource.Set("dimensions", dimensions)
}
func fetchCloudwatchAlarmMetrics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	alarm := parent.Item.(types.MetricAlarm)
	res <- alarm.Metrics
	return nil
}
func resolveCloudwatchAlarmMetricMetricStatMetricDimensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
