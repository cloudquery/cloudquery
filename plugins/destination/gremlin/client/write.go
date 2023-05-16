package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cenkalti/backoff/v4"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, records []arrow.Record) error {
	session, closer, err := c.newSession()
	if err != nil {
		return err
	}
	defer closer()

	rows := make([]map[string]any, 0)
	for _, record := range records {
		rows = append(rows, transformValues(record)...)
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

	g := gremlingo.Traversal_().WithRemote(session).V().HasLabel(table.Name)
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

	for retryCount <= c.pluginSpec.MaxRetries {
		retryCount++

		err = <-g.Iterate()
		if err == nil {
			return nil
		}

		if !strings.Contains(err.Error(), "ConcurrentModificationException") {
			return fmt.Errorf("Iterate: %w", err)
		}

		if retryCount > c.pluginSpec.MaxRetries {
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

	return fmt.Errorf("Max retries (%d) reached. Iterate: %w", c.pluginSpec.MaxRetries, err)
}
