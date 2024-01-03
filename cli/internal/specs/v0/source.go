package specs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/thoas/go-funk"
)

type BackendOptions struct {
	TableName  string `json:"table_name,omitempty"`
	Connection string `json:"connection,omitempty"`
}

// Source is the spec for a source plugin
type Source struct {
	// Name of the source plugin to use
	Name string `json:"name,omitempty"`
	// Version of the source plugin to use
	Version string `json:"version,omitempty"`
	// Path is the canonical path to the source plugin in a given registry
	// For example:
	// in github the path will be: org/repo
	// For the local registry the path will be the path to the binary: ./path/to/binary
	// For the gRPC registry the path will be the address of the gRPC server: host:port
	Path string `json:"path,omitempty"`
	// Registry can be "", "github", "local", "grpc", "docker", "cloudquery"
	Registry Registry `json:"registry,omitempty"`
	// Tables to sync from the source plugin
	Tables []string `json:"tables,omitempty"`
	// SkipTables defines tables to skip when syncing data. Useful if a glob pattern is used in Tables
	SkipTables []string `json:"skip_tables,omitempty"`
	// SkipDependentTables changes the matching behavior with regard to dependent tables. If set to true, dependent tables will not be synced unless they are explicitly matched by Tables.
	SkipDependentTables bool `json:"skip_dependent_tables,omitempty"`
	// Destinations are the names of destination plugins to send sync data to
	Destinations []string `json:"destinations,omitempty"`

	// Optional Backend options for sync operation
	BackendOptions *BackendOptions `json:"backend_options,omitempty"`

	// Spec defines plugin specific configuration
	// This is different in every source plugin.
	Spec map[string]any `json:"spec,omitempty"`

	// DeterministicCQID is a flag that indicates whether the source plugin should generate a random UUID as the value of _cq_id
	// or whether it should calculate a UUID that is a hash of the primary keys (if they exist) or the entire resource.
	DeterministicCQID bool `json:"deterministic_cq_id,omitempty"`

	// If specified this will spawn the plugin with --otel-endpoint
	OtelEndpoint string `json:"otel_endpoint,omitempty"`
	// If specified this will spawn the plugin with --otel-endpoint-insecure
	OtelEndpointInsecure bool `json:"otel_endpoint_insecure,omitempty"`

	// registryInferred is a flag that indicates whether the registry was inferred from an empty value
	registryInferred bool
}

// GetWarnings returns a list of deprecated options that were used in the source config. This should be
// called before SetDefaults.
func (s *Source) GetWarnings() Warnings {
	warnings := make(map[string]string)
	if s.SkipDependentTables && slices.Contains(s.Tables, "*") {
		warnings["skip_dependent_tables"] = "the `skip_dependent_tables` option is ineffective when used with '*' `tables`"
	}
	if slices.Contains(s.Tables, "*") && len(s.Tables) > 1 {
		warnings["all_tables_with_more_tables"] = "`tables` option contains '*' as well as other tables. '*' will match all tables"
	}

	return warnings
}

func (s *Source) SetDefaults() {
	if s.Spec == nil {
		s.Spec = make(map[string]any)
	}
	if s.Registry == RegistryUnset {
		s.Registry = RegistryCloudQuery
		s.registryInferred = true
	}
}

// UnmarshalSpec unmarshals the internal spec into the given interface
func (s *Source) UnmarshalSpec(out any) error {
	b, err := json.Marshal(s.Spec)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	dec.DisallowUnknownFields()
	return dec.Decode(out)
}

func (s *Source) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("name is required")
	}
	if s.Path == "" {
		msg := "path is required"
		// give a small hint to help users transition from the old config format that didn't require path
		officialPlugins := []string{"aws", "azure", "gcp", "digitalocean", "github", "heroku", "k8s", "okta", "terraform", "cloudflare"}
		if funk.ContainsString(officialPlugins, s.Name) {
			msg += fmt.Sprintf(". Hint: try setting path to cloudquery/%s in your config", s.Name)
		}
		return fmt.Errorf(msg)
	}

	if len(s.Tables) == 0 {
		return fmt.Errorf("tables configuration is required. Hint: set the tables you want to sync by adding `tables: [...]` or use `cloudquery tables` to list available tables")
	}

	if s.Registry.NeedVersion() {
		if s.Version == "" {
			return fmt.Errorf("version is required")
		}
		if !strings.HasPrefix(s.Version, "v") {
			return fmt.Errorf("version must start with v")
		}
	}
	if len(s.Destinations) == 0 {
		return fmt.Errorf("at least one destination is required")
	}

	return nil
}

func (s Source) VersionString() string {
	if s.Registry != RegistryGithub {
		return fmt.Sprintf("%s (%s@%s)", s.Name, s.Registry, s.Path)
	}
	pathParts := strings.Split(s.Path, "/")
	if len(pathParts) != 2 {
		return fmt.Sprintf("%s (%s@%s)", s.Name, s.Path, s.Version)
	}
	if s.Name == pathParts[1] {
		return fmt.Sprintf("%s (%s)", s.Name, s.Version)
	}
	return fmt.Sprintf("%s (%s@%s)", s.Name, pathParts[1], s.Version)
}

func (s Source) RegistryInferred() bool {
	return s.registryInferred
}
