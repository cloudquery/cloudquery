package client

import (
	"github.com/apache/arrow/go/v12/arrow"
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

func reverseTransform(sc *arrow.Schema, data []any) (arrow.Record, error) {
	return nil, nil
}
