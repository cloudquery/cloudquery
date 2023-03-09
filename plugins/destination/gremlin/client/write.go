package client

import (
	"context"
	"fmt"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	session, err := c.client.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	rows := make([]map[string]any, len(resources))
	for i, resource := range resources {
		rows[i] = make(map[string]any)
		for j, column := range table.Columns {
			rows[i][column.Name] = resource[j]
		}
	}

	pks := table.PrimaryKeys()
	if len(pks) == 0 {
		// If no primary keys are defined, use all columns
		pks = table.Columns.Names()
	}
	nonPKs := make(map[string]struct{})
	for _, column := range table.Columns {
		nonPKs[column.Name] = struct{}{}
	}
	for _, pk := range pks {
		delete(nonPKs, pk)
	}

	g := gremlingo.Traversal_().WithRemote(session).V().HasLabel(table.Name)
	for i := range rows {
		for _, column := range pks {
			g = g.Has(column, rows[i][column])
		}
		g = g.Fold()

		ins := AnonT.AddV(table.Name)
		for _, column := range pks {
			ins = ins.Property(column, rows[i][column])
		}
		g = g.Coalesce(
			AnonT.Unfold(),
			ins,
		)

		for column := range nonPKs {
			g = g.Property(gremlingo.Cardinality.Single, column, rows[i][column])
		}
	}

	if err := <-g.Iterate(); err != nil {
		return fmt.Errorf("Iterate: %w", err)
	}

	return nil
}
