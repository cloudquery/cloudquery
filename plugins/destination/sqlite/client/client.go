package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/rs/zerolog"

	// Import sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	destination.UnimplementedManagedWriter
	db      *sql.DB
	logger  zerolog.Logger
	spec    specs.Destination
	metrics destination.Metrics
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "sqlite-dest").Logger(),
	}

	var sqliteSpec Spec
	c.spec = spec
	if err := spec.UnmarshalSpec(&sqliteSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal sqlite spec: %w", err)
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
