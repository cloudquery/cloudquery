package policy

import (
	_ "embed"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

//go:embed policy_schema.json
var schema []byte

func UnmarshalPolicy(content []byte) (*Policy, *gojsonschema.Result, error) {
	c := Policy{}
	err := yaml.Unmarshal([]byte(content), &c)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse yaml: %w", err)
	}

	schemaLoader := gojsonschema.NewBytesLoader(schema)
	documentLoader := gojsonschema.NewGoLoader(c)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to validate config: %w", err)
	}
	return &c, result, err
}
