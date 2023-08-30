package client

import (
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/pgarrow"
)

func (c *Client) SchemaTypeToPg(t arrow.DataType) string {
	switch c.pgType {
	case pgTypeCockroachDB:
		return pgarrow.ArrowToCockroach(t)
	default:
		return pgarrow.ArrowToPg10(t)
	}
}

func (c *Client) PgToSchemaType(t string) arrow.DataType {
	switch c.pgType {
	case pgTypeCockroachDB:
		return pgarrow.CockroachToArrow(t)
	default:
		return pgarrow.Pg10ToArrow(t)
	}
}
