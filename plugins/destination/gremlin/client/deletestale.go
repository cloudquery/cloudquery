package client

import (
	"context"
	"time"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Schemas, source string, syncTime time.Time) error {
	session, closer, err := c.newSession()
	if err != nil {
		return err
	}
	defer closer()

	for _, table := range tables {
		tableName := schema.TableName(table)
		g := gremlingo.Traversal_().WithRemote(session).
			V().
			HasLabel(tableName).
			Has(schema.CqSourceNameColumn.Name, source).
			Has(schema.CqSyncTimeColumn.Name, gremlingo.P.Lt(syncTime)).
			SideEffect(AnonT.Drop())
		if err := <-g.Iterate(); err != nil {
			return err
		}
	}
	return nil
}
