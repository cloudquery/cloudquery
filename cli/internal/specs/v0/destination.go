package specs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/thoas/go-funk"
)

type Destination struct {
	Name        string      `json:"name,omitempty"`
	Version     string      `json:"version,omitempty"`
	Path        string      `json:"path,omitempty"`
	Registry    Registry    `json:"registry,omitempty"`
	WriteMode   WriteMode   `json:"write_mode,omitempty"`
	MigrateMode MigrateMode `json:"migrate_mode,omitempty"`
	// Deprecated: use plugin-spec batch_size_bytes instead
	BatchSize int `json:"batch_size,omitempty"`
	// Deprecated: use plugin-spec batch_size_bytes instead
	BatchSizeBytes int            `json:"batch_size_bytes,omitempty"`
	Spec           map[string]any `json:"spec,omitempty"`
	PKMode         PKMode         `json:"pk_mode,omitempty"`
}

func (d *Destination) GetWarnings() Warnings {
	warnings := make(map[string]string)
	if d.BatchSize != 0 {
		warnings["batch_size"] = "batch_size in the top-level spec is deprecated and will be removed in a future release. Please use the plugin-spec `batch_size` option instead."
	}
	if d.BatchSizeBytes != 0 {
		warnings["batch_size_bytes"] = "batch_size_bytes in the top-level spec is deprecated and will be removed in a future release. Please use the plugin-spec `batch_size_bytes` option instead."
	}
	return warnings
}

func (d *Destination) SetDefaults(defaultBatchSize, defaultBatchSizeBytes int) {
	if d.Spec == nil {
		d.Spec = make(map[string]any)
	}
	if d.Registry.String() == "" {
		d.Registry = RegistryGithub
	}
	if d.BatchSize == 0 {
		d.BatchSize = defaultBatchSize
	}
	if d.BatchSizeBytes == 0 {
		d.BatchSizeBytes = defaultBatchSizeBytes
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
	if d.Name == "" {
		return fmt.Errorf("name is required")
	}
	if d.Path == "" {
		msg := "path is required"
		// give a small hint to help users transition from the old config format that didn't require path
		officialPlugins := []string{"postgresql", "csv"}
		if funk.ContainsString(officialPlugins, d.Name) {
			msg += fmt.Sprintf(". Hint: try setting path to cloudquery/%s in your config", d.Name)
		}
		return fmt.Errorf(msg)
	}

	if d.Registry == RegistryGithub {
		if d.Version == "" {
			return fmt.Errorf("version is required")
		}
		if !strings.HasPrefix(d.Version, "v") {
			return fmt.Errorf("version must start with v")
		}
	}
	if d.BatchSize < 0 {
		return fmt.Errorf("batch_size must be greater than 0")
	}
	return nil
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
