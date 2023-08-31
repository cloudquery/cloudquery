package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/streamingbatchwriter"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	streamingbatchwriter.IgnoreMigrateTable
	streamingbatchwriter.UnimplementedDeleteStale

	logger zerolog.Logger
	spec   *Spec

	*filetypes.Client
	writer *streamingbatchwriter.StreamingBatchWriter
}

func New(_ context.Context, logger zerolog.Logger, spec []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "file").Logger(),
	}
	if opts.NoConnection {
		return c, nil
	}

	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file spec: %w", err)
	}
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	if c.spec.Directory != "" {
		c.logger.Warn().Msg("deprecated: the `directory` configuration option will be removed in a future version, please use `path` instead")
	}
	c.spec.SetDefaults()

	filetypesClient, err := filetypes.NewClient(c.spec.FileSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to create filetypes client: %w", err)
	}
	c.Client = filetypesClient

	c.writer, err = streamingbatchwriter.New(c,
		streamingbatchwriter.WithBatchSizeRows(*c.spec.BatchSize),
		streamingbatchwriter.WithBatchSizeBytes(*c.spec.BatchSizeBytes),
		streamingbatchwriter.WithBatchTimeout(c.spec.BatchTimeout.Duration()),
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.writer.Close(ctx)
}
