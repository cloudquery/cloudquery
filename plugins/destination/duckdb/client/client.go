package client

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
	"github.com/rs/zerolog"

	// import duckdb driver
	"github.com/marcboeker/go-duckdb"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	db        *sql.DB
	connector driver.Connector
	logger    zerolog.Logger
	spec      specs.Destination
	metrics   destination.Metrics
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "duckdb-dest").Logger(),
	}

	var duckdbSpec Spec
	c.spec = spec
	if err := spec.UnmarshalSpec(&duckdbSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal duckdb spec: %w", err)
	}
	c.connector, err = duckdb.NewConnector(duckdbSpec.ConnectionString, nil)
	db := sql.OpenDB(c.connector)
	if err != nil {
		return nil, err
	}
	c.db = db
	_, err = c.db.Exec("INSTALL 'json'; LOAD 'json';")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	var err error
	if c.db == nil {
		return fmt.Errorf("client already closed or not initialized")
	}
	err = c.db.Close()
	c.db = nil
	return err
}
