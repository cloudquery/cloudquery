package client

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	mssql "github.com/microsoft/go-mssqldb"
)

func (*Client) TransformUUID(v *schema.UUID) any {
	if v.Status != schema.Present {
		return nil
	}
	val, _ := mssql.UniqueIdentifier(v.Bytes).Value()
	return val
}
