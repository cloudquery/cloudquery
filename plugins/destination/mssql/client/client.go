package client

import (
	"context"
	"database/sql"
	"errors"

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
	batchwriter.UnimplementedDeleteRecord
}

var _ plugin.Client = (*Client)(nil)
var _ batchwriter.Client = (*Client)(nil)

var errInvalidSpec = errors.New("invalid spec")

func (c *Client) Close(ctx context.Context) error {
	if err := c.writer.Close(ctx); err != nil {
		_ = c.db.Close()
		return err
	}
	return c.db.Close()
}

func New(_ context.Context, logger zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var spec Spec
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		return nil, errors.Join(errInvalidSpec, err)
	}
	spec.SetDefaults()

	if err := spec.Validate(); err != nil {
		return nil, errors.Join(errInvalidSpec, err)
	}

	connector, err := spec.Connector()
	if err != nil {
		return nil, errors.Join(errInvalidSpec, err)
	}

	c := &Client{
		logger: logger.With().Str("module", "dest-mssql").Logger(),
		spec:   spec,
	}
	// set ctx logger
	mssql.SetContextLogger(c)
	c.db = sql.OpenDB(connector)

	err = c.db.Ping()
	if err != nil {
		return nil, err
	}

	c.writer, err = batchwriter.New(c,
		batchwriter.WithLogger(c.logger),
		batchwriter.WithBatchSize(spec.BatchSize),
		batchwriter.WithBatchSizeBytes(spec.BatchSizeBytes),
		batchwriter.WithBatchTimeout(spec.BatchTimeout.Duration()),
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}
