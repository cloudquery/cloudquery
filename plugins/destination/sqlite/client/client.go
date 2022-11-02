package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	db        *sql.DB
	logger    zerolog.Logger
	spec      specs.Destination
	metrics   plugins.DestinationMetrics
	batchSize int
}

type pgColumn struct {
	name string
	typ  string
}

type pgTableColumns struct {
	name    string
	columns []pgColumn
}

const sqlSelectColumnTypes = `
SELECT
pg_attribute.attname AS column_name,
pg_catalog.format_type(pg_attribute.atttypid, pg_attribute.atttypmod) AS data_type
FROM
pg_catalog.pg_attribute
INNER JOIN
pg_catalog.pg_class ON pg_class.oid = pg_attribute.attrelid
INNER JOIN
pg_catalog.pg_namespace ON pg_namespace.oid = pg_class.relnamespace
WHERE
pg_attribute.attnum > 0
AND NOT pg_attribute.attisdropped
AND pg_class.relname = $1
ORDER BY
attnum ASC;
`


func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (plugins.DestinationClient, error) {
	c := &Client{
		logger: logger.With().Str("module", "pg-dest").Logger(),
	}

	var specPostgreSql Spec
	c.spec = spec
	if err := spec.UnmarshalSpec(&specPostgreSql); err != nil {
		return nil, fmt.Errorf("failed to unmarshal postgresql spec: %w", err)
	}
	specPostgreSql.SetDefaults()
	c.batchSize = specPostgreSql.BatchSize

	db, err := sql.Open("sqlite3", specPostgreSql.ConnectionString)
	if err != nil {
		return nil, err
	}
	c.db = db
	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	var err error
	if c.db == nil {
		return fmt.Errorf("client already closed or not initialized")
	}
	err = c.db.Close()
	c.db = nil
	return err
}

// func (c *Client) getPgTableColumns(ctx context.Context, tableName string) (*pgTableColumns, error) {
// 	tc := pgTableColumns{
// 		name: tableName,
// 	}
// 	rows, err := c.db.Query(ctx, sqlSelectColumnTypes, tableName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var name string
// 		var typ string
// 		if err := rows.Scan(&name, &typ); err != nil {
// 			return nil, err
// 		}
// 		tc.columns = append(tc.columns, pgColumn{
// 			name: strings.ToLower(name),
// 			typ:  strings.ToLower(typ),
// 		})
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return &tc, nil
// }

func (c *pgTableColumns) getPgColumn(column string) *pgColumn {
	for _, col := range c.columns {
		if col.name == column {
			return &col
		}
	}
	return nil
}
