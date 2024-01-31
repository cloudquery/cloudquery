package spec

import (
	_ "embed"
	"fmt"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/invopop/jsonschema"
)

const (
	varFormat = "{{FORMAT}}"
	varTable  = "{{TABLE}}"
	varUUID   = "{{UUID}}"
	varYear   = "{{YEAR}}"
	varMonth  = "{{MONTH}}"
	varDay    = "{{DAY}}"
	varHour   = "{{HOUR}}"
	varMinute = "{{MINUTE}}"
)

type Spec struct {
	filetypes.FileSpec

	// Path template string that determines where files will be written.
	//
	// The path supports the following placeholder variables:
	// - `{{TABLE}}` will be replaced with the table name
	// - `{{FORMAT}}` will be replaced with the file format, such as `csv`, `json` or `parquet`. If compression is enabled, the format will be `csv.gz`, `json.gz` etc.
	// - `{{UUID}}` will be replaced with a random UUID to uniquely identify each file
	// - `{{YEAR}}` will be replaced with the current year in `YYYY` format
	// - `{{MONTH}}` will be replaced with the current month in `MM` format
	// - `{{DAY}}` will be replaced with the current day in `DD` format
	// - `{{HOUR}}` will be replaced with the current hour in `HH` format
	// - `{{MINUTE}}` will be replaced with the current minute in `mm` format
	//
	//  **Note** that timestamps are in `UTC` and will be the current time at the time the file is written, not when the sync started.
	Path string `json:"path,omitempty" jsonschema:"required,minLength=1"`

	// If set to `true`, the plugin will write to one file per table.
	// Otherwise, for every batch a new file will be created with a different `.<UUID>` suffix.
	NoRotate bool `json:"no_rotate,omitempty" jsonschema:"default=false"`

	// This parameter controls the maximum amount of items may be grouped together to be written in a single write.
	//
	// Defaults to `10000` unless `no_rotate` is `true` (will be `0` then).
	BatchSize *int64 `json:"batch_size" jsonschema:"minimum=1,default=10000"`

	// This parameter controls the maximum size of items that may be grouped together to be written in a single write.
	//
	// Defaults to `52428800` (50 MiB) unless `no_rotate` is `true` (will be `0` then).
	BatchSizeBytes *int64 `json:"batch_size_bytes" jsonschema:"minimum=1,default=52428800"`

	// This parameter controls the maximum interval between batch writes.
	//
	// Defaults to `30s` unless `no_rotate` is `true` (will be `0s` then).
	BatchTimeout *configtype.Duration `json:"batch_timeout" jsonschema:"default=30s"`
}

func (s *Spec) SetDefaults() {
	if !strings.Contains(s.Path, varTable) {
		s.Path = path.Join(s.Path, fmt.Sprintf("%s.%s", varTable, s.Format))
	}
	if s.BatchSize == nil {
		if s.NoRotate {
			s.BatchSize = ptr(int64(0))
		} else {
			s.BatchSize = ptr(int64(10000))
		}
	}
	if s.BatchSizeBytes == nil {
		if s.NoRotate {
			s.BatchSizeBytes = ptr(int64(0))
		} else {
			s.BatchSizeBytes = ptr(int64(50 * 1024 * 1024)) // 50 MiB
		}
	}
	if s.BatchTimeout == nil {
		if s.NoRotate {
			d := configtype.NewDuration(0)
			s.BatchTimeout = &d
		} else {
			d := configtype.NewDuration(30 * time.Second)
			s.BatchTimeout = &d
		}
	}
}

func (s *Spec) Validate() error {
	if len(s.Path) == 0 {
		return fmt.Errorf("`path` must be set")
	}

	if s.NoRotate && strings.Contains(s.Path, varUUID) {
		return fmt.Errorf("`path` should not contain %s when `no_rotate` = true", varUUID)
	}

	if !strings.Contains(s.Path, varUUID) && s.batchingEnabled() {
		return fmt.Errorf("`path` should contain %s when using a non-zero `batch_size`, `batch_size_bytes` or `batch_timeout_ms`", varUUID)
	}

	if s.NoRotate && ((s.BatchSize != nil && *s.BatchSize > 0) || (s.BatchSizeBytes != nil && *s.BatchSizeBytes > 0) || (s.BatchTimeout != nil && s.BatchTimeout.Duration() > 0)) {
		return fmt.Errorf("`no_rotate` cannot be used with non-zero `batch_size`, `batch_size_bytes` or `batch_timeout_ms`")
	}

	// required for s.FileSpec.Validate call
	err := s.FileSpec.UnmarshalSpec()
	if err != nil {
		return err
	}
	s.FileSpec.SetDefaults()

	return s.FileSpec.Validate()
}

func (s *Spec) ReplacePathVariables(table string, fileIdentifier string, t time.Time) string {
	name := strings.ReplaceAll(s.Path, varTable, table)
	if strings.Contains(name, varFormat) {
		e := string(s.Format) + s.Compression.Extension()
		name = strings.ReplaceAll(name, varFormat, e)
	}
	name = strings.ReplaceAll(name, varUUID, fileIdentifier)
	name = strings.ReplaceAll(name, varYear, t.Format("2006"))
	name = strings.ReplaceAll(name, varMonth, t.Format("01"))
	name = strings.ReplaceAll(name, varDay, t.Format("02"))
	name = strings.ReplaceAll(name, varHour, t.Format("15"))
	name = strings.ReplaceAll(name, varMinute, t.Format("04"))
	return filepath.Clean(name)
}

func (s *Spec) PathContainsUUID() bool {
	return strings.Contains(s.Path, varUUID)
}

func (s *Spec) batchingEnabled() bool {
	if !s.NoRotate && (s.BatchSize == nil || s.BatchSizeBytes == nil || s.BatchTimeout == nil) {
		return true
	}

	return (s.BatchSize != nil && *s.BatchSize > 0) ||
		(s.BatchSizeBytes != nil && *s.BatchSizeBytes > 0) ||
		(s.BatchTimeout != nil && s.BatchTimeout.Duration() > 0)
}

func (s Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	s.FileSpec.JSONSchemaExtend(sc) // need to call manually

	batchTimeout := sc.Properties.Value("batch_timeout").OneOf[0] // 0 - val, 1 - null
	batchTimeout.Default = "30s"
}

//go:embed schema.json
var JSONSchema string

func ptr[A any](a A) *A {
	return &a
}
