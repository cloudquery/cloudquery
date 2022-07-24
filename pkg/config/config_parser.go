package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/spf13/viper"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

//go:embed schema.json
var configSchemaYAML []byte

func (p *Parser) LoadConfigFromSource(data []byte) (*Config, *gojsonschema.Result, error) {
	newData := os.Expand(string(data), p.getVariableValue)
	config, results, err := decodeConfig(strings.NewReader(newData))
	if err != nil || !results.Valid() {
		return nil, results, err
	}

	if err := ProcessConfig(config); err != nil {
		return nil, nil, err
	}
	return config, nil, nil
}

func (p *Parser) LoadConfigFile(path string) (*Config, *gojsonschema.Result, error) {
	contents, err := p.LoadFile(path)
	if err != nil {
		return nil, nil, err
	}
	return p.LoadConfigFromSource(contents)
}

// ProcessConfig handles the configuration after it was loaded and parsed
// 1. Assigns defaults after decoding the raw configuration format
// 2. Overrides configuration values from CLI flags
// 3. Validates the configuration provided by the user
// 4. Normalizes the configuration to make it easier to use
func ProcessConfig(config *Config) error {
	assignDefaults(config)
	overrideFromCLIFlags(config)
	if err := validate(config); err != nil {
		return err
	}

	normalize(config)
	return nil
}

func ParseVersion(version string) (*semver.Version, error) {
	return semver.NewVersion(version)
}

func FormatVersion(version *semver.Version) string {
	return "v" + version.String()
}

func isVersionLatest(version string) bool {
	return version == "latest"
}

func validate(config *Config) error {
	if err := validateCloudQueryProviders(config.CloudQuery.Providers); err != nil {
		return err
	}
	if err := validateConnection(config.CloudQuery.Connection); err != nil {
		return err
	}

	if err := validateProvidersBlock(config); err != nil {
		return err
	}
	return nil
}

func assignDefaults(config *Config) {
	// TODO: decode in a more generic way
	if config.CloudQuery.Connection == nil {
		config.CloudQuery.Connection = &Connection{}
	}
}

func overrideFromCLIFlags(config *Config) {
	datadir := viper.GetString("data-dir")
	if datadir != "" {
		config.CloudQuery.PluginDirectory = filepath.Join(datadir, "providers")
	}

	if datadir != "" {
		config.CloudQuery.PolicyDirectory = filepath.Join(datadir, "policies")
	}

	if ds := viper.GetString("dsn"); ds != "" {
		config.CloudQuery.Connection.DSN = ds
	}
}

func validateCloudQueryProviders(providers RequiredProviders) error {
	for _, cp := range providers {
		if isVersionLatest(cp.Version) {
			continue
		}

		_, err := ParseVersion(cp.Version)
		if err != nil {
			return fmt.Errorf("Provider %q version %q is invalid. Please set to 'latest' a or valid semantic version", cp.Name, cp.Version)
		}
	}

	return nil
}

func decodeConfig(r io.Reader) (*Config, *gojsonschema.Result, error) {
	var yc struct {
		CloudQuery CloudQuery  `yaml:"cloudquery" json:"cloudquery"`
		Providers  []*Provider `yaml:"providers" json:"providers"`
	}

	d := yaml.NewDecoder(r)
	d.KnownFields(true)
	if err := d.Decode(&yc); err != nil {
		return nil, nil, fmt.Errorf("failed to decode config: %w", err)
	}

	schemaLoader := gojsonschema.NewBytesLoader(configSchemaYAML)
	documentLoader := gojsonschema.NewGoLoader(yc)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to validate config: %w", err)
	}
	if !result.Valid() {
		return nil, result, nil
	}

	providers := yc.Providers
	for _, p := range providers {
		p.ConfigBytes, err = yaml.Marshal(p.Configuration)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to marshal provider configuration: %w", err)
		}
	}

	return &Config{
		CloudQuery: yc.CloudQuery,
		Providers:  providers,
	}, nil, nil
}

func validateConnection(connection *Connection) error {

	dsnFromFlag := viper.GetString("dsn")

	if dsnFromFlag == "" && connection.DSNFile != "" {
		if connection.DSN != "" {
			return fmt.Errorf("DSN file specified along with literal DSN, only one type is supported")
		}

		dsnBytes, err := ioutil.ReadFile(connection.DSNFile)
		if err != nil {
			return fmt.Errorf("failed to read DSN file: %w", err)
		}

		connection.DSN = string(bytes.TrimSpace(dsnBytes))
	}

	// We support both a `dsn` string for backwards compatibility, a dsn file, or connection configuration parameters (host, database, etc.)
	// If a user configured multiple, we error out unless the dsn was configured via a CLI flag (e.g. `cloudquery fetch --dsn <dsn>`)
	if connection.DSN != "" {
		// allow using a DSN flag even if the config file has explicitly attributes (user, password, host, database, etc.)
		if dsnFromFlag == "" && connection.IsAnyConnParamsSet() {
			return fmt.Errorf("DSN specified along with explicit attributes, only one type is supported")
		}
		return nil
	}

	if connection.Host == "" {
		return fmt.Errorf("host is required")
	}
	if connection.Database == "" {
		return fmt.Errorf("database is required")
	}

	return nil
}

// Validates the `cloudquery.providers` configuration block
func validateProvidersBlock(config *Config) error {
	existingProviders := make(map[string]bool, len(config.Providers))

	// We don't allow duplicate provider names or aliases
	for _, provider := range config.Providers {
		if provider.Alias != "" {
			_, aliasExists := existingProviders[provider.Alias]
			if aliasExists {
				continue
			}
			existingProviders[provider.Alias] = true
		} else {
			_, nameExists := existingProviders[provider.Name]
			if nameExists {
				continue
			}
			existingProviders[provider.Name] = true
		}
	}
	return nil
}

func normalize(config *Config) {
	for _, cloudqueryProvider := range config.CloudQuery.Providers {
		if isVersionLatest(cloudqueryProvider.Version) {
			continue
		}

		ver, _ := ParseVersion(cloudqueryProvider.Version)
		// convert partial versions such as "0.10" to "v0.10.0"
		cloudqueryProvider.Version = FormatVersion(ver)
	}

	// Backwards compatibility. Don't override DSN if was provided by the user
	if config.CloudQuery.Connection.DSN == "" {
		config.CloudQuery.Connection.BuildFromConnParams()
	}
	for _, provider := range config.Providers {
		// Alias should default to provider name
		if provider.Alias == "" {
			provider.Alias = provider.Name
		}
	}
}
