package spec

import (
	"errors"
	"time"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

type Spec struct {
	filetypes.FileSpec

	// Storage account where to sync the files.
	StorageAccount string `json:"storage_account,omitempty" jsonschema:"required,minLength=1"`

	// Storage container inside the storage account where to sync the files.
	Container string `json:"container,omitempty" jsonschema:"required,minLength=1"`

	// Path to where the files will be uploaded in the storage container.
	Path string `json:"path,omitempty" jsonschema:"required,minLength=1"`

	// If set to `true`, the plugin will write to one file per table.
	// Otherwise, for every batch a new file will be created with a different `.<UUID>` suffix.
	NoRotate bool `json:"no_rotate,omitempty" jsonschema:"default=false"`

	// Maximum number of items that may be grouped together to be written in a single object.
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
	if len(s.StorageAccount) == 0 {
		return errors.New("`storage_account` is required")
	}
	if len(s.Container) == 0 {
		return errors.New("`container` is required")
	}
	if len(s.Path) == 0 {
		return errors.New("`path` is required")
	}

	if s.NoRotate && ((s.BatchSize != nil && *s.BatchSize > 0) || (s.BatchSizeBytes != nil && *s.BatchSizeBytes > 0) || (s.BatchTimeout != nil && s.BatchTimeout.Duration() > 0)) {
		return errors.New("`no_rotate` cannot be used with non-zero `batch_size`, `batch_size_bytes` or `batch_timeout`")
	}

	// required for s.FileSpec.Validate call
	err := s.FileSpec.UnmarshalSpec()
	if err != nil {
		return err
	}
	s.FileSpec.SetDefaults()

	return s.FileSpec.Validate()
}

func int64ptr(i int64) *int64 {
	return &i
}
