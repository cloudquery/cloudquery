package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/cloudquery/plugin-sdk/v4/writers/mixedbatchwriter"
	"github.com/cloudquery/plugin-sdk/v4/writers/streamingbatchwriter"

	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	spec   Spec
	writer genericBatchWriter

	plugin.UnimplementedSource
}

var (
	_ plugin.Client = (*Client)(nil)

	ErrOnWrite   = errors.New("error_on_write is true")
	ErrOnMigrate = errors.New("error_on_migrate is true")
	ErrOnInsert  = errors.New("error_on_insert is true")
)

func New(_ context.Context, logger zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var spec Spec
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("invalid spec: %w", err)
	}

	c := &Client{
		logger: logger.With().Str("module", "test").Logger(),
		spec:   spec,
	}

	var err error
	switch {
	case spec.BatchWriter:
		c.writer, err = batchwriter.New(c,
			batchwriter.WithBatchSize(*c.spec.BatchSize),
			batchwriter.WithBatchSizeBytes(*c.spec.BatchSizeBytes),
			batchwriter.WithBatchTimeout(c.spec.BatchTimeout.Duration()),
			batchwriter.WithLogger(c.logger),
		)
	case spec.MixedBatchWriter:
		var mbw *mixedbatchwriter.MixedBatchWriter
		mbw, err = mixedbatchwriter.New(c,
			mixedbatchwriter.WithBatchSize(*c.spec.BatchSize),
			mixedbatchwriter.WithBatchSizeBytes(*c.spec.BatchSizeBytes),
			mixedbatchwriter.WithBatchTimeout(c.spec.BatchTimeout.Duration()),
			mixedbatchwriter.WithLogger(c.logger),
		)
		c.writer = adaptMixedBatchWriter(mbw)
	case spec.StreamBatchWriter:
		c.writer, err = streamingbatchwriter.New(newClientForStreamingBatchWriter(c),
			streamingbatchwriter.WithBatchSizeRows(*c.spec.BatchSize),
			streamingbatchwriter.WithBatchSizeBytes(*c.spec.BatchSizeBytes),
			streamingbatchwriter.WithBatchTimeout(c.spec.BatchTimeout.Duration()),
			streamingbatchwriter.WithLogger(c.logger),
		)
	}
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (*Client) Read(context.Context, *schema.Table, chan<- arrow.Record) error {
	return nil
}

//revive:disable We need to range over the channel to clear it, but revive thinks it can be removed
func (c *Client) Write(ctx context.Context, messages <-chan message.WriteMessage) error {
	if c.spec.ErrorOnWrite {
		return ErrOnWrite
	}

	if c.spec.ExitOnWrite {
		os.Exit(1)
	}

	if c.spec.BatchWriter || c.spec.MixedBatchWriter || c.spec.StreamBatchWriter {
		if err := c.writer.Write(ctx, messages); err != nil {
			return fmt.Errorf("failed to write: %w", err)
		}

		if err := c.writer.Flush(ctx); err != nil {
			return fmt.Errorf("failed to flush messages: %w", err)
		}
		return nil
	}

	for m := range messages {
		if c.spec.ErrorOnMigrate {
			if _, ok := m.(*message.WriteMigrateTable); ok {
				return ErrOnMigrate
			}
		}
		if c.spec.ExitOnMigrate {
			if _, ok := m.(*message.WriteMigrateTable); ok {
				os.Exit(1)
			}
		}

		if m, ok := m.(*message.WriteInsert); ok {
			m.Record.Release()
			if c.spec.ErrorOnInsert {
				return ErrOnInsert
			}
			if c.spec.ExitOnInsert {
				os.Exit(1)
			}
		}
	}
	return nil
}

func (c *Client) Close(ctx context.Context) error {
	if c.writer != nil {
		return c.writer.Close(ctx)
	}
	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, name string, msgs message.WriteInserts) error {
	if c.spec.ErrorOnInsert {
		return ErrOnInsert
	}
	if c.spec.ExitOnInsert {
		os.Exit(1)
	}
	return nil
}

