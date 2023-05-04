package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	mssql "github.com/microsoft/go-mssqldb"
	"github.com/rs/zerolog"
)

type Client struct {
	db         *sql.DB
	schemaName string

	logger zerolog.Logger

	spec specs.Destination

	destination.UnimplementedUnmanagedWriter
}

var _ destination.Client = (*Client)(nil)

func (c *Client) Close(context.Context) error {
	return c.db.Close()
}

func New(_ context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	var pluginSpec Spec
	if err := spec.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	pluginSpec.SetDefaults()

	connector, err := pluginSpec.Connector()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare connection %w", err)
	}

	c := &Client{
		schemaName: pluginSpec.Schema,
		logger:     logger.With().Str("module", "dest-mssql").Logger(),
		spec:       spec,
	}
	// set ctx logger
	mssql.SetContextLogger(c)
	c.db = sql.OpenDB(connector)

	return c, nil
}
