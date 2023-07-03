package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cenkalti/backoff/v4"
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
		for i := range table.Columns {
			pks = append(pks, table.Columns[i].Name)
		}
	}
	nonPKs := make(map[string]struct{})
	for _, c := range table.Columns {
		if !c.PrimaryKey {
			nonPKs[c.Name] = struct{}{}
		}
	}

	g := gremlingo.Traversal_().WithRemote(session).V().HasLabel(tableName)
	for i := range rows {
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

		for colName := range nonPKs {
			g = g.Property(gremlingo.Cardinality.Single, colName, rows[i][colName])
		}
	}

	bo := backoff.NewExponentialBackOff()
	retryCount := 0

	for retryCount <= c.spec.MaxRetries {
		retryCount++

		err = <-g.Iterate()
		if err == nil {
			return nil
		}

		if !strings.Contains(err.Error(), "ConcurrentModificationException") {
			return fmt.Errorf("Iterate: %w", err)
		}

		if retryCount > c.spec.MaxRetries {
			break
		}

		nb := bo.NextBackOff()
		c.logger.Debug().Err(err).Str("backoff_duration", nb.String()).Msg("Iterate failed, retrying")

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(nb):
		}
	}

	return fmt.Errorf("Max retries (%d) reached. Iterate: %w", c.spec.MaxRetries, err)
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
