package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (c *Client) useTVP(sc *arrow.Schema) bool {
	return c.pkEnabled() && len(schema.PrimaryKeyIndices(sc)) > 0
}

func (c *Client) ensureTVP(ctx context.Context, sc *arrow.Schema) (err error) {
	if !c.useTVP(sc) {
		return nil
	}

	name := schema.TableName(sc)

	query, params := queries.TVPDropProc(c.schemaName, sc)
	_, err = c.db.ExecContext(ctx, query, params...)
	if err != nil {
		return fmt.Errorf("failed to drop TVP proc for table %s: %w", name, err)
	}

	query, params = queries.TVPDropType(c.schemaName, sc)
	_, err = c.db.ExecContext(ctx, query, params...)
	if err != nil {
		return fmt.Errorf("failed to drop TVP type for table %s: %w", name, err)
	}

	_, err = c.db.ExecContext(ctx, queries.TVPAddType(c.schemaName, sc))
	if err != nil {
		return fmt.Errorf("failed to create TVP type for table %s: %w", name, err)
	}

	_, err = c.db.ExecContext(ctx, queries.TVPAddProc(c.schemaName, sc))
	if err != nil {
		return fmt.Errorf("failed to create TVP proc for table %s: %w", name, err)
	}

	return nil
}

func (c *Client) insertTVP(ctx context.Context, sc *arrow.Schema, records []arrow.Record) error {
	query, params, err := queries.TVPQuery(c.schemaName, sc, records)
	if err != nil {
		return err
	}

	return c.doInTx(ctx, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, query, params...)
		return err
	})
}
