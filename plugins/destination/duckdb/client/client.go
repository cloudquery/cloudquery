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

	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/duckdb/v5/resources/plugin"
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

	logger zerolog.Logger
	spec   Spec
	writer *batchwriter.BatchWriter
}

var _ plugin.Client = (*Client)(nil)

func New(ctx context.Context, logger zerolog.Logger, spec []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "duckdb-dest").Logger(),
	}
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	c.spec.SetDefaults()
	c.writer, err = batchwriter.New(c, batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes), batchwriter.WithLogger(c.logger))
	if err != nil {
		return nil, fmt.Errorf("failed to create batch writer: %w", err)
	}

	c.connector, err = duckdb.NewConnector(amendConnectionString(c.spec.ConnectionString), nil)
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
		return errors.New("client already closed or not initialized")
	}

	err1 := c.writer.Close(ctx)
	if err1 != nil {
		err1 = fmt.Errorf("failed to close writer: %w", err1)
	}

	err := errors.Join(err1, c.db.Close())
	c.db = nil
	return err
}

func TestConnection(ctx context.Context, _ zerolog.Logger, specBytes []byte) error {
	var s Spec
	if err := json.Unmarshal(specBytes, &s); err != nil {
		return &plugin.TestConnError{
			Code:    "INVALID_SPEC",
			Message: fmt.Errorf("failed to unmarshal spec: %w", err),
		}
	}
	s.SetDefaults()

	connector, err := duckdb.NewConnector(amendConnectionString(s.ConnectionString), nil)
	if err != nil {
		return err
	}
	db := sql.OpenDB(connector)
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return err
	}

	return db.Close()
}

func (c *Client) exec(ctx context.Context, query string, args ...any) error {
	r, err := c.db.ExecContext(ctx, query, args...)
	if c.spec.Debug {
		logEvent := c.logger.Debug().Str("query", query).Any("values", args)
		if err != nil {
			logEvent.Err(err).Msg("exec query")
		} else {
			rowsAffected, rowsErr := r.RowsAffected()
			logEvent.Int64("rowsAffected", rowsAffected).Err(rowsErr).Msg("exec query")
		}
	}
	return err
}

func amendConnectionString(s string) string {
	if !strings.HasPrefix(s, "md:") {
		return s
	}

	// Motherduck, add 'custom_user_agent' to the connection string
	if strings.Contains(s, "?") {
		s += "&"
	} else {
		s += "?"
	}
	s += fmt.Sprintf("custom_user_agent=%s_%s_%s/%s(%s_%s)",
		internalPlugin.Team, internalPlugin.Kind, internalPlugin.Name, strings.TrimPrefix(internalPlugin.Version, "v"), runtime.GOOS, runtime.GOARCH)
	return s
}
