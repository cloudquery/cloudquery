package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource

	logger zerolog.Logger
	spec   *Spec

	*filetypes.Client
	writer *StreamingBatchWriter
}

func New(_ context.Context, logger zerolog.Logger, spec []byte) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "file").Logger(),
	}
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file spec: %w", err)
	}

	// if c.spec.WriteMode != specs.WriteModeAppend {
	//	return nil, fmt.Errorf("file destination only supports append mode")
	// }

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

	if err != nil {
		return nil, fmt.Errorf("failed to create filetype client: %w", err)
	}

	c.writer, err = NewStreamingBatchWriter(c, WithStreamingBatchWriterBatchSize(c.spec.BatchSize), WithStreamingBatchWriterBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (*Client) Close(_ context.Context) error {
	return nil
}
