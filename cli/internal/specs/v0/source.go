package specs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
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

// Sync modes supported by the `sync_mode` option.
const (
	// SyncModeFull syncs every selected table on every run.
	SyncModeFull = "full"
	// SyncModeEvents syncs change events from the source, optionally after a full sync.
	SyncModeEvents = "events"
)

// AllSyncModes lists the accepted values of `sync_mode`.
var AllSyncModes = []string{SyncModeFull, SyncModeEvents}

// Policies supported by the `events.full_sync` option.
const (
	// FullSyncAlways runs a full sync at the start of every events run.
	FullSyncAlways = "always"
	// FullSyncOnFirstRun runs a full sync only when the source has no checkpoint yet.
	FullSyncOnFirstRun = "on_first_run"
	// FullSyncNever never runs a full sync, even on the first run.
	FullSyncNever = "never"
)

// AllFullSyncPolicies lists the accepted values of `events.full_sync`.
var AllFullSyncPolicies = []string{FullSyncAlways, FullSyncOnFirstRun, FullSyncNever}

// Defaults applied by EventsOptions.SetDefaults.
const (
	defaultBatchMaxChanges = 500
	defaultBatchMaxWait    = 5 * time.Second
	defaultSummaryInterval = 5 * time.Minute
)

// EventsOptions tunes `sync_mode: events`. It is only valid with that sync mode.
type EventsOptions struct {
	// Follow keeps the sync running and streaming changes until it is stopped. When `false` the sync exits once the source is caught up.
	Follow bool `json:"follow,omitempty" jsonschema:"default=false"`

	// FullSync decides when a full sync runs before the CLI starts following the change stream.
	FullSync string `json:"full_sync,omitempty" jsonschema:"enum=always,enum=on_first_run,enum=never,default=on_first_run"`

	// MaxCheckpointAge forces a full sync when the source's checkpoint is older than this duration. Unset means the checkpoint never expires.
	MaxCheckpointAge *configtype.Duration `json:"max_checkpoint_age,omitempty"`

	// BatchMaxChanges is the number of changes that closes a refresh batch.
	BatchMaxChanges int `json:"batch_max_changes,omitempty" jsonschema:"default=500,minimum=1"`

	// BatchMaxWait is how long a refresh batch waits for more changes after the first change of the batch arrived.
	BatchMaxWait *configtype.Duration `json:"batch_max_wait,omitempty" jsonschema:"default=5s"`

	// SummaryInterval is how often a sync summary is written while following the change stream.
	SummaryInterval *configtype.Duration `json:"summary_interval,omitempty" jsonschema:"default=5m"`
}

func (e *EventsOptions) SetDefaults() {
	if e.FullSync == "" {
		e.FullSync = FullSyncOnFirstRun
	}
	if e.BatchMaxChanges == 0 {
		e.BatchMaxChanges = defaultBatchMaxChanges
	}
	if e.BatchMaxWait == nil {
		e.BatchMaxWait = ptrDuration(defaultBatchMaxWait)
	}
	if e.SummaryInterval == nil {
		e.SummaryInterval = ptrDuration(defaultSummaryInterval)
	}
}

func (e *EventsOptions) Validate() error {
	if !slices.Contains(AllFullSyncPolicies, e.FullSync) {
		return fmt.Errorf("invalid events.full_sync %q, must be one of %v", e.FullSync, AllFullSyncPolicies)
	}
	if e.BatchMaxChanges <= 0 {
		return errors.New("events.batch_max_changes must be greater than 0")
	}
	if e.BatchMaxWait != nil && e.BatchMaxWait.Duration() <= 0 {
		return errors.New("events.batch_max_wait must be greater than 0")
	}
	if e.SummaryInterval != nil && e.SummaryInterval.Duration() <= 0 {
		return errors.New("events.summary_interval must be greater than 0")
	}
	if e.MaxCheckpointAge != nil && e.MaxCheckpointAge.Duration() <= 0 {
		return errors.New("events.max_checkpoint_age must be greater than 0")
	}
	return nil
}

func ptrDuration(d time.Duration) *configtype.Duration {
	v := configtype.NewDuration(d)
	return &v
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

	// SyncMode selects how the source is synced: `full` syncs every selected table, `events` follows the source's change events and refreshes only the resources that changed.
	SyncMode string `json:"sync_mode,omitempty" jsonschema:"enum=full,enum=events,default=full"`

	// Events tunes `sync_mode: events`. It is only valid with that sync mode.
	Events *EventsOptions `json:"events,omitempty"`

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
	if s.SyncMode == "" {
		s.SyncMode = SyncModeFull
	}
	if s.SyncMode == SyncModeEvents {
		if s.Events == nil {
			s.Events = &EventsOptions{}
		}
		s.Events.SetDefaults()
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

func (s *Source) Validate(relaxed bool) error {
	if err := s.Metadata.Validate(); err != nil {
		return err
	}

	if len(s.Tables) == 0 {
		return errors.New("tables configuration is required. Hint: set the tables you want to sync by adding `tables: [...]` or use `cloudquery tables` to list available tables")
	}

	if !relaxed && len(s.Destinations) == 0 {
		return errors.New("at least one destination is required")
	}

	return s.validateSyncMode()
}

func (s *Source) validateSyncMode() error {
	switch s.SyncMode {
	case "", SyncModeFull:
		if s.Events != nil {
			return errors.New("`events` options require `sync_mode: events`")
		}
		return nil
	case SyncModeEvents:
		if s.BackendOptions == nil || s.BackendOptions.TableName == "" {
			return errors.New("`sync_mode: events` requires `backend_options` with a `table_name`. Hint: point `backend_options.connection` at one of your destinations")
		}
		if s.Events == nil {
			return nil
		}
		return s.Events.Validate()
	default:
		return fmt.Errorf("invalid sync_mode %q, must be one of %v", s.SyncMode, AllSyncModes)
	}
}
