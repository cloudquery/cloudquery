package terraform

import (
	"encoding/json"
	"fmt"
)

// LoadState Terraform show json output into struct
// We are getting content bytes here and not a stream as it's not possible to lazyload a
// json and it needs to be all in memory anyway.
func LoadTerraformShowOutput(content []byte) (*ShowOutput, error) {
	o := ShowOutput{}
	if err := json.Unmarshal(content, &o); err != nil {
		return nil, fmt.Errorf("error parsing terraform show output json: %w", err)
	}
	return &o, nil
}
