package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) Read(ctx context.Context, table *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	session, closer, err := c.newSession()
	if err != nil {
		return err
	}
	defer closer()

	tableName := schema.TableName(table)
	g := gremlingo.Traversal_().WithRemote(session).
		V().
		HasLabel(tableName).
		Has(schema.CqSourceNameColumn.Name, sourceName).
		Group().By(gremlingo.T.Id).
		By(AnonT.ValueMap())

	rs, err := g.GetResultSet()
	if err != nil {
		return fmt.Errorf("GetResultSet: %w", err)
	}
	defer rs.Close()

	for row := range rs.Channel() {
		m := row.Data.(map[any]any)
		for _, rowCols := range m {
			rowData := rowCols.(map[any]any)
			rec, err := reverseTransformer(table, rowData)
			if err != nil {
				return err
			}
			res <- rec
		}
	}

	return rs.GetError()
}
