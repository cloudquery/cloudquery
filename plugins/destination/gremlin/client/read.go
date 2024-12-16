package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) Read(_ context.Context, table *schema.Table, res chan<- arrow.Record) error {
	session, closer, err := c.newSession()
	if err != nil {
		return err
	}
	defer closer()

	g := gremlingo.Traversal_().WithRemote(session).
		V().
		HasLabel(table.Name).
		Group().By(gremlingo.T.Id).
		By(AnonT.ValueMap())

	rs, err := g.GetResultSet()
	if err != nil {
		return fmt.Errorf("GetResultSet: %w", err)
	}
	defer rs.Close()

	sc := table.ToArrowSchema()
	for row := range rs.Channel() {
		m := row.Data.(map[any]any)
		for _, rowCols := range m {
			rowData := rowCols.(map[any]any)
			rec, err := reverseTransformer(sc, rowData)
			if err != nil {
				return err
			}
			res <- rec
		}
	}

	return rs.GetError()
}
