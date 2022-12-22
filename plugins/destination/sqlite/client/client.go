package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"

	"database/sql"

	// Import sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
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
		logger: logger.With().Str("module", "pg-dest").Logger(),
	}

	var sqliteSpec Spec
	c.spec = spec
	if err := spec.UnmarshalSpec(&sqliteSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal postgresql spec: %w", err)
	}
	sqliteSpec.SetDefaults()

	db, err := sql.Open("sqlite3", sqliteSpec.ConnectionString)
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
