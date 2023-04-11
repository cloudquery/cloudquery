package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"github.com/rs/zerolog"

	// import duckdb driver
	_ "github.com/marcboeker/go-duckdb"
)

type Client struct {
	destination.UnimplementedManagedWriter
	destination.DefaultReverseTransformer
	db      *sql.DB
	logger  zerolog.Logger
	spec    specs.Destination
	metrics destination.Metrics
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "duckdb-dest").Logger(),
	}

	var duckdbSpec Spec
	c.spec = spec
	if err := spec.UnmarshalSpec(&duckdbSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal duckdb spec: %w", err)
	}
	duckdbSpec.SetDefaults()

	db, err := sql.Open("duckdb", duckdbSpec.ConnectionString)
	if err != nil {
		return nil, err
	}
	c.db = db
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
