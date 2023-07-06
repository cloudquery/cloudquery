package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/mixedbatchwriter"
	pgx_zero_log "github.com/jackc/pgx-zerolog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
)

type Client struct {
	conn                *pgxpool.Pool
	logger              zerolog.Logger
	currentDatabaseName string
	currentSchemaName   string
	pgType              pgType
	batchSize           int
	writer              *mixedbatchwriter.MixedBatchWriter

	plugin.UnimplementedSource
}

// Assert Client implements plugin.Client interface.
var _ plugin.Client = (*Client)(nil)

type pgType int

const (
	invalid pgType = iota
	pgTypePostgreSQL
	pgTypeCockroachDB
)

func New(ctx context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "pg-dest").Logger(),
	}
	if opts.NoConnection {
		return c, nil
	}

	var spec Spec
	err := json.Unmarshal(specBytes, &spec)
	if err != nil {
		return nil, err
	}
	spec.SetDefaults()
	c.batchSize = spec.BatchSize
	logLevel, err := tracelog.LogLevelFromString(spec.PgxLogLevel.String())
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgx log level %s: %w", spec.PgxLogLevel, err)
	}
	c.logger.Info().Str("pgx_log_level", spec.PgxLogLevel.String()).Msg("Initializing postgresql destination")
	pgxConfig, err := pgxpool.ParseConfig(spec.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string %w", err)
	}
	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return nil
	}

	pgxConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgx_zero_log.NewLogger(c.logger),
		LogLevel: logLevel,
	}
	// maybe expose this to the user?
	pgxConfig.ConnConfig.RuntimeParams["timezone"] = "UTC"
	c.conn, err = pgxpool.NewWithConfig(ctx, pgxConfig)
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
	c.writer, err = mixedbatchwriter.New(c,
		mixedbatchwriter.WithLogger(c.logger),
		mixedbatchwriter.WithBatchSize(spec.BatchSize),
		mixedbatchwriter.WithBatchSizeBytes(spec.BatchSizeBytes),
		mixedbatchwriter.WithBatchTimeout(spec.BatchTimeout.Duration()),
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) Write(ctx context.Context, res <-chan message.WriteMessage) error {
	return c.writer.Write(ctx, res)
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

func (c *Client) currentSchema(ctx context.Context) (string, error) {
	var schema string
	err := c.conn.QueryRow(ctx, "select current_schema()").Scan(&schema)
	if err != nil {
		return "", err
	}

	return schema, nil
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
