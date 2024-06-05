package client

import (
	_ "embed"
	"errors"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/invopop/jsonschema"
	mssql "github.com/microsoft/go-mssqldb"
	"github.com/microsoft/go-mssqldb/azuread"
)

type AuthMode string

const (
	AuthModeAzure = AuthMode("azure")
	AuthModeMS    = AuthMode("ms")
)

type Spec struct {
	// Connection string to connect to the database.
	// See [SDK documentation](https://github.com/microsoft/go-mssqldb#connection-parameters-and-dsn) for details.
	//
	// Example connection strings:
	// - `"sqlserver://username:password@hostname/instance"` basic connection using a named instance
	// - `"sqlserver://username:password@localhost?database=master&connection+timeout=30"` select "master" database and set connection timeout (default instance)
	ConnectionString string `json:"connection_string" jsonschema:"required,minLength=1,example=${MSSQL_CONNECTION_STRING}"`

	//  If you need to authenticate via Azure Active Directory ensure you specify `azure` value.
	//  See [SDK documentation](https://github.com/microsoft/go-mssqldb#azure-active-directory-authentication) for more information.
	//  Supported values:
	//
	//    - `ms` _connect to Microsoft SQL Server instance_
	//    - `azure` _connect to Azure SQL Server instance_
	AuthMode AuthMode `json:"auth_mode,omitempty" jsonschema:"default=ms"`

	// By default, Microsoft SQL Server destination plugin will use the [default](https://learn.microsoft.com/en-us/sql/relational-databases/security/authentication-access/ownership-and-user-schema-separation?view=sql-server-ver16#the-dbo-schema) schema named `dbo`.
	Schema string `json:"schema,omitempty" jsonschema:"default=dbo"`

	// Maximum number of items that may be grouped together to be written in a single write.
	BatchSize int `json:"batch_size,omitempty" jsonschema:"minimum=1,default=1000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	BatchSizeBytes int `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=5242880"`

	// Timeout for writing a single batch.
	BatchTimeout *configtype.Duration `json:"batch_timeout,omitempty"`
}

//go:embed schema.json
var JSONSchema string

func (s *Spec) Validate() error {
	if len(s.ConnectionString) == 0 {
		return errors.New("missing required \"connection_string\" option")
	}
	return nil
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
	if strings.EqualFold(string(s.AuthMode), string(AuthModeAzure)) {
		return azuread.NewConnector(s.ConnectionString)
	}
	return mssql.NewConnector(s.ConnectionString)
}

func (AuthMode) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Type = "string"
	sc.Enum = []any{AuthModeAzure, AuthModeMS}
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	batchTimeout := sc.Properties.Value("batch_timeout").OneOf[0] // 0 - val, 1 - null
	batchTimeout.Default = "20s"
}
