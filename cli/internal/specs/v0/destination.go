package specs

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"
)

const (
	varYear   = "{{YEAR}}"
	varMonth  = "{{MONTH}}"
	varDay    = "{{DAY}}"
	varHour   = "{{HOUR}}"
	varMinute = "{{MINUTE}}"
	varSyncId = "{{SYNC_ID}}"
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

	SyncGroupId string `json:"sync_group_id,omitempty"`

	SyncSummary bool `json:"send_sync_summary,omitempty"`

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

func (d *Destination) RenderedSyncGroupId(t time.Time, sync_id string) string {
	renderedValue := strings.ReplaceAll(d.SyncGroupId, varYear, t.Format("2006"))
	renderedValue = strings.ReplaceAll(renderedValue, varMonth, t.Format("01"))
	renderedValue = strings.ReplaceAll(renderedValue, varDay, t.Format("02"))
	renderedValue = strings.ReplaceAll(renderedValue, varHour, t.Format("15"))
	renderedValue = strings.ReplaceAll(renderedValue, varMinute, t.Format("04"))
	renderedValue = strings.ReplaceAll(renderedValue, varSyncId, sync_id)
	return renderedValue
}
