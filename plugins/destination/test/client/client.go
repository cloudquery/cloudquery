package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v3/schema"
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

func (*Client) Read(context.Context, *schema.Table, string, chan<- arrow.Record) error {
	return nil
}

func (*Client) Migrate(context.Context, schema.Tables) error {
	return nil
}

//revive:disable We need to range over the channel to clear it, but revive thinks it can be removed
func (c *Client) Write(_ context.Context, _ schema.Tables, records <-chan arrow.Record) error {
	if c.spec.ErrorOnWrite {
		return errors.New("error_on_write is true")
	}
	for record := range records {
		record.Release()
	}
	return nil
}

func (c *Client) WriteTableBatch(_ context.Context, _ *schema.Table, records []arrow.Record) error {
	defer func() {
		for _, record := range records {
			record.Release()
		}
	}()
	if c.spec.ErrorOnWrite {
		return errors.New("error_on_write is true")
	}
	return nil
}

func (*Client) Close(context.Context) error {
	return nil
}

func (*Client) DeleteStale(context.Context, schema.Tables, string, time.Time) error {
	return nil
}
