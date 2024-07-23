package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/reverser/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	plugin.UnimplementedDestination

	logger zerolog.Logger
	spec   *spec.Spec
}

func New(_ context.Context, logger zerolog.Logger, s []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "reverser").Logger(),
	}
	if opts.NoConnection {
		return c, nil
	}

	if err := json.Unmarshal(s, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file spec: %w", err)
	}
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.spec.SetDefaults()
	return c, nil
}

func (c *Client) Transform(ctx context.Context, recvRecords <-chan arrow.Record, sendRecords chan<- arrow.Record) error {
	for {
		select {
		case record, ok := <-recvRecords:
			if !ok {
				return nil
			}
			reversedRecord, err := c.reverseStrings(record)
			if err != nil {
				return err
			}
			sendRecords <- reversedRecord
		case <-ctx.Done():
			return nil
		}
	}
}

func (*Client) reverseStrings(record arrow.Record) (arrow.Record, error) {
	for i, column := range record.Columns() {
		if column.DataType().ID() != arrow.STRING {
			continue
		}
		newColumnData := []string{}
		for i := 0; i < column.Len(); i++ {
			if !column.IsValid(i) {
				continue
			}
			s := column.ValueStr(i)
			runes := []rune(s)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			newColumnData = append(newColumnData, string(runes))
		}
		mem := memory.NewGoAllocator()
		bld := array.NewStringBuilder(mem)

		// create an array with 4 values, no null
		bld.AppendValues(newColumnData, nil)
		var err error
		record, err = record.SetColumn(i, bld.NewStringArray())
		if err != nil {
			return nil, err
		}
	}
	return record, nil
}

func (c *Client) Close(ctx context.Context) error {
	return nil
}
