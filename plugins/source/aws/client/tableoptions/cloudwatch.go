package tableoptions

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/cloudquery/plugin-sdk/v3/caser"
)

type CloudwatchCustomMetricStatistics struct {
	GetMetricStatisticsOpts []CloudwatchGetMetricStatisticsInput `json:"get_metric_statistics,omitempty"`
}

type CloudwatchGetMetricStatisticsInput struct {
	cloudwatch.GetMetricStatisticsInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CloudwatchGetMetricStatisticsInput type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (c *CloudwatchGetMetricStatisticsInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := processRelativeTimes(m, time.Now().UTC(), []string{"start_time", "end_time"}); err != nil {
		return err
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &c.GetMetricStatisticsInput)
}

type CloudwatchMetrics struct {
	ListMetricsOpts         []CloudwatchListMetricsInput         `json:"list_metrics,omitempty"`
	GetMetricStatisticsOpts []CloudwatchGetMetricStatisticsInput `json:"get_metric_statistics,omitempty"`
}

type CloudwatchListMetricsInput struct {
	cloudwatch.ListMetricsInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CloudwatchListMetricsInput type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (c *CloudwatchListMetricsInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := processRelativeTimes(m, time.Now().UTC(), []string{"start_time", "end_time"}); err != nil {
		return err
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &c.ListMetricsInput)
}

func (c *CloudwatchMetrics) validateListMetricsOpts() error {
	for _, opt := range c.ListMetricsOpts {
		if aws.ToString(opt.NextToken) != "" {
			return errors.New("invalid input: cannot set NextToken in CloudwatchMetrics.ListMetricsOpts")
		}
	}
	return nil
}

func (c *CloudwatchMetrics) validateGetMetricStatisticsOpts() error {
	for _, opt := range c.GetMetricStatisticsOpts {
		if aws.ToString(opt.Namespace) != "" {
			return errors.New("invalid input: cannot set Namespace in CloudwatchMetrics.GetMetricStatisticsOpts")
		}
		if aws.ToString(opt.MetricName) != "" {
			return errors.New("invalid input: cannot set MetricName in CloudwatchMetrics.GetMetricStatisticsOpts")
		}
		if len(opt.Dimensions) > 0 {
			return errors.New("invalid input: cannot set Dimensions in CloudwatchMetrics.GetMetricStatisticsOpts")
		}
	}
	return nil
}

func (c *CloudwatchMetrics) Validate() error {
	if err := c.validateListMetricsOpts(); err != nil {
		return err
	}
	return c.validateGetMetricStatisticsOpts()
}
