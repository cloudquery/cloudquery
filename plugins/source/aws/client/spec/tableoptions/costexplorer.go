package tableoptions

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/invopop/jsonschema"
)

type CostExplorerAPIs struct {
	GetCostAndUsageOpts []CustomGetCostAndUsageInput `json:"get_cost_and_usage,omitempty"`
}

type CustomGetCostAndUsageInput struct {
	costexplorer.GetCostAndUsageInput
}

// JSONSchemaExtend is required to remove `NextPageToken`.
// We use value receiver because of https://github.com/invopop/jsonschema/issues/102
func (CustomGetCostAndUsageInput) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Properties.Delete("NextPageToken")
}

// UnmarshalJSON implements the json.Unmarshaler interface for the CustomLookupEventsOpts type.
// It is the same as default, but allows the use of underscore in the JSON field names.
func (c *CustomGetCostAndUsageInput) UnmarshalJSON(data []byte) error {
	m := map[string]any{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	csr := caser.New()
	changeCaseForObject(m, csr.ToPascal)
	b, _ := json.Marshal(m)
	return json.Unmarshal(b, &c.GetCostAndUsageInput)
}

func (c *CostExplorerAPIs) validateCustomGetCostAndUsage() error {
	for _, opt := range c.GetCostAndUsageOpts {
		if aws.ToString(opt.NextPageToken) != "" {
			return errors.New("invalid input: cannot set NextToken in GetCostAndUsage")
		}
	}
	return nil
}

func (c *CostExplorerAPIs) Validate() error {
	return c.validateCustomGetCostAndUsage()
}
