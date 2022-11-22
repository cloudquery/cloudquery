package client

import (
	"fmt"
)

type FormatType string

type BackendType string


const (
	remoteBackendMaxFileSize = uint64(1024 * 1024 * 4) // 4MB

	BackendTypeLocal = "local"
	BackendTypeGCS = "gcs"
	BackendTypeS3 = "s3"

	FormatTypeCSV = "csv"
	FormatTypeJSON = "json"
)


type Spec struct {
	Directory   string      `json:"directory,omitempty"`
	Backend     BackendType `json:"backend,omitempty"`
	Format      FormatType  `json:"format,omitempty"`
	MaxFileSize uint64      `json:"max_file_size,omitempty"`
	// This is used for debugging purposes only
	NoRotate		bool        `json:"no_rotate,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Backend == "" {
		s.Backend = BackendTypeLocal
	}
	if s.Format == "" {
		s.Format = FormatTypeCSV
	}
	if s.MaxFileSize == 0 {
		s.MaxFileSize = remoteBackendMaxFileSize
	}
	// this debug flag override default rotation
	if s.NoRotate {
		s.MaxFileSize = 0	
	}
}

func (s *Spec) Validate() error {
	if s.Directory == "" {
		return fmt.Errorf("directory is required")
	}
	return nil
}
