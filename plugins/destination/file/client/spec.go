package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type FormatType int

type BackendType int

const (
	remoteBackendMaxFileSize = 1024 * 1024 * 4 // 4MB
)

const (
	BackendTypeLocal BackendType = iota
	BackendTypeGCS
	BackendTypeS3
)

const (
	FormatTypeCSV  FormatType = iota
	FormatTypeJSON
)

type Spec struct {
	Directory   string      `json:"directory,omitempty"`
	Backend     BackendType `json:"backend,omitempty"`
	Format      FormatType  `json:"format,omitempty"`
	MaxFileSize uint64      `json:"max_file_size,omitempty"`
}

func (s *Spec) SetDefaults() {
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

func (r BackendType) String() string {
	return [...]string{"local", "gcs", "s3"}[r]
}

func (r BackendType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(r.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (r *BackendType) UnmarshalJSON(data []byte) (err error) {
	var backendType string
	if err := json.Unmarshal(data, &backendType); err != nil {
		return err
	}
	if *r, err = BackendTypeFromString(backendType); err != nil {
		return err
	}
	return nil
}

func BackendTypeFromString(s string) (BackendType, error) {
	switch s {
	case "local":
		return BackendTypeLocal, nil
	case "gcs":
		return BackendTypeGCS, nil
	case "s3":
		return BackendTypeS3, nil
	default:
		return BackendTypeLocal, fmt.Errorf("invalid level %s", s)
	}
}

func (r FormatType) String() string {
	return [...]string{"csv", "json"}[r]
}

func (r FormatType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(r.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (r *FormatType) UnmarshalJSON(data []byte) (err error) {
	var formatType string
	if err := json.Unmarshal(data, &formatType); err != nil {
		return err
	}
	if *r, err = FormatTypeFromString(formatType); err != nil {
		return err
	}
	return nil
}

func FormatTypeFromString(s string) (FormatType, error) {
	switch s {
	case "csv":
		return FormatTypeCSV, nil
	case "json":
		return FormatTypeJSON, nil
	default:
		return FormatTypeJSON, fmt.Errorf("invalid level %s", s)
	}
}