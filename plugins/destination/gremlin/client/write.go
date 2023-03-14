package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cenkalti/backoff/v4"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	session, closer, err := c.newSession()
	if err != nil {
		return err
	}
	defer closer()

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
