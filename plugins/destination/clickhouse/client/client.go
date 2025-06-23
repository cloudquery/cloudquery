package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
)

type Client struct {
	conn     clickhouse.Conn
	database string
	spec     *spec.Spec

	logger zerolog.Logger
	writer *batchwriter.BatchWriter
	plugin.UnimplementedSource
}

var _ plugin.Client = (*Client)(nil)
var _ batchwriter.Client = (*Client)(nil)

func (c *Client) Close(ctx context.Context) error {
	if err := c.writer.Close(ctx); err != nil {
		_ = c.conn.Close()
		return err
	}
	return c.conn.Close()
}

var errInvalidSpec = errors.New("invalid spec")

func New(_ context.Context, logger zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var s spec.Spec
	if err := json.Unmarshal(specBytes, &s); err != nil {
		return nil, errors.Join(errInvalidSpec, err)
	}
	s.SetDefaults()
	if err := s.Validate(); err != nil {
		return nil, errors.Join(errInvalidSpec, err)
	}

	options, err := s.Options()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare connect options %w", err)
	}

	l := logger.With().
		Str("module", "dest-clickhouse").
		Str("database", options.Auth.Database).
		Str("cluster", s.Cluster).
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

	minVer := proto.Version{Major: 24, Minor: 8, Patch: 1}
	if !proto.CheckMinVersion(minVer, ver.Version) {
		defer conn.Close()
		return nil, fmt.Errorf("server version is %s, minimum version supported is %s", ver.Version, minVer)
	}

	c := &Client{
		conn:     conn,
		database: options.Auth.Database,
		spec:     &s,
		logger:   l,
	}
	c.writer, err = batchwriter.New(c,
		batchwriter.WithLogger(c.logger),
		batchwriter.WithBatchSize(s.BatchSize),
		batchwriter.WithBatchSizeBytes(s.BatchSizeBytes),
		batchwriter.WithBatchTimeout(s.BatchTimeout.Duration()),
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}
