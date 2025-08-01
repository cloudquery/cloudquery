package client

import (
	"context"
	"fmt"
	"strings"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cenkalti/backoff/v5"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) WriteTableBatch(ctx context.Context, tableName string, msgs message.WriteInserts) error {
	if len(msgs) == 0 {
		return nil
	}

	table, err := schema.NewTableFromArrowSchema(msgs[0].Record.Schema())
	if err != nil {
		return err
	}

	session, closer, err := c.newSession()
	if err != nil {
		return err
	}
	defer closer()

	cqTimeIndex := -1
	for i := range table.Columns {
		if table.Columns[i].Name == schema.CqSyncTimeColumn.Name {
			cqTimeIndex = i
			break
		}
	}

	rows := make([]map[string]any, 0, len(msgs))
	for i := range msgs {
		rows = append(rows, c.transformValues(msgs[i].Record, cqTimeIndex)...)
	}

	pks := table.PrimaryKeys()
	if len(pks) == 0 {
		// If no primary keys are defined, use all columns
		pks = table.Columns.Names()
	}
	valueColumns := make([]string, 0, len(table.Columns)-len(pks))
	if len(table.Columns)-len(pks) > 0 {
		// not all columns are a part of "pk", so we need to account for the values
		for _, col := range table.Columns {
			if !col.PrimaryKey {
				valueColumns = append(valueColumns, col.Name)
			}
		}
	}

	g := gremlingo.Traversal_().WithRemote(session).V().HasLabel(tableName)
	for i := range rows {
		g = g.V().HasLabel(tableName)
		for _, colName := range pks {
			g = g.Has(colName, rows[i][colName])
		}
		g = g.Fold()

		ins := AnonT.AddV(table.Name)
		for _, colName := range pks {
			ins = ins.Property(colName, rows[i][colName])
		}
		g = g.Coalesce(
			AnonT.Unfold(),
			ins,
		)

		for _, colName := range valueColumns {
			g = g.Property(gremlingo.Cardinality.Single, colName, rows[i][colName])
		}
	}

	bo := backoff.WithContext(
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), uint64(c.spec.MaxRetries)),
		ctx,
	)
	return backoff.Retry(func() error {
		err = <-g.Iterate()
		if err == nil {
			return nil
		}
		if !strings.Contains(err.Error(), "ConcurrentModificationException") {
			return backoff.Permanent(fmt.Errorf("Iterate: %w", err))
		}
		return err
	}, bo)
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	if err := c.writer.Write(ctx, msgs); err != nil {
		return err
	}
	if err := c.writer.Flush(ctx); err != nil {
		return fmt.Errorf("failed to flush: %w", err)
	}
	return nil
}
