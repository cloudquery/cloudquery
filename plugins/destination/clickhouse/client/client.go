package client

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
)

type Client struct {
	conn     clickhouse.Conn
	database string
	spec     *Spec

	logger zerolog.Logger
	writer *batchwriter.BatchWriter
	plugin.UnimplementedSource
}

var _ plugin.Client = (*Client)(nil)
var _ batchwriter.Client = (*Client)(nil)

func (*Client) DeleteStale(context.Context, message.WriteDeleteStales) error {
	return plugin.ErrNotImplemented
}

func (c *Client) Close(ctx context.Context) error {
	if err := c.writer.Close(ctx); err != nil {
		_ = c.conn.Close()
		return err
	}
	return c.conn.Close()
}

func New(_ context.Context, logger zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var spec Spec
	if err := json.Unmarshal(specBytes, &spec); err != nil {
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

	c := &Client{
		conn:     conn,
		database: options.Auth.Database,
		spec:     &spec,
		logger:   l,
	}
	c.writer, err = batchwriter.New(c,
		batchwriter.WithLogger(l),
		batchwriter.WithBatchSize(spec.BatchSize),
		batchwriter.WithBatchSizeBytes(spec.BatchSizeBytes),
		batchwriter.WithBatchTimeout(spec.BatchTimeout.Duration()),
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}
