package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v4"
)

type Spec struct {
	ConnectionString string   `json:"connection_string,omitempty"`
	PgxLogLevel      LogLevel `json:"pgx_log_level,omitempty"`
}

type Client struct {
	conn                *pgxpool.Pool
	logger              zerolog.Logger
	spec                specs.Destination
	currentDatabaseName string
	currentSchemaName   string
}

type pgTablePrimaryKeys struct {
	name    string
	columns []string
}

type pgColumn struct {
	name string
	typ  string
}

type pgTableColumns struct {
	name    string
	columns []pgColumn
}

const sqlSelectColumnTypes = `SELECT a.attname as column_name, format_type(a.atttypid, a.atttypmod) AS data_type
FROM pg_attribute a JOIN pg_class b ON a.attrelid = b.relfilenode AND relnamespace = (SELECT oid FROM pg_catalog.pg_namespace WHERE nspname = $1)
WHERE a.attnum > 0 -- hide internal columns
AND NOT a.attisdropped -- hide deleted columns
AND b.relname = $2`

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (plugins.DestinationClient, error) {
	c := &Client{
		logger: logger.With().Str("module", "pg-dest").Logger(),
	}
	var specPostgreSql Spec
	c.spec = spec
	if err := spec.UnmarshalSpec(&specPostgreSql); err != nil {
		return nil, fmt.Errorf("failed to unmarshal postgresql spec: %w", err)
	}

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
		conn.ConnInfo().RegisterDataType(pgtype.DataType{Value: &pgxUUID.UUID{}, Name: "uuid", OID: pgtype.UUIDOID})
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

	c.currentDatabaseName, err = c.currentDatabase()
	if err != nil {
		return nil, fmt.Errorf("failed to get current database: %w", err)
	}
	c.currentSchemaName = "public"
	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	if c.conn == nil {
		return fmt.Errorf("client already closed or not initialized")
	}
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	return nil
}

func (c *Client) currentDatabase() (string, error) {
	var db string
	err := c.conn.QueryRow(context.Background(), "select current_database()").Scan(&db)
	if err != nil {
		return "", err
	}
	return db, nil
}

func (c *Client) getPgTableColumns(ctx context.Context, tableName string) (*pgTableColumns, error) {
	tc := pgTableColumns{
		name: tableName,
	}
	rows, err := c.conn.Query(ctx, sqlSelectColumnTypes, c.currentSchemaName, tableName)
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

func (c *pgTablePrimaryKeys) columnExist(column string) bool {
	for _, col := range c.columns {
		if col == column {
			return true
		}
	}
	return false
}

func (c *pgTableColumns) getPgColumn(column string) *pgColumn {
	for _, col := range c.columns {
		if col.name == column {
			return &col
		}
	}
	return nil
}
