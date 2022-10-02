package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgx/v4"
)

const sqlDropTable = "drop table if exists "

func (p *Client) Drop(ctx context.Context, tables schema.Tables) error {
	p.logger.Info().Strs("tables", tables.TableNames()).Msg("Dropping tables")
	for _, table := range tables {
		var tableName pgx.Identifier = []string{table.Name}
		if _, err := p.conn.Exec(ctx, sqlDropTable+tableName.Sanitize()); err != nil {
			return fmt.Errorf("failed to drop table %s: %w", table.Name, err)
		}
	}
	return nil
}
