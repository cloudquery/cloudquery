package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/test/pk"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.DefaultReverseTransformer
	logger zerolog.Logger
	spec   Spec
	pks    pk.Store
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	var testConfig Spec
	err := spec.UnmarshalSpec(&testConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	return &Client{
		logger: logger.With().Str("module", "test").Logger(),
		spec:   testConfig,
		pks:    pk.NewStore(),
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

func (c *Client) checkPKViolation(table *schema.Table, res []any) {
	if c.pks.IsDuplicate(table, res) {
		c.logger.Error().
			Str("table", table.Name).
			Str("columns", "("+pk.Columns(table)+")").
			Str("value", "("+pk.Convert(table, res)+")").
			Msg("Duplicate PK")
	}
}

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *destination.ClientResource) error {
	if c.spec.ErrorOnWrite {
		return errors.New("error_on_write is true")
	}
	for r := range res {
		// check for PK issues
		table := tables.Get(r.TableName)

		c.checkPKViolation(table, r.Data)
	}
	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, res [][]any) error {
	if c.spec.ErrorOnWrite {
		return errors.New("error_on_write is true")
	}

	for _, r := range res {
		c.checkPKViolation(table, r)
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
