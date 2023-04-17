package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	spec   Spec
}

func New(_ context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	var testConfig Spec
	err := spec.UnmarshalSpec(&testConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	return &Client{
		logger: logger.With().Str("module", "test").Logger(),
		spec:   testConfig,
	}, nil
}

func (*Client) Metrics() destination.Metrics {
	return destination.Metrics{}
}

func (*Client) Read(context.Context, *arrow.Schema, string, chan<- arrow.Record) error {
	return nil
}

func (*Client) Migrate(context.Context, schema.Schemas) error {
	return nil
}

//revive:disable We need to range over the channel to clear it, but revive thinks it can be removed
func (c *Client) Write(_ context.Context, _ schema.Schemas, res <-chan arrow.Record) error {
	if c.spec.ErrorOnWrite {
		return errors.New("error_on_write is true")
	}
	for range res {
		// do nothing
	}
	return nil
}

func (c *Client) WriteTableBatch(context.Context, *arrow.Schema, []arrow.Record) error {
	if c.spec.ErrorOnWrite {
		return errors.New("error_on_write is true")
	}
	return nil
}

func (*Client) Close(context.Context) error {
	return nil
}

func (*Client) DeleteStale(context.Context, schema.Schemas, string, time.Time) error {
	return nil
}
