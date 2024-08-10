package specs

import (
	"bytes"
	"encoding/json"
	"errors"
	"slices"
	"strings"

	"github.com/invopop/jsonschema"
)

// Backend options to be used in conjunction with incremental tables (stores the incremental progres)
type BackendOptions struct {
	// The name of the table to store the key-value pairs for incremental progress.
	TableName string `json:"table_name,omitempty" jsonschema:"required,minLength=1"`

	// Connection string for the destination plugin.
	// Can be either `@@plugin.name.connection` or a fully-qualified gRPC connection string.
	Connection string `json:"connection,omitempty" jsonschema:"required,minLength=1"`
}

// PluginName returns the name of the plugin from the connection string variable.
//
// Note that `Connection` gets string replaced with the actual connection value during the sync
// process, so calling this function will only work before the sync process starts.
func (b *BackendOptions) PluginName() string {
	if b == nil || !strings.HasPrefix(b.Connection, "@@plugins.") {
		return ""
	}
	return strings.Split(b.Connection, ".")[1]
}

// Source plugin spec
type Source struct {
	Metadata

	// Tables to sync from the source plugin
	Tables []string `json:"tables,omitempty" jsonschema:"required,minItems=1,minLength=1"`
	// SkipTables defines tables to skip when syncing data. Useful if a glob pattern is used in Tables
	SkipTables []string `json:"skip_tables,omitempty" jsonschema:"minLength=1"`
	// SkipDependentTables changes the matching behavior with regard to dependent tables. If set to `false`, dependent tables will be included in the sync when their parents are matched, even if not explicitly included by the `tables` configuration.
	SkipDependentTables *bool `json:"skip_dependent_tables,omitempty" jsonschema:"default=true"`

	// Destinations are the names of destination plugins to send sync data to
	Destinations []string `json:"destinations,omitempty" jsonschema:"required,minItems=1,minLength=1"`

	// Optional Backend options for sync operation
	BackendOptions *BackendOptions `json:"backend_options,omitempty"`

	// Source plugin own (nested) spec
	Spec map[string]any `json:"spec,omitempty"`

	// DeterministicCQID is a flag that indicates whether the source plugin should generate a random UUID as the value of `_cq_id`
	// or whether it should calculate a UUID that is a hash of the primary keys (if they exist) or the entire resource.
	DeterministicCQID bool `json:"deterministic_cq_id,omitempty" jsonschema:"default=false"`

	// If specified this will spawn the plugin with `--otel-endpoint`
	OtelEndpoint string `json:"otel_endpoint,omitempty" jsonschema:"default="`
	// If specified this will spawn the plugin with `--otel-endpoint-insecure`
	OtelEndpointInsecure bool `json:"otel_endpoint_insecure,omitempty" jsonschema:"default=false"`
}

// GetWarnings returns a list of deprecated options that were used in the source config. This should be
// called before SetDefaults.
func (s *Source) GetWarnings() Warnings {
	warnings := make(map[string]string)

	if s.SkipDependentTables != nil && *s.SkipDependentTables && slices.Contains(s.Tables, "*") {
		warnings["skip_dependent_tables"] = "the `skip_dependent_tables` option is ineffective when used with '*' `tables`"
	}

	if slices.Contains(s.Tables, "*") && len(s.Tables) > 1 {
		warnings["all_tables_with_more_tables"] = "`tables` option contains '*' as well as other tables. '*' will match all tables"
	}

	return warnings
}

func (s *Source) SetDefaults() {
	s.Metadata.SetDefaults()
	if s.Spec == nil {
		s.Spec = make(map[string]any)
	}
	if s.SkipDependentTables == nil {
		b := true
		s.SkipDependentTables = &b
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

func (Source) JSONSchemaExtend(sc *jsonschema.Schema) {
	tables := sc.Properties.Value("tables")
	*tables = *tables.OneOf[0] // only value

	destinations := sc.Properties.Value("destinations")
	*destinations = *destinations.OneOf[0] // only value

	Metadata{}.JSONSchemaExtend(sc) // have to call manually
}

func (s *Source) Validate() error {
	if err := s.Metadata.Validate(); err != nil {
		return err
	}

	if len(s.Tables) == 0 {
		return errors.New("tables configuration is required. Hint: set the tables you want to sync by adding `tables: [...]` or use `cloudquery tables` to list available tables")
	}

	if len(s.Destinations) == 0 {
		return errors.New("at least one destination is required")
	}

	return nil
}
