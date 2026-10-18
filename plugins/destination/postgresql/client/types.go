package client

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/pgarrow"
)

func (c *Client) SchemaTypeToPg(t arrow.DataType) string {
	return c.pgType.schemaTypeToPg(t)
}

func (c *Client) PgToSchemaType(t string) arrow.DataType {
	return c.pgType.pgToSchemaType(t)
}

func (p pgType) schemaTypeToPg(t arrow.DataType) string {
	switch p {
	case pgTypeCockroachDB:
		return pgarrow.ArrowToCockroach(t)
	case pgTypeCrateDB:
		return pgarrow.ArrowToCrateDB(t)
	default:
		return pgarrow.ArrowToPg10(t)
	}
}

func (p pgType) pgToSchemaType(t string) arrow.DataType {
	switch p {
	case pgTypeCockroachDB:
		return pgarrow.CockroachToArrow(t)
	case pgTypeCrateDB:
		return pgarrow.CrateDBToArrow(t)
	default:
		return pgarrow.Pg10ToArrow(t)
	}
}
