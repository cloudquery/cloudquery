package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/queries"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (c *Client) createTable(ctx context.Context, sc *arrow.Schema) (err error) {
	name := schema.TableName(sc)
	c.logger.Debug().Str("table", name).Msg("Table doesn't exist, creating")
	defer func() {
		if err != nil {
			c.logErr(err)
		}
	}()

	_, err = c.db.ExecContext(ctx, queries.CreateTable(c.schemaName, sc, c.pkEnabled()))
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", name, err)
	}

	return c.ensureTVP(ctx, sc)
}

func (c *Client) dropTable(ctx context.Context, sc *arrow.Schema) (err error) {
	name := schema.TableName(sc)
	c.logger.Debug().Str("table", name).Msg("Table exists, dropping")
	defer func() {
		if err != nil {
			c.logErr(err)
		}
	}()

	_, err = c.db.ExecContext(ctx, queries.DropTable(c.schemaName, sc))
	if err != nil {
		return fmt.Errorf("failed to drop table %s: %w", name, err)
	}

	return nil
}

func (c *Client) recreateTable(ctx context.Context, sc *arrow.Schema) error {
	err := c.dropTable(ctx, sc)
	if err != nil {
		return err
	}

	return c.createTable(ctx, sc)
}
