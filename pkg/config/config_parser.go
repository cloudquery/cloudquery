package config

import (
	_ "embed"
	"fmt"
	"net/url"

	"github.com/spf13/afero"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

//go:embed schema.json
var schema []byte

// ReadConfigFile reads cloudquery config files from either local or remote locations
// and reutrn bytes and diags
// Example of path supported paths:
// `./local/relative/path/to/config.hcl`
// `/absolute/path/to/config.hcl`
// `s3://object/in/remote/location/absolute/path/to/config.hcl`
func ReadConfigFile(path string) ([]byte, error) {
	// Example of path supported paths:
	// `./local/relative/path/to/config.hcl`
	// `/absolute/path/to/config.hcl`
	// `s3://object/in/remote/location/absolute/path/to/config.hcl`
	sanitizedPath, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	fs := afero.Afero{Fs: afero.OsFs{}}
	var content []byte
	if sanitizedPath.Scheme == "" {
		content, err = fs.ReadFile(path)
	} else {
		content, err = loadRemoteFile(path)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	return content, nil
}

func UnmarshalConfig(data []byte) (*Config, *gojsonschema.Result, error) {
	c := Config{}
	err := yaml.Unmarshal([]byte(data), &c)

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
