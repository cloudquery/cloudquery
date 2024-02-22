package client

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strings"

	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/motherduck/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/marcboeker/go-duckdb"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	batchwriter.UnimplementedDeleteRecord

	connector driver.Connector
	db        *sql.DB
	conn      driver.Conn // used in Appender

	logger zerolog.Logger
	spec   Spec
	writer *batchwriter.BatchWriter
}

var _ plugin.Client = (*Client)(nil)

var localDuckDB bool // used in testing

func New(ctx context.Context, logger zerolog.Logger, spec []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "motherduck-dest").Logger(),
	}
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	c.spec.SetDefaults()
	if err := c.spec.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	c.writer, err = batchwriter.New(c, batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes), batchwriter.WithLogger(c.logger))
	if err != nil {
		return nil, fmt.Errorf("failed to create batch writer: %w", err)
	}

	// Add 'custom_user_agent' to the connection string
	if strings.Contains(c.spec.ConnectionString, "?") {
		c.spec.ConnectionString += "&"
	} else {
		c.spec.ConnectionString += "?"
	}
	c.spec.ConnectionString += fmt.Sprintf("custom_user_agent=%s_%s_%s/%s(%s_%s)",
		internalPlugin.Team, internalPlugin.Kind, internalPlugin.Name, strings.TrimPrefix(internalPlugin.Version, "v"), runtime.GOOS, runtime.GOARCH)

	if len(c.spec.Token) > 0 {
		c.spec.ConnectionString += "&motherduck_token=" + c.spec.Token
	}

	dbUri := c.spec.ConnectionString
	if !localDuckDB {
		dbUri = "md:" + dbUri
	}

	c.connector, err = duckdb.NewConnector(dbUri, nil)
	if err != nil {
		return nil, err
	}
	c.db = sql.OpenDB(c.connector)

	err = c.exec(ctx, "INSTALL 'json'; LOAD 'json';")
	if err != nil {
		return nil, err
	}
	err = c.exec(ctx, "INSTALL 'parquet'; LOAD 'parquet';")
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	if c.db == nil {
		return fmt.Errorf("client already closed or not initialized")
	}

	err1 := c.writer.Close(ctx)
	if err1 != nil {
		err1 = fmt.Errorf("failed to close writer: %w", err1)
	}

	err := errors.Join(err1, c.db.Close(), func() error {
		if c.conn == nil {
			return nil
		}
		return c.conn.Close()
	}())
	c.db, c.conn = nil, nil
	return err
}

func (c *Client) exec(ctx context.Context, query string, args ...any) error {
	r, err := c.db.ExecContext(ctx, query, args...)
	if c.spec.Debug {
		rowsAffected, rowsErr := r.RowsAffected()
		errs := errors.Join(err, rowsErr)
		c.logger.Debug().Str("query", query).Any("values", args).Int64("rowsAffected", rowsAffected).Err(errs).Msg("exec query")
	}
	return err
}
