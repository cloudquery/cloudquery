package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/mixedbatchwriter"
	pgx_zero_log "github.com/jackc/pgx-zerolog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/otan/gopgkrb5"
	"github.com/rs/zerolog"
)

// This holds the details of the PK Constraint
type pkConstraintDetails struct {
	name    string
	columns []string
}

type Client struct {
	conn                *pgxpool.Pool
	logger              zerolog.Logger
	currentDatabaseName string
	currentSchemaName   string
	pgType              pgType
	batchSize           int64
	writer              *mixedbatchwriter.MixedBatchWriter

	spec *spec.Spec

	pgTablesToPKConstraints   map[string]*pkConstraintDetails
	pgTablesToPKConstraintsMu sync.RWMutex

	plugin.UnimplementedSource
}

// Assert Client implements plugin.Client interface.
var _ plugin.Client = (*Client)(nil)

type pgType int

const (
	invalid pgType = iota
	pgTypePostgreSQL
	pgTypeCockroachDB
	pgTypeCrateDB
)

func init() {
	pgconn.RegisterGSSProvider(func() (pgconn.GSS, error) { return gopgkrb5.NewGSS() })
}

func New(ctx context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "pg-dest").Logger(),
	}
	if opts.NoConnection {
		return c, nil
	}

	var s spec.Spec
	if err := json.Unmarshal(specBytes, &s); err != nil {
		return nil, err
	}
	s.SetDefaults()
	if err := s.Validate(); err != nil {
		return nil, err
	}
	c.spec = &s
	c.batchSize = s.BatchSize
	c.logger.Info().Str("pgx_log_level", s.PgxLogLevel.String()).Msg("Initializing postgresql destination")

	pgxConfig, err := pgxpool.ParseConfig(s.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string %w", err)
	}
	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return nil
	}

	pgxConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgx_zero_log.NewLogger(c.logger),
		LogLevel: s.PgxLogLevel.LogLevel(),
	}
	// maybe expose this to the user?
	pgxConfig.ConnConfig.RuntimeParams["timezone"] = "UTC"
	c.conn, err = pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgresql: %w", err)
	}

	c.currentDatabaseName, err = currentDatabase(ctx, c.conn)
	if err != nil {
		return nil, fmt.Errorf("failed to get current database: %w", err)
	}
	c.currentSchemaName, err = currentSchema(ctx, c.conn)
	if err != nil {
		return nil, fmt.Errorf("failed to get current schema: %w", err)
	}
	c.pgType, err = c.getPgType(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get database type: %w", err)
	}
	c.writer, err = mixedbatchwriter.New(c,
		mixedbatchwriter.WithLogger(c.logger),
		mixedbatchwriter.WithBatchSize(s.BatchSize),
		mixedbatchwriter.WithBatchSizeBytes(s.BatchSizeBytes),
		mixedbatchwriter.WithBatchTimeout(s.BatchTimeout.Duration()),
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
	if c.conn == nil {
		return errors.New("client already closed or not initialized")
	}

	c.conn.Close()
	c.conn = nil
	return nil
}

func currentDatabase(ctx context.Context, conn *pgxpool.Pool) (string, error) {
	var db string
	err := conn.QueryRow(ctx, "select current_database()").Scan(&db)
	if err != nil {
		return "", err
	}
	return db, nil
}

func currentSchema(ctx context.Context, conn *pgxpool.Pool) (string, error) {
	var schema string
	err := conn.QueryRow(ctx, "select current_schema()").Scan(&schema)
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
	case "cratedb":
		typ = pgTypeCrateDB
	default:
		return typ, fmt.Errorf("unknown database type %s", name)
	}

	return typ, nil
}
