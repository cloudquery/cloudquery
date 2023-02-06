package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/jackc/pglogrepl"
	pgx_zero_log "github.com/jackc/pgx-zerolog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
)

type Client struct {
	Conn                *pgxpool.Pool
	logger              zerolog.Logger
	spec                specs.Source
	pluginSpec          Spec
	currentDatabaseName string
	currentSchemaName   string
	pgType              pgType
	Tables              schema.Tables
	createReplicationSlotResult pglogrepl.CreateReplicationSlotResult
}

type pgType int

const (
	invalid pgType = iota
	pgTypePostgreSQL
	pgTypeCockroachDB
)

var _ schema.ClientMeta = (*Client)(nil)

func (*Client) ID() string {
	return "source-pg"
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	c := &Client{
		logger: logger.With().Str("module", "pg-dest").Logger(),
	}
	var pluginSpec Spec
	c.spec = spec
	if err := spec.UnmarshalSpec(&pluginSpec); err != nil {
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
	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return nil
	}

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
	c.currentSchemaName = "public"
	c.pgType, err = c.getPgType(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get database type: %w", err)
	}
	c.Tables, err = c.listTables(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}

	// if c.pluginSpec.CDC {
	// 	if _, err := c.Conn.Exec(ctx, "CREATE SCHEMA IF NOT EXISTS cq_source_pg_cdc"); err != nil {
	// 		return nil, fmt.Errorf("failed to create cq_source_pg_cdc schema: %w", err)
	// 	}
	// 	if _, err := c.Conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS cq_source_pg_cdc.state (slot_name text PRIMARY KEY, lsn pg_lsn)"); err != nil {
	// 		return nil, fmt.Errorf("failed to create cq_source_pg_cdc.state table: %w", err)
	// 	}
	// }
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

func (c *Client) currentDatabase(ctx context.Context) (string, error) {
	var db string
	err := c.Conn.QueryRow(ctx, "select current_database()").Scan(&db)
	if err != nil {
		return "", err
	}
	return db, nil
}

func (c *Client) createReplication(ctx context.Context) error {
	cfg, err := pgconn.ParseConfig(c.pluginSpec.ConnectionString)
	if err != nil {
		return err
	}
	conn, err := pgconn.ConnectConfig(ctx, cfg)
	if err != nil {
		return err
	}
	tables := strings.Join(c.Tables.TableNames(), ",")
	reader := conn.Exec(ctx, fmt.Sprintf("CREATE PUBLICATION %s FOR TABLE %s;", pgx.Identifier{c.spec.Name}.Sanitize(), tables))
	_, err = reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to create publication: %w", err)
	}

	sysident, err := pglogrepl.IdentifySystem(ctx, conn)
	if err != nil {
		return fmt.Errorf("failed to identify system: %w", err)
	}

	sql := fmt.Sprintf("CREATE_REPLICATION_SLOT %s LOGICAL pgoutput EXPORT_SNAPSHOT", c.spec.Name)
	c.createReplicationSlotResult, err = pglogrepl.ParseCreateReplicationSlot(conn.Exec(ctx, sql))
	if err != nil {
		return fmt.Errorf("failed to create replication slot: %w", err)
	}

	if err := pglogrepl.StartReplication(ctx, conn, c.createReplicationSlotResult.SlotName, sysident.XLogPos,
		pglogrepl.StartReplicationOptions{
			PluginArgs: []string{"proto_version '1'", "publication_names '" + c.spec.Name + "'"},
		}); err != nil {
		return fmt.Errorf("failed to start replication: %w", err)
	}

	return nil
}
