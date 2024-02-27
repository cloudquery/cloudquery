package specs

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

// Spec part to define exact plugin: name, version & location.
type Metadata struct {
	// Name of the plugin to use
	Name string `json:"name" jsonschema:"required,minLength=1"`

	// Version of the plugin to be used
	Version string `json:"version"`

	// Path is the canonical path to the plugin in a given registry
	// For example:
	// * for `registry: github` the `path` will be: `org/repo`
	// * for `registry: local` the `path` will be the path to the binary: `./path/to/binary`
	// * for `registry: grpc` the `path` will be the address of the gRPC server: `host:port`
	// * for `registry: cloudquery` the `path` will be: `team/name`
	Path string `json:"path" jsonschema:"required,minLength=1"`

	// Registry can be "", "github", "local", "grpc", "docker", "cloudquery"
	Registry Registry `json:"registry,omitempty" jsonschema:"default=cloudquery"`

	// DockerRegistryAuthToken is the token to use to authenticate with the docker registry
	DockerRegistryAuthToken string `json:"docker_registry_auth_token,omitempty"`

	// registryInferred is a flag that indicates whether the registry was inferred from an empty value
	registryInferred bool
}

func (m Metadata) VersionString() string {
	switch m.Registry {
	case RegistryCloudQuery:
		return fmt.Sprintf("%s (%s@%s)", m.Name, m.Path, m.Version)
	case RegistryLocal, RegistryGRPC:
		return fmt.Sprintf("%s (%s@%s)", m.Name, m.Registry, m.Path)
	default:
		return fmt.Sprintf("%s@%s (%s@%s)", m.Name, m.Registry, m.Path, m.Version)
	}
}

func (m *Metadata) Validate() error {
	if m.Name == "" {
		return errors.New("name is required")
	}

	if m.Path == "" {
		msg := "path is required"
		// give a small hint to help users transition from the old config format that didn't require path
		officialPlugins := []string{
			"aws",
			"azure",
			"gcp",
			"digitalocean",
			"github",
			"heroku",
			"k8s",
			"okta",
			"terraform",
			"cloudflare",
			"postgresql",
			"csv",
			"clickhouse",
		}
		if slices.Contains(officialPlugins, m.Name) {
			msg += fmt.Sprintf(". Hint: try setting path to cloudquery/%s in your config", m.Name)
		}
		return errors.New(msg)
	}

	if m.Registry.NeedVersion() {
		if m.Version == "" {
			return errors.New("version is required")
		}
		if !strings.HasPrefix(m.Version, "v") {
			return errors.New("version must start with v")
		}
	}

	return nil
}

// JSONSchemaExtend has to be in sync with Registry.NeedVersion
func (Metadata) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.If = &jsonschema.Schema{
		Title: "if registry is unset or is either `github` or `cloudquery`",
		Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
			properties := orderedmap.New[string, *jsonschema.Schema]()
			properties.Set("registry", &jsonschema.Schema{Enum: []any{RegistryUnset.String(), RegistryGitHub.String(), RegistryCloudQuery.String()}})
			return properties
		}(),
	}
	sc.Then = &jsonschema.Schema{
		Title: "require version to be present",
		Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
			properties := orderedmap.New[string, *jsonschema.Schema]()
			version := *sc.Properties.Value("version")
			version.Pattern = `^v.*$` // v1.2.3, v1, v0
			version.Description = ""
			version.Default = nil
			properties.Set("version", &version)
			return properties
		}(),
		Required: []string{"version"},
	}
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
