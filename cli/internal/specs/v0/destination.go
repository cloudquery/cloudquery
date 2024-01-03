package specs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Destination struct {
	Metadata

	WriteMode   WriteMode      `json:"write_mode,omitempty"`
	MigrateMode MigrateMode    `json:"migrate_mode,omitempty"`
	PKMode      PKMode         `json:"pk_mode,omitempty"`
	Spec        map[string]any `json:"spec,omitempty"`
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

func (d Destination) VersionString() string {
	if d.Registry != RegistryGithub {
		return fmt.Sprintf("%s (%s@%s)", d.Name, d.Registry, d.Path)
	}
	pathParts := strings.Split(d.Path, "/")
	if len(pathParts) != 2 {
		return fmt.Sprintf("%s (%s@%s)", d.Name, d.Path, d.Version)
	}
	if d.Name == pathParts[1] {
		return fmt.Sprintf("%s (%s)", d.Name, d.Version)
	}
	return fmt.Sprintf("%s (%s@%s)", d.Name, pathParts[1], d.Version)
}
