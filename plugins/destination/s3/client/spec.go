package client

import (
	"fmt"
	"path"
	"strings"

	"github.com/cloudquery/filetypes/v3"
)

type Spec struct {
	*filetypes.FileSpec
	NoRotate  bool   `json:"no_rotate,omitempty"`
	Bucket    string `json:"bucket,omitempty"`
	Region    string `json:"region,omitempty"`
	Path      string `json:"path,omitempty"`
	Athena    bool   `json:"athena,omitempty"`
	TestWrite *bool  `json:"test_write,omitempty"`
	Endpoint  string `json:"endpoint,omitempty"`
}

func (s *Spec) SetDefaults() {
	if !strings.Contains(s.Path, PathVarTable) {
		// for backwards-compatibility, default to given path plus /{{TABLE}}.[format].{{UUID}} if
		// no {{TABLE}} value is found in the path string
		s.Path += fmt.Sprintf("/%s.%s", PathVarTable, s.Format)
		if !s.NoRotate {
			s.Path += "." + PathVarUUID
		}
	}
	if s.TestWrite == nil {
		b := true
		s.TestWrite = &b
	}
}

func (s *Spec) Validate() error {
	if s.Bucket == "" {
		return fmt.Errorf("bucket is required")
	}
	if s.Path == "" {
		return fmt.Errorf("path is required")
	}
	if s.Region == "" {
		return fmt.Errorf("region is required")
	}
	if s.NoRotate && strings.Contains(s.Path, PathVarUUID) {
		return fmt.Errorf("path should not contain %s when no_rotate = true", PathVarUUID)
	}
	if path.IsAbs(s.Path) {
		return fmt.Errorf(`path should not start with a "/"`)
	}

	if s.Path != path.Clean(s.Path) {
		return fmt.Errorf("path should not contain relative paths or duplicate slashes")
	}

	if s.Format == "" {
		return fmt.Errorf("format is required")
	}

	return nil
}
