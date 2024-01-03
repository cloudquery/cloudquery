package specs

import (
	"fmt"
	"slices"
	"strings"
)

// Spec part to define exact plugin: name, version & location.
type Metadata struct {
	// Name of the plugin to use
	Name string `json:"name" jsonschema:"required,minLength=1"`

	// Version of the plugin to be used
	Version string `json:"version"`

	// Path is the canonical path to the source plugin in a given registry
	// For example:
	// * for `registry: github` the `path` will be: `org/repo`
	// * for `registry: local` the `path` will be the path to the binary: `./path/to/binary`
	// * for `registry: grpc` the `path` will be the address of the gRPC server: `host:port`
	Path string `json:"path" jsonschema:"required,minLength=1"`

	// Registry can be "", "github", "local", "grpc", "docker", "cloudquery"
	Registry Registry `json:"registry,omitempty" jsonschema:"default=cloudquery"`

	// registryInferred is a flag that indicates whether the registry was inferred from an empty value
	registryInferred bool
}

func (m *Metadata) Validate() error {
	if m.Name == "" {
		return fmt.Errorf("name is required")
	}

	if m.Path == "" {
		msg := "path is required"
		// give a small hint to help users transition from the old config format that didn't require path
		officialPlugins := []string{"aws", "azure", "gcp", "k8s", "postgresql", "clickhouse"}
		if slices.Contains(officialPlugins, m.Name) {
			msg += fmt.Sprintf(". Hint: try setting path to cloudquery/%s in your config", m.Name)
		}
		return fmt.Errorf(msg)
	}

	if m.Registry.NeedVersion() {
		if m.Version == "" {
			return fmt.Errorf("version is required")
		}
		if !strings.HasPrefix(m.Version, "v") {
			return fmt.Errorf("version must start with v")
		}
	}

	return nil
}

func (m *Metadata) SetDefaults() {
	if m.Registry == RegistryUnset {
		m.Registry = RegistryCloudQuery
		m.registryInferred = true
	}
}

func (m *Metadata) RegistryInferred() bool {
	return m.registryInferred
}
