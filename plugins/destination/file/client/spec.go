package client

import "fmt"

type BackendType string
type FormatType string

const (
	remoteBackendMaxFileSize = 1024 * 1024 * 4 // 4MB
)

const (
	BackendTypeLocal BackendType = "local"
	BackendTypeS3    BackendType = "s3"
	BackendTypeGCS   BackendType = "gcs"

	FormatTypeCSV  FormatType = "csv"
	FormatTypeJSON FormatType = "json"
)

type Spec struct {
	Directory   string      `json:"directory,omitempty"`
	Backend     BackendType `json:"backend,omitempty"`
	Format      FormatType  `json:"format,omitempty"`
	MaxFileSize uint64      `json:"max_file_size,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Backend == "" {
		s.Backend = BackendTypeLocal
	}
	if s.Format == "" {
		s.Format = FormatTypeCSV
	}

	if s.Backend != BackendTypeLocal && s.MaxFileSize == 0 {
		s.MaxFileSize = remoteBackendMaxFileSize
	}
}

func (s *Spec) Validate() error {
	if s.Directory == "" {
		return fmt.Errorf("directory is required")
	}
	return nil
}
