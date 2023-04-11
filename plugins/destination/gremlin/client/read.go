package client

import (
	"context"
	"fmt"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	session, closer, err := c.newSession()
	if err != nil {
		return err
	}
	defer closer()

	g := gremlingo.Traversal_().WithRemote(session).
		V().
		HasLabel(table.Name).
		Has("_cq_source_name", sourceName).
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

			madeRow := make([]any, len(table.Columns))
			for i, col := range table.Columns {
				if rowData[col.Name] == nil {
					continue
				}

				data := rowData[col.Name].([]any)
				if l := len(data); l == 1 {
					madeRow[i] = data[0]
				} else if l > 1 {
					return fmt.Errorf("expected 1 value for %v, got %v", col.Name, l)
				}
			}
			res <- madeRow
		}
	}

	return rs.GetError()
}
