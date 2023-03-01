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
	database string
	cluster  string
	conn     clickhouse.Conn

	logger zerolog.Logger

	spec specs.Destination

	destination.UnimplementedUnmanagedWriter
}

var _ destination.Client = (*Client)(nil)

func (*Client) DeleteStale(context.Context, schema.Tables, string, time.Time) error {
	return errors.New("DeleteStale is not implemented")
}

func (c *Client) Close(context.Context) error {
	return c.conn.Close()
}

func New(_ context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("clickhouse destination only supports append mode")
	}

	var pluginSpec Spec
	if err := spec.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	options, err := pluginSpec.Options()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare connect options %w", err)
	}

	l := logger.With().
		Str("module", "dest-clickhouse").
		Str("database", options.Auth.Database).
		Str("cluster", pluginSpec.Cluster).
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
		database: options.Auth.Database,
		cluster:  pluginSpec.Cluster,
		conn:     conn,
		logger:   l,
		spec:     spec,
	}, nil
}
