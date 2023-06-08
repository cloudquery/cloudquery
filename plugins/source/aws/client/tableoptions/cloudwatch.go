package tableoptions

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/cloudquery/plugin-sdk/v3/caser"
)

type CloudwatchMetricStatistics struct {
	GetMetricStatisticsOpts []CustomCloudwatchGetMetricStatisticsInput `json:"get_metric_statistics,omitempty"`
}

type CustomCloudwatchGetMetricStatisticsInput struct {
	cloudwatch.GetMetricStatisticsInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomCloudwatchGetMetricStatisticsInput type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (c *CustomCloudwatchGetMetricStatisticsInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &c.GetMetricStatisticsInput)
}
