package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/goccy/go-json"
	mssql "github.com/microsoft/go-mssqldb"
	"github.com/rs/zerolog"
)

type Client struct {
	db   *sql.DB
	spec Spec

	logger zerolog.Logger
	writer *batchwriter.BatchWriter
	plugin.UnimplementedSource
}

var _ plugin.Client = (*Client)(nil)
var _ batchwriter.Client = (*Client)(nil)

func (c *Client) Close(ctx context.Context) error {
	if err := c.writer.Close(ctx); err != nil {
		_ = c.db.Close()
		return err
	}
	return c.db.Close()
}

func New(_ context.Context, logger zerolog.Logger, specBytes []byte) (plugin.Client, error) {
	var spec Spec
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	spec.SetDefaults()

	connector, err := spec.Connector()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare connection %w", err)
	}

	c := &Client{
		logger: logger.With().Str("module", "dest-mssql").Logger(),
		spec:   spec,
	}
	// set ctx logger
	mssql.SetContextLogger(c)
	c.db = sql.OpenDB(connector)

	return c, nil
}
