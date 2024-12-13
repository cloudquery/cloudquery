package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/spec"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/transformers"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	plugin.UnimplementedDestination

	logger zerolog.Logger
	spec   spec.Spec
	tfs    []*transformers.Transformer
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

	tf, err := transformers.NewFromSpec(c.spec)
	if err != nil {
		return nil, err
	}
	c.tfs = append(c.tfs, tf)

	return c, nil
}

func (c *Client) Transform(ctx context.Context, recvRecords <-chan arrow.Record, sendRecords chan<- arrow.Record) error {
	for {
		select {
		case record, ok := <-recvRecords:
			if !ok {
				return nil
			}

			// Run all transformers sequentially on the record
			for _, tf := range c.tfs {
				var err error
				record, err = tf.Transform(record)
				if err != nil {
					return err
				}
			}

			sendRecords <- record
		case <-ctx.Done():
			return nil
		}
	}
}

func (c *Client) TransformSchema(ctx context.Context, schema *arrow.Schema) (*arrow.Schema, error) {
	for _, tf := range c.tfs {
		var err error
		schema, err = tf.TransformSchema(schema)
		if err != nil {
			return nil, err
		}
	}
	return schema, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}
