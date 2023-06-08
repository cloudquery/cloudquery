package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/cloudquery/plugin-sdk/v3/caser"
)

type CloudwatchAPIs struct {
	ListMetricsOpts []CustomCloudwatchListMetricsInput `json:"list_metrics,omitempty"`
}

type CustomCloudwatchListMetricsInput struct {
	cloudwatch.ListMetricsInput
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomCloudwatchListMetricsInput type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (c *CustomCloudwatchListMetricsInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &c.ListMetricsInput)
}

func (c *CloudwatchAPIs) validateListMetrics() error {
	for _, opt := range c.ListMetricsOpts {
		if aws.ToString(opt.NextToken) != "" {
			return errors.New("invalid input: cannot set NextToken in ListMetrics")
		}
	}
	return nil
}

func (c *CloudwatchAPIs) Validate() error {
	return c.validateListMetrics()
}
