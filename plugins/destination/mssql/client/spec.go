package client

import (
	"strings"

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
}

func (s *Spec) SetDefaults() {
	const dboSchema = "dbo"
	if len(s.Schema) == 0 {
		s.Schema = dboSchema
	}

	if len(s.AuthMode) == 0 {
		s.AuthMode = AuthModeMS
	}
}

func (s *Spec) Connector() (*mssql.Connector, error) {
	if strings.EqualFold(string(s.AuthMode), AuthModeAzure) {
		return azuread.NewConnector(s.ConnectionString)
	}
	return mssql.NewConnector(s.ConnectionString)
}
