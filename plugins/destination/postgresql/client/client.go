package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

type Client struct {
	plugins.DefaultReverseTransformer
	conn                *pgxpool.Pool
	logger              zerolog.Logger
	spec                specs.Destination
	currentDatabaseName string
	currentSchemaName   string
	pgType              pgType
	metrics             plugins.DestinationMetrics
	batchSize           int
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

type pgType int

const (
	invalid pgType = iota
	pgTypePostgreSQL
	pgTypeCockroachDB
)

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

	logLevel, err := pgx.LogLevelFromString(specPostgreSql.PgxLogLevel.String())
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgx log level %s: %w", specPostgreSql.PgxLogLevel, err)
	}
	c.logger.Info().Str("pgx_log_level", specPostgreSql.PgxLogLevel.String()).Msg("Initializing postgresql destination")

	pgxConfig, err := pgxpool.ParseConfig(specPostgreSql.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string %w", err)
	}
	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return nil
	}
	l := zerologadapter.NewLogger(c.logger)
	pgxConfig.ConnConfig.Logger = l
	pgxConfig.ConnConfig.LogLevel = logLevel
	// maybe expose this to the user?
	pgxConfig.ConnConfig.RuntimeParams["timezone"] = "UTC"
	c.conn, err = pgxpool.ConnectConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgresql: %w", err)
	}

	c.currentDatabaseName, err = c.currentDatabase(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get current database: %w", err)
	}
	c.currentSchemaName = "public"
	c.pgType, err = c.getPgType(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get database type: %w", err)
	}
	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	var err error
	if c.conn == nil {
		return fmt.Errorf("client already closed or not initialized")
	}
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	return err
}

func (c *Client) currentDatabase(ctx context.Context) (string, error) {
	var db string
	err := c.conn.QueryRow(ctx, "select current_database()").Scan(&db)
	if err != nil {
		return "", err
	}
	return db, nil
}

func (c *Client) getPgType(ctx context.Context) (pgType, error) {
	var version string
	var typ pgType
	err := c.conn.QueryRow(ctx, "select version()").Scan(&version)
	if err != nil {
		return typ, err
	}
	versionTokens := strings.Split(version, " ")
	if len(versionTokens) == 0 {
		return typ, fmt.Errorf("failed to parse version string %s", version)
	}
	name := strings.ToLower(versionTokens[0])
	switch name {
	case "postgresql":
		typ = pgTypePostgreSQL
	case "cockroachdb":
		typ = pgTypeCockroachDB
	default:
		return typ, fmt.Errorf("unknown database type %s", name)
	}

	return typ, nil
}

func (c *Client) getPgTableColumns(ctx context.Context, tableName string) (*pgTableColumns, error) {
	tc := pgTableColumns{
		name: tableName,
	}
	rows, err := c.conn.Query(ctx, sqlSelectColumnTypes, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var typ string
		if err := rows.Scan(&name, &typ); err != nil {
			return nil, err
		}
		tc.columns = append(tc.columns, pgColumn{
			name: strings.ToLower(name),
			typ:  strings.ToLower(typ),
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &tc, nil
}

func (c *pgTableColumns) getPgColumn(column string) *pgColumn {
	for _, col := range c.columns {
		if col.name == column {
			return &col
		}
	}
	return nil
}
