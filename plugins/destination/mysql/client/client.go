package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/rs/zerolog"

	mysql "github.com/go-sql-driver/mysql"
)

type ServerType int64

const (
	ServerTypeMySQL   ServerType = 0
	ServerTypeMariaDB ServerType = 1
)

// Client is the MySQL client

type Client struct {
	plugin.UnimplementedSource
	logger        zerolog.Logger
	spec          Spec
	db            *sql.DB
	writer        *batchwriter.BatchWriter
	serverType    ServerType
	serverVersion string

	maxIndexLength int
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{logger: logger.With().Str("module", "mysql").Logger()}
	var err error

	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	c.spec.SetDefaults()
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.writer, err = batchwriter.New(c, batchwriter.WithLogger(logger), batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create batch writer: %w", err)
	}

	dsn, err := mysql.ParseDSN(c.spec.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("invalid MySQL connection string: %w", err)
	}
	if dsn.Params == nil {
		dsn.Params = map[string]string{}
	}
	dsn.Params["parseTime"] = "true"
	db, err := sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open mysql connection: %w", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	c.db = db

	if err := c.validateConnection(ctx); err != nil {
		return nil, err
	}

	if err := c.getVersion(ctx); err != nil {
		return nil, err
	}

	c.setMaxIndexLength(ctx)

	return c, nil
}

func (c *Client) validateConnection(ctx context.Context) error {
	rows, err := c.db.QueryContext(ctx, "select database()")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var name *string
		if err := rows.Scan(&name); err != nil {
			return err
		}
		if name == nil {
			return fmt.Errorf("default database is not selected. Update connection string to include database name")
		}
	}
	return nil
}

func (c *Client) getVersion(ctx context.Context) error {
	rows, err := c.db.QueryContext(ctx, "SELECT VERSION()")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var versionString *string
		if err := rows.Scan(&versionString); err != nil {
			return err
		}
		if strings.Contains(*versionString, "-MariaDB") {
			c.serverType = ServerTypeMariaDB
			c.logger.Warn().Msg("MariaDB detected. Some features may not work as expected")
		}
		c.serverVersion = strings.Split(*versionString, "-")[0]
	}
	return nil
}

func (c *Client) setMaxIndexLength(ctx context.Context) {
	const maxIndexLengthInBytes = 3072
	row := c.db.QueryRowContext(ctx, "show variables like 'innodb_default_row_format'")
	var varName sql.NullString
	var rowFormat sql.NullString
	err := row.Scan(&varName, &rowFormat)
	if err != nil {
		c.logger.Warn().Err(err).Msgf("failed to detect max index length, using default value of %d bytes", maxIndexLengthInBytes)
	}

	// In MySQL >= 5.7 the max PK length is 3072 bytes for dynamic or compressed row format, and 767 bytes for redundant or compact row format.
	// We need to divide the max length in bytes by 4, since we use utf8mb4 charset, which can take up to 4 bytes per character.
	switch rowFormat.String {
	case "redundant", "compact":
		c.maxIndexLength = 767 / 4
	default:
		c.maxIndexLength = maxIndexLengthInBytes / 4
	}
}

func (c *Client) Close(ctx context.Context) error {
	if err := c.writer.Close(ctx); err != nil {
		_ = c.db.Close()
		return fmt.Errorf("failed to close writer: %w", err)
	}
	return c.db.Close()
}
