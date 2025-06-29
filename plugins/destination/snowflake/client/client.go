package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/rs/zerolog"

	_ "github.com/snowflakedb/gosnowflake" // "snowflake" database/sql driver.
)

var errInvalidSpec = errors.New("invalid spec")

type Client struct {
	plugin.UnimplementedSource
	batchwriter.UnimplementedDeleteRecord
	db     *sql.DB
	logger zerolog.Logger
	spec   Spec
	writer *batchwriter.BatchWriter

	setupWriteOnce *sync.Once
}

func New(_ context.Context, logger zerolog.Logger, spec []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var err error
	c := &Client{
		logger:         logger.With().Str("module", "sf-dest").Logger(),
		setupWriteOnce: &sync.Once{},
	}
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, errors.Join(errInvalidSpec, err)
	}
	c.spec.SetDefaults()
	c.writer, err = batchwriter.New(c, batchwriter.WithLogger(c.logger), batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, errors.Join(errInvalidSpec, err)
	}
	dsn, err := c.spec.DSN()
	if err != nil {
		return nil, errors.Join(errInvalidSpec, err)
	}
	if strings.Contains(dsn, "token=token") {
		// if DSN contains token=token, it means we are running in a Snowflake Container Service
		// https://docs.snowflake.com/en/developer-guide/snowpark-container-services/overview
		// we need to read the token that is automatically mounted and use it (/snowflake/session/token)
		if _, err := os.Stat("/snowflake/session/token"); err == nil {
			token, err := os.ReadFile("/snowflake/session/token")
			if err != nil {
				return nil, errors.Join(errInvalidSpec, err)
			}
			// the token contains "/" that we need to escape because we use a connection string - ideally we would want to move from connection string in the config
			escapedToken := url.QueryEscape(string(token))
			dsn = strings.Replace(dsn, "token=token", fmt.Sprintf("token=%s", string(escapedToken)), 1)
		} else {
			return nil, errors.New("token not found in Snowflake Container Service at /snowflake/session/token")
		}
	}

	db, err := sql.Open("snowflake", dsn+"&BINARY_INPUT_FORMAT=BASE64&BINARY_OUTPUT_FORMAT=BASE64&timezone=UTC")
	if err != nil {
		return nil, errors.Join(errInvalidSpec, err, fmt.Errorf("dsn is %s", dsn))
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Join(fmt.Errorf("ping failed"), err, fmt.Errorf("dsn is %s", dsn))
	}

	c.db = db
	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	if c.db == nil {
		return errors.New("client already closed or not initialized")
	}

	if err := c.writer.Close(ctx); err != nil {
		_ = c.db.Close()
		c.db = nil
		return fmt.Errorf("failed to close writer: %w", err)
	}

	err := c.db.Close()
	c.db = nil
	return err
}
