package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/invopop/jsonschema"
)

type (
	CloudwatchMetrics []CloudwatchMetric
	CloudwatchMetric  struct {
		ListMetricsOpts         CloudwatchListMetricsInput           `json:"list_metrics,omitempty"`
		GetMetricStatisticsOpts []CloudwatchGetMetricStatisticsInput `json:"get_metric_statistics,omitempty"`
	}
)

func (c CloudwatchMetrics) Validate() error {
	for _, m := range c {
		if err := m.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (c *CloudwatchMetric) Validate() error {
	return errors.Join(c.validateListMetricsOpts(), c.validateGetMetricStatisticsOpts())
}

func (c *CloudwatchMetric) validateListMetricsOpts() error {
	if aws.ToString(c.ListMetricsOpts.NextToken) != "" {
		return errors.New("invalid input: cannot set NextToken in CloudwatchMetrics.ListMetricsOpts")
	}
	return nil
}

func (c CloudwatchMetric) validateGetMetricStatisticsOpts() error {
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

type CloudwatchGetMetricStatisticsInput struct {
	cloudwatch.GetMetricStatisticsInput
}

// JSONSchemaExtend is required to remove `namespace`, `metric_name` & `dimensions`.
// We use value receiver because of https://github.com/invopop/jsonschema/issues/102
func (CloudwatchGetMetricStatisticsInput) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Properties.Delete("Namespace")  // we don't allow `namespace`
	sc.Properties.Delete("MetricName") // we don't allow `metric_name`
	sc.Properties.Delete("Dimensions") // we don't allow `dimensions`
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CloudwatchGetMetricStatisticsInput type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (c *CloudwatchGetMetricStatisticsInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &c.GetMetricStatisticsInput)
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
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &c.ListMetricsInput)
}

// JSONSchemaExtend is required to remove `next_token`.
// We use value receiver because of https://github.com/invopop/jsonschema/issues/102
func (CloudwatchListMetricsInput) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Properties.Delete("NextToken") // we don't allow `next_token`
}
