package spec

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
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

	NoRotate  bool   `json:"no_rotate,omitempty"`
	Bucket    string `json:"bucket,omitempty"`
	Region    string `json:"region,omitempty"`
	Path      string `json:"path,omitempty"`
	Athena    bool   `json:"athena,omitempty"`
	TestWrite *bool  `json:"test_write,omitempty"`

	Endpoint              string               `json:"endpoint,omitempty"`
	UsePathStyle          bool                 `json:"use_path_style,omitempty"`
	EndpointSkipTLSVerify bool                 `json:"endpoint_skip_tls_verify,omitempty"`
	BatchSize             *int64               `json:"batch_size"`
	BatchSizeBytes        *int64               `json:"batch_size_bytes"`
	BatchTimeout          *configtype.Duration `json:"batch_timeout"`
}

func (s *Spec) SetDefaults() {
	if !strings.Contains(s.Path, varTable) {
		// for backwards-compatibility, default to given path plus /{{TABLE}}.[format].{{UUID}} if
		// no {{TABLE}} value is found in the path string
		s.Path += fmt.Sprintf("/%s.%s", varTable, s.Format)
		if !s.NoRotate {
			s.Path += "." + varUUID
		}
	}
	if s.TestWrite == nil {
		b := true
		s.TestWrite = &b
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
	if len(s.Bucket) == 0 {
		return fmt.Errorf("`bucket` is required")
	}
	if len(s.Region) == 0 {
		return fmt.Errorf("`region` is required")
	}

	if len(s.Path) == 0 {
		return fmt.Errorf("`path` is required")
	}
	if path.IsAbs(s.Path) {
		return fmt.Errorf("`path` should not start with a \"/\"")
	}
	if s.Path != path.Clean(s.Path) {
		return fmt.Errorf("`path` should not contain relative paths or duplicate slashes")
	}

	if s.NoRotate {
		if strings.Contains(s.Path, varUUID) {
			return fmt.Errorf("`path` should not contain %s when `no_rotate` = true", varUUID)
		}

		if (s.BatchSize != nil && *s.BatchSize > 0) || (s.BatchSizeBytes != nil && *s.BatchSizeBytes > 0) || (s.BatchTimeout != nil && s.BatchTimeout.Duration() > 0) {
			return fmt.Errorf("`no_rotate` cannot be used with non-zero `batch_size`, `batch_size_bytes` or `batch_timeout_ms`")
		}
	}

	if !strings.Contains(s.Path, varUUID) && s.batchingEnabled() {
		return fmt.Errorf("`path` should contain %s when using a non-zero `batch_size`, `batch_size_bytes` or `batch_timeout_ms`", varUUID)
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
	if s.NoRotate {
		// if that's set we don't allow batching
		return false
	}

	return (s.BatchSize == nil || *s.BatchSize > 0) ||
		(s.BatchSizeBytes == nil || *s.BatchSizeBytes > 0) ||
		(s.BatchTimeout == nil || s.BatchTimeout.Duration() > 0)
}

func ptr[A any](a A) *A {
	return &a
}