func (c *Client) MigrateTables(ctx context.Context, msgs message.WriteMigrateTables) error {
	if c.spec.ErrorOnMigrate {
		return ErrOnMigrate
	}
	if c.spec.ExitOnMigrate {
		os.Exit(1)
	}
	return nil
}

func (c *Client) DeleteStale(ctx context.Context, msgs message.WriteDeleteStales) error {
	return nil
}

func (c *Client) DeleteRecord(ctx context.Context, msgs message.WriteDeleteRecords) error {
	return nil
}

func (c *Client) DeleteRecordsBatch(ctx context.Context, messages message.WriteDeleteRecords) error {
	return nil
}

func (c *Client) MigrateTableBatch(ctx context.Context, messages message.WriteMigrateTables) error {
	if c.spec.ErrorOnMigrate {
		return ErrOnMigrate
	}
	if c.spec.ExitOnMigrate {
		os.Exit(1)
	}
	return nil
}

func (c *Client) InsertBatch(ctx context.Context, messages message.WriteInserts) error {
	if c.spec.ErrorOnInsert {
		return ErrOnInsert
	}
	if c.spec.ExitOnInsert {
		os.Exit(1)
	}
	return nil
}

func (c *Client) DeleteStaleBatch(ctx context.Context, messages message.WriteDeleteStales) error {
	return nil
}

type ClientForStreamingBatchWriter struct {
	*Client
}

func newClientForStreamingBatchWriter(c *Client) streamingbatchwriter.Client {
	return &ClientForStreamingBatchWriter{Client: c}
}

func (c *ClientForStreamingBatchWriter) DeleteRecords(ctx context.Context, msgs <-chan *message.WriteDeleteRecord) error {
	return nil
}
func (c *ClientForStreamingBatchWriter) MigrateTable(context.Context, <-chan *message.WriteMigrateTable) error {
	if c.spec.ErrorOnMigrate {
		return ErrOnMigrate
	}
	if c.spec.ExitOnMigrate {
		os.Exit(1)
	}
	return nil
}

func (c *ClientForStreamingBatchWriter) DeleteStale(context.Context, <-chan *message.WriteDeleteStale) error {
	return nil
}

func (c *ClientForStreamingBatchWriter) WriteTable(context.Context, <-chan *message.WriteInsert) error {
	if c.spec.ErrorOnInsert {
		return ErrOnInsert
	}
	if c.spec.ExitOnInsert {
		os.Exit(1)
	}
	return nil
}

func TestConnection(ctx context.Context, _ zerolog.Logger, specBytes []byte) error {
	var s Spec
	if err := json.Unmarshal(specBytes, &s); err != nil {
		return &plugin.TestConnError{
			Code:    "INVALID_SPEC",
			Message: fmt.Errorf("failed to unmarshal spec: %w", err),
		}
	}
	s.SetDefaults()

	return nil
}

// genericBatchWriter is a generic interface for batch writers. There are currently 3 implementations, but the mixedbatchwriter
// is different so we need to adapt it to the interface.
type genericBatchWriter interface {
	Write(ctx context.Context, messages <-chan message.WriteMessage) error
	Flush(ctx context.Context) error
	Close(ctx context.Context) error
}

// adaptedMixedBatchWriter is a wrapper around a mixedbatchwriter.MixedBatchWriter that implements the genericBatchWriter interface.
type adaptedMixedBatchWriter struct {
	mbw *mixedbatchwriter.MixedBatchWriter
}

func adaptMixedBatchWriter(mbw *mixedbatchwriter.MixedBatchWriter) adaptedMixedBatchWriter {
	return adaptedMixedBatchWriter{
		mbw: mbw,
	}
}

func (a adaptedMixedBatchWriter) Write(ctx context.Context, messages <-chan message.WriteMessage) error {
	return a.mbw.Write(ctx, messages)
}

func (a adaptedMixedBatchWriter) Flush(ctx context.Context) error {
	return nil
}

func (a adaptedMixedBatchWriter) Close(ctx context.Context) error {
	return nil
}
