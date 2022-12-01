package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"

	"database/sql"

	"github.com/snowflakedb/gosnowflake"
)

type Client struct {
	plugins.DefaultReverseTransformer
	db      *sql.DB
	logger  zerolog.Logger
	spec    specs.Destination
	metrics plugins.DestinationMetrics
}

func New(ctx context.Context, logger zerolog.Logger, destSpec specs.Destination) (plugins.DestinationClient, error) {
	if destSpec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("snowflake destination only supports append mode")
	}
	c := &Client{
		logger: logger.With().Str("module", "sf-dest").Logger(),
	}
	var spec Spec
	c.spec = destSpec
	if err := destSpec.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal snowflake spec: %w", err)
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}
	_, err := gosnowflake.ParseDSN(spec.ConnectionString)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("snowflake", spec.ConnectionString)
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
