package client

import (
	"strings"
	"time"

	mssql "github.com/microsoft/go-mssqldb"
	"github.com/microsoft/go-mssqldb/azuread"
)

type AuthMode string

const (
	AuthModeAzure = "azure"
	AuthModeMS    = "ms"
)

type Spec struct {
	ConnectionString string   `json:"connection_string,omitempty"`
	AuthMode         AuthMode `json:"auth_mode,omitempty"`
	Schema           string   `json:"schema,omitempty"`

	BatchSize      int           `json:"batch_size,omitempty"`
	BatchSizeBytes int           `json:"batch_size_bytes,omitempty"`
	BatchTimeout   time.Duration `json:"batch_timeout,omitempty"`
}

func (s *Spec) SetDefaults() {
	const dboSchema = "dbo"
	if len(s.Schema) == 0 {
		s.Schema = dboSchema
	}

	if len(s.AuthMode) == 0 {
		s.AuthMode = AuthModeMS
	}

	if s.BatchSize == 0 {
		const defaultBatchSize = 1000 // 1K
		s.BatchSize = defaultBatchSize
	}

	if s.BatchSizeBytes == 0 {
		const defaultBatchSizeBytes = 5 << 20 // 5 MiB
		s.BatchSizeBytes = defaultBatchSizeBytes
	}

	if s.BatchTimeout == 0 {
		const defaultBatchTimeout = 20 * time.Second // 20s
		s.BatchTimeout = defaultBatchTimeout
	}
}

func (s *Spec) Connector() (*mssql.Connector, error) {
	if strings.EqualFold(string(s.AuthMode), AuthModeAzure) {
		return azuread.NewConnector(s.ConnectionString)
	}
	return mssql.NewConnector(s.ConnectionString)
}
