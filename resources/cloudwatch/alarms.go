package cloudwatch

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/mitchellh/mapstructure"
	"time"
)


type MetricAlarm struct {
	ID uint `gorm:"primarykey"`
	AccountID string
	Region string
	ActionsEnabled *bool
	AlarmActions []*MetricAlarmActions `gorm:"constraint:OnDelete:CASCADE;"`
	AlarmArn *string
	AlarmConfigurationUpdatedTimestamp *time.Time
	AlarmDescription *string
	AlarmName *string
	ComparisonOperator *string
	DatapointsToAlarm *int32
	EvaluateLowSampleCountPercentile *string
	EvaluationPeriods *int32
	ExtendedStatistic *string
	MetricName *string
	Metrics []*MetricAlarmMetric `gorm:"constraint:OnDelete:CASCADE;"`
	Namespace *string
	Period *int32
	StateReason *string
	StateReasonData *string
	StateUpdatedTimestamp *time.Time
	StateValue *string
	Statistic *string
	Threshold *float64
	ThresholdMetricId *string
	TreatMissingData *string
	Unit *string
}

func (MetricAlarm) TableName() string {
	return "aws_cloudwatch_metric_alarms"
}
type MetricAlarmActions struct {
	ID uint `gorm:"primarykey"`
	AccountID string `gorm:"-"`
	Region string `gorm:"-"`
	MetricAlarmID uint `neo:"ignore"`
	Value string
}

func (MetricAlarmActions) TableName() string {
	return "aws_cloudwatch_metric_alarm_actions"
}

type MetricAlarmMetric struct {
	ID uint `gorm:"primarykey"`
	AccountID string `gorm:"-"`
	Region string `gorm:"-"`
	MetricAlarmID uint `neo:"ignore"`
	Expression *string
	ResourceID *string
	Label *string

	Name      *string
	Namespace *string

	StatPeriod *int32
	StatStat   *string
	StatUnit   string

	Period *int32
	ReturnData *bool
}

func (MetricAlarmMetric) TableName() string {
	return "aws_cloudwatch_metric_alarm_metrics"
}


func (c *Client) transformMetricAlarms(values *[]types.MetricAlarm) []*MetricAlarm {
	var tValues []*MetricAlarm
	for _, value := range *values {
		tValue := MetricAlarm{
			AccountID: c.accountID,
			Region: c.region,
			ActionsEnabled: value.ActionsEnabled,
			AlarmActions: c.transformMetricAlarmActions(&value.AlarmActions),
			AlarmArn: value.AlarmArn,
			AlarmConfigurationUpdatedTimestamp: value.AlarmConfigurationUpdatedTimestamp,
			AlarmDescription: value.AlarmDescription,
			AlarmName: value.AlarmName,
			ComparisonOperator: aws.String(string(value.ComparisonOperator)),
			DatapointsToAlarm: value.DatapointsToAlarm,
			EvaluateLowSampleCountPercentile: value.EvaluateLowSampleCountPercentile,
			EvaluationPeriods: value.EvaluationPeriods,
			ExtendedStatistic: value.ExtendedStatistic,
			MetricName: value.MetricName,
			Metrics: c.transformMetricAlarmMetrics(&value.Metrics),
			Namespace: value.Namespace,
			Period: value.Period,
			StateReason: value.StateReason,
			StateReasonData: value.StateReasonData,
			StateUpdatedTimestamp: value.StateUpdatedTimestamp,
			StateValue: aws.String(string(value.StateValue)),
			Statistic: aws.String(string(value.Statistic)),
			Threshold: value.Threshold,
			ThresholdMetricId: value.ThresholdMetricId,
			TreatMissingData: value.TreatMissingData,
			Unit: aws.String(string(value.Unit)),
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformMetricAlarmMetrics(values *[]types.MetricDataQuery) []*MetricAlarmMetric {
	var tValues []*MetricAlarmMetric
	for _, value := range *values {
		tValue := MetricAlarmMetric{
			AccountID: c.accountID,
			Region: c.region,
			Expression: value.Expression,
			ResourceID: value.Id,
			Label: value.Label,
			Period: value.Period,
			ReturnData: value.ReturnData,
		}
		if value.MetricStat != nil {
			tValue.StatPeriod = value.MetricStat.Period
			tValue.StatStat = value.MetricStat.Stat
			tValue.StatUnit = string(value.MetricStat.Unit)
			if value.MetricStat.Metric != nil {
				tValue.Name = value.MetricStat.Metric.MetricName
				tValue.Namespace = value.MetricStat.Metric.Namespace
			}

		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformMetricAlarmActions(values *[]string) []*MetricAlarmActions {
	var tValues []*MetricAlarmActions
	for _, v := range *values {
		tValues = append(tValues, &MetricAlarmActions{
			Value: v,
		})
	}
	return tValues
}

type MetricAlarmConfig struct {
	Filter string
}

var MetricAlarmTables = []interface{} {
	&MetricAlarm{},
	&MetricAlarmActions{},
	&MetricAlarmMetric{},
}

func (c *Client)alarms(gConfig interface{}) error {
	ctx := context.Background()
	var config cloudwatch.DescribeAlarmsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(MetricAlarmTables...)

	for {
		output, err := c.svc.DescribeAlarms(ctx, &config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformMetricAlarms(&output.MetricAlarms))
		c.log.Info("Fetched resources", "resource", "cloudwatch.alarms", "count", len(output.MetricAlarms))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}

