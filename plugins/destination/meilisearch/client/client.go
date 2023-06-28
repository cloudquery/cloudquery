package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/meilisearch/meilisearch-go"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	Meilisearch *meilisearch.Client

	logger   zerolog.Logger
	spec     Spec
	pkColumn string
	writer   *batchwriter.BatchWriter
}

var _ plugin.Client = (*Client)(nil)

func (c *Client) Close(ctx context.Context) error {
	if err := c.writer.Close(ctx); err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}
	c.Meilisearch = nil
	return nil
}

func (*Client) DeleteStale(context.Context, []*message.WriteDeleteStale) error {
	return fmt.Errorf("DeleteStale not supported")
}

func (c *Client) verifyVersion() error {
	version, err := c.Meilisearch.Version()
	if err != nil {
		return err
	}

	parts := strings.Split(version.PkgVersion, ".")
	if len(parts) < 3 {
		return fmt.Errorf("malformed version %q (expected \"major.minor.patch\"", version)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("failed to parse major version (%q): %w", parts[0], err)
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("failed to parse minor version (%q): %w", parts[1], err)
	}

	const (
		minMajor = 1
		minMinor = 1
	)

	if (major > minMajor) || (major == minMajor && minor >= minMinor) {
		return nil
	}

	return fmt.Errorf("unsupported Meilisearch version %s (must be >= 1.1)", version.PkgVersion)
}

func New(_ context.Context, logger zerolog.Logger, specBytes []byte) (plugin.Client, error) {
	var pkColumn string

	spec := Spec{}
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	spec.setDefaults()
	if err := spec.validate(); err != nil {
		return nil, err
	}

	mClient, err := spec.getClient()
	if err != nil {
		return nil, err
	}

	client := &Client{
		Meilisearch: mClient,
		logger:      logger.With().Str("module", "dest-meilisearch").Str("host", spec.Host).Logger(),
		pkColumn:    pkColumn,
		spec:        spec,
	}
	writer, err := batchwriter.New(client, batchwriter.WithBatchSize(spec.BatchSize), batchwriter.WithBatchSizeBytes(spec.BatchSizeBytes))
	if err != nil {
		return nil, err
	}
	client.writer = writer

	return client, client.verifyVersion()
}
