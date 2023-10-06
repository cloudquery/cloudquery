package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	pgx_zero_log "github.com/jackc/pgx-zerolog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedDestination
	Conn                *pgxpool.Pool
	logger              zerolog.Logger
	pluginSpec          Spec
	currentDatabaseName string
	currentSchemaName   string
	pgType              pgType
	tables              schema.Tables
	cdcId               string
	options             plugin.NewClientOptions
}

type pgType int

const (
	invalid pgType = iota
	pgTypePostgreSQL
	pgTypeCockroachDB
)

func (*Client) ID() string {
	return "source-pg"
}

func Configure(ctx context.Context, logger zerolog.Logger, spec []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger:  logger,
			options: opts,
			tables:  schema.Tables{},
		}, nil
	}
	c := &Client{
		logger: logger.With().Str("module", "pg-source").Logger(),
	}
	var pluginSpec Spec
	if err := json.Unmarshal(spec, &pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal postgresql spec: %w", err)
	}
	pluginSpec.SetDefaults()

	c.pluginSpec = pluginSpec
	logLevel, err := tracelog.LogLevelFromString(pluginSpec.PgxLogLevel.String())
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgx log level %s: %w", pluginSpec.PgxLogLevel, err)
	}
	c.logger.Info().Str("pgx_log_level", pluginSpec.PgxLogLevel.String()).Msg("Initializing postgresql destination")
	pgxConfig, err := pgxpool.ParseConfig(pluginSpec.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string %w", err)
	}
	if c.pluginSpec.CDCId != "" {
		// if cdc is specified the connection must be in replication mode
		// https://www.postgresql.org/docs/current/libpq-connect.html
		pgxConfig.ConnConfig.RuntimeParams["replication"] = "database"
	}
	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return nil
	}
	pgxConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	pgxConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgx_zero_log.NewLogger(c.logger),
		LogLevel: logLevel,
	}
	// maybe expose this to the user?
	pgxConfig.ConnConfig.RuntimeParams["timezone"] = "UTC"
	c.Conn, err = pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgresql: %w", err)
	}

	c.currentDatabaseName, err = c.currentDatabase(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get current database: %w", err)
	}
	c.currentSchemaName, err = c.currentSchema(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get current schema: %w", err)
	}
	c.pgType, err = c.getPgType(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get database type: %w", err)
	}
	c.tables, err = c.listTables(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}

	if c.pluginSpec.CDCId != "" {
		walLevel, err := c.walLevel(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get wal_level: %w", err)
		}
		if walLevel != "logical" {
			return nil, fmt.Errorf("cdc is enabled but wal_level is not logical")
		}
	}
	c.cdcId = pluginSpec.CDCId
	c.options = opts

	return c, nil
}

func (c *Client) getPgType(ctx context.Context) (pgType, error) {
	var version string
	var typ pgType
	err := c.Conn.QueryRow(ctx, "select version()").Scan(&version)
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

func (c *Client) walLevel(ctx context.Context) (string, error) {
	var walLevel string
	err := c.Conn.QueryRow(ctx, "SELECT setting FROM pg_settings WHERE name='wal_level'").Scan(&walLevel)
	if err != nil {
		return "", err
	}
	return walLevel, nil
}

func (c *Client) currentDatabase(ctx context.Context) (string, error) {
	var db string
	err := c.Conn.QueryRow(ctx, "select current_database()").Scan(&db)
	if err != nil {
		return "", err
	}
	if db == "" {
		return "", fmt.Errorf("failed to get CURRENT_DATABASE")
	}
	return db, nil
}

func (c *Client) currentSchema(ctx context.Context) (string, error) {
	var schemaName string
	err := c.Conn.QueryRow(ctx, "select current_schema()").Scan(&schemaName)
	if err != nil {
		return "", err
	}
	if schemaName == "" {
		return "", fmt.Errorf("failed to get CURRENT_SCHEMA")
	}

	return schemaName, nil
}

func (c Client) Tables(_ context.Context, opts plugin.TableOptions) (schema.Tables, error) {
	if c.options.NoConnection {
		return schema.Tables{}, nil
	}
	return c.tables.FilterDfs(opts.Tables, opts.SkipTables, opts.SkipDependentTables)
}

func (c Client) Close(_ context.Context) error {
	c.Conn.Close()
	return nil
}
