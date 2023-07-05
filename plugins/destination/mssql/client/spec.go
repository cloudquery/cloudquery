package client

import (
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
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

	BatchSize      int                  `json:"batch_size,omitempty"`
	BatchSizeBytes int                  `json:"batch_size_bytes,omitempty"`
	BatchTimeout   *configtype.Duration `json:"batch_timeout,omitempty"`
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
		s.BatchSize = 1000 // 1K
	}

	if s.BatchSizeBytes == 0 {
		s.BatchSizeBytes = 5 << 20 // 5 MiB
	}

	if s.BatchTimeout == nil {
		d := configtype.NewDuration(20 * time.Second) // 20s
		s.BatchTimeout = &d
	}
}

func (s *Spec) Connector() (*mssql.Connector, error) {
	if strings.EqualFold(string(s.AuthMode), AuthModeAzure) {
		return azuread.NewConnector(s.ConnectionString)
	}
	return mssql.NewConnector(s.ConnectionString)
}
