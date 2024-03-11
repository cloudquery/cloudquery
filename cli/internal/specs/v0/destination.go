package specs

import (
	"bytes"
	"encoding/json"
)

// Destination plugin spec
type Destination struct {
	Metadata

	// Destination plugin write mode
	WriteMode WriteMode `json:"write_mode,omitempty" jsonschema:"default=overwrite-delete-stale"`

	// Destination plugin migrate mode
	MigrateMode MigrateMode `json:"migrate_mode,omitempty" jsonschema:"default=safe"`

	// Destination plugin PK mode
	PKMode PKMode `json:"pk_mode,omitempty" jsonschema:"default=default"`

	SyncSummary bool `json:"summary,omitempty" jsonschema:"default=false"`

	// Destination plugin own (nested) spec
	Spec map[string]any `json:"spec,omitempty"`
}

func (*Destination) GetWarnings() Warnings {
	warnings := make(map[string]string)
	return warnings
}

func (d *Destination) SetDefaults() {
	d.Metadata.SetDefaults()
	if d.Spec == nil {
		d.Spec = make(map[string]any)
	}
}

func (d *Destination) UnmarshalSpec(out any) error {
	b, err := json.Marshal(d.Spec)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	dec.DisallowUnknownFields()
	return dec.Decode(out)
}

func (d *Destination) Validate() error {
	return d.Metadata.Validate()
}
