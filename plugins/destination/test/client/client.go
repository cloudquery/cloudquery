package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Client struct {
	destination.DefaultReverseTransformer
	logger zerolog.Logger
	spec   Spec
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	var testConfig Spec
	err := spec.UnmarshalSpec(&testConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	return &Client{
		logger: log.With().Str("module", "test").Logger(),
		spec:   testConfig,
	}, nil
}

func (*Client) Metrics() destination.Metrics {
	return destination.Metrics{}
}

func (*Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	return nil
}

func (*Client) Migrate(ctx context.Context, tables schema.Tables) error {
	return nil
}

//revive:disable We need to range over the channel to clear it, but revive thinks it can be removed
func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *destination.ClientResource) error {
	if c.spec.ErrorOnWrite {
		return errors.New("error_on_write is true")
	}
	for range res {
		// do nothing
	}
	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, res [][]any) error {
	if c.spec.ErrorOnWrite {
		return errors.New("error_on_write is true")
	}
	return nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}

func (*Client) DeleteStale(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) error {
	return nil
}

func (*Client) ReverseTransformValues(table *schema.Table, values []any) (schema.CQTypes, error) {
	return nil, nil
}
