package spec

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

type Spec struct {
	filetypes.FileSpec

	// Bucket where to sync the files.
	Bucket string `json:"bucket,omitempty" jsonschema:"required,minLength=1"`

	// Path to where the files will be uploaded in the above bucket.
	Path string `json:"path,omitempty" jsonschema:"required,minLength=1"`

	// If set to `true`, the plugin will write to one file per table.
	// Otherwise, for every batch a new file will be created with a different `.<UUID>` suffix.
	NoRotate bool `json:"no_rotate,omitempty" jsonschema:"default=false"`

	// Maximum number of items may be grouped together to be written in a single object.
	//
	// Defaults to `10000` unless `no_rotate` is `true` (will be `0` then).
	BatchSize *int64 `json:"batch_size" jsonschema:"minimum=1,default=10000"`

	// Maximum size of items that may be grouped together to be written in a single object.
	//
	// Defaults to `52428800` (50 MiB) unless `no_rotate` is `true` (will be `0` then).
	BatchSizeBytes *int64 `json:"batch_size_bytes" jsonschema:"minimum=1,default=52428800"`

	// Maximum interval between batch writes.
	//
	// Defaults to `30s` unless `no_rotate` is `true` (will be `0s` then).
	BatchTimeout *configtype.Duration `json:"batch_timeout" jsonschema:"default=30s"`

	// GCP service account key content.
	// This allows for using different service accounts for the GCP source and GCS destination.
	// If using service account keys, it is best to use [environment or file variable substitution](/docs/advanced-topics/environment-variable-substitution).
	ServiceAccountKeyJSON string `json:"service_account_key_json"`
}

func (s *Spec) SetDefaults() {
	if s.BatchSize == nil {
		if s.NoRotate {
			s.BatchSize = int64ptr(0)
		} else {
			s.BatchSize = int64ptr(10000)
		}
	}
	if s.BatchSizeBytes == nil {
		if s.NoRotate {
			s.BatchSizeBytes = int64ptr(0)
		} else {
			s.BatchSizeBytes = int64ptr(50 * 1024 * 1024) // 50 MiB
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
	if len(s.Bucket) == 0 {
		return fmt.Errorf("`bucket` is required")
	}
	if len(s.Path) == 0 {
		return fmt.Errorf("`path` is required")
	}

	if s.NoRotate && ((s.BatchSize != nil && *s.BatchSize > 0) || (s.BatchSizeBytes != nil && *s.BatchSizeBytes > 0) || (s.BatchTimeout != nil && s.BatchTimeout.Duration() > 0)) {
		return fmt.Errorf("`no_rotate` cannot be used with non-zero `batch_size`, `batch_size_bytes` or `batch_timeout_ms`")
	}

	if len(s.ServiceAccountKeyJSON) > 0 {
		if err := isValidJson(s.ServiceAccountKeyJSON); err != nil {
			return fmt.Errorf("invalid json for service_account_key_json: %w", err)
		}
	}

	// required for s.FileSpec.Validate call
	err := s.FileSpec.UnmarshalSpec()
	if err != nil {
		return err
	}
	s.FileSpec.SetDefaults()

	return s.FileSpec.Validate()
}

func isValidJson(content string) error {
	var v map[string]any
	err := json.Unmarshal([]byte(content), &v)
	if err != nil {
		return err
	}
	return nil
}

func int64ptr(i int64) *int64 {
	return &i
}
