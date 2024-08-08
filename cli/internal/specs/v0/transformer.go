package specs

import (
	"bytes"
	"encoding/json"
)

// Transformer plugin spec
type Transformer struct {
	Metadata

	// Transformer plugin own (nested) spec
	Spec map[string]any `json:"spec,omitempty"`
}

func (*Transformer) GetWarnings() Warnings {
	warnings := make(map[string]string)
	return warnings
}

func (d *Transformer) UnmarshalSpec(out any) error {
	b, err := json.Marshal(d.Spec)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	dec.DisallowUnknownFields()
	return dec.Decode(out)
}

func (d *Transformer) Validate() error {
	return d.Metadata.Validate()
}
