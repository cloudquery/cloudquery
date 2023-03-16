package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	conn     clickhouse.Conn
	database string
	spec     *Spec

	logger zerolog.Logger
	mode   specs.MigrateMode
	destination.UnimplementedUnmanagedWriter
}

var _ destination.Client = (*Client)(nil)

func (*Client) DeleteStale(context.Context, schema.Tables, string, time.Time) error {
	return errors.New("DeleteStale is not implemented")
}

func (c *Client) Close(context.Context) error {
	return c.conn.Close()
}

func New(_ context.Context, logger zerolog.Logger, dstSpec specs.Destination) (destination.Client, error) {
	if dstSpec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("clickhouse destination only supports append mode")
	}

	var spec Spec
	if err := dstSpec.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	options, err := spec.Options()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare connect options %w", err)
	}

	l := logger.With().
		Str("module", "dest-clickhouse").
		Str("database", options.Auth.Database).
		Str("cluster", spec.Cluster).
		Logger()
	options.Debugf = l.Printf

	conn, err := clickhouse.Open(options)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare connection %w", err)
	}

	ver, err := conn.ServerVersion()
	if err != nil {
		return nil, fmt.Errorf("failed to verify server version %w", err)
	}

	minVer := proto.Version{Major: 22, Minor: 1, Patch: 2}
	if !proto.CheckMinVersion(minVer, ver.Version) {
		defer conn.Close()
		return nil, fmt.Errorf("server version is %s, minimum version supported is %s", ver.Version, minVer)
	}

	return &Client{
		conn:     conn,
		database: options.Auth.Database,
		spec:     &spec,
		logger:   l,
		mode:     dstSpec.MigrateMode,
	}, nil
}
