package client

import (
	"context"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) DeleteStale(_ context.Context, msgs message.WriteDeleteStales) error {
	session, closer, err := c.newSession()
	if err != nil {
		return err
	}
	defer closer()

	for _, msg := range msgs {
		g := gremlingo.Traversal_().With(session).
			V().
			HasLabel(msg.GetTable().Name).
			Has(schema.CqSourceNameColumn.Name, msg.SourceName).
			Has(schema.CqSyncTimeColumn.Name, gremlingo.P.Lt(msg.SyncTime)).
			SideEffect(AnonT.Drop())
		if err := <-g.Iterate(); err != nil {
			return err
		}
	}

	return nil
}
