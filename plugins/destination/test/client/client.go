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
	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	spec   Spec
	writer *batchwriter.BatchWriter

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

	c := &Client{
		logger: logger.With().Str("module", "test").Logger(),
		spec:   spec,
	}

	if spec.BatchWriter {
		var err error
		c.writer, err = batchwriter.New(c,
			batchwriter.WithBatchSize(*c.spec.BatchSize),
			batchwriter.WithBatchSizeBytes(*c.spec.BatchSizeBytes),
			batchwriter.WithBatchTimeout(c.spec.BatchTimeout.Duration()),
			batchwriter.WithLogger(c.logger),
		)
		if err != nil {
			return nil, err
		}
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

	if c.spec.BatchWriter {
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
