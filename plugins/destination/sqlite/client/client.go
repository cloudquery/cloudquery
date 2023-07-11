package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"

	// Import sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	plugin.UnimplementedSource
	db     *sql.DB
	logger zerolog.Logger
	spec   Spec
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "sqlite-dest").Logger(),
	}

	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	c.spec.SetDefaults()

	db, err := sql.Open("sqlite3", c.spec.ConnectionString)
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
