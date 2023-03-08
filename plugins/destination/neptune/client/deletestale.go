package client

import (
	"context"
	"time"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, source string, syncTime time.Time) error {
	session, err := c.client.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	for _, table := range tables {
		g := gremlingo.Traversal_().WithRemote(session).
			V().
			HasLabel(table.Name).
			Has("_cq_source_name", source).
			Has("_cq_sync_time", gremlingo.P.Lt(syncTime)).
			SideEffect(AnonT.Drop())
		if err := <-g.Iterate(); err != nil {
			return err
		}
	}
	return nil
}
