package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/transformer/test/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	plugin.UnimplementedDestination

	logger zerolog.Logger
	spec   spec.Spec
}

func New(_ context.Context, logger zerolog.Logger, s []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", opts.PluginMeta.Name).Logger(),
	}
	if opts.NoConnection {
		return c, nil
	}

	if err := json.Unmarshal(s, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file spec: %w", err)
	}
	c.spec.SetDefaults()
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Transform(ctx context.Context, recvRecords <-chan arrow.Record, sendRecords chan<- arrow.Record) error {
	sourceRecords := 0
	for {
		select {
		case record, ok := <-recvRecords:
			if !ok {
				return nil
			}

			sourceRecords++
			if c.spec.FailImmediately || (c.spec.FailAfterNSourceRecords > 0 && sourceRecords > c.spec.FailAfterNSourceRecords) {
				return fmt.Errorf("failing at the transformer stage according to spec requirements")
			}
			if c.spec.ExitImmediately {
				os.Exit(1)
			}

			sendRecords <- record
		case <-ctx.Done():
			return nil
		}
	}
}

func (*Client) TransformSchema(ctx context.Context, schema *arrow.Schema) (*arrow.Schema, error) {
	return schema, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}
