package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	spec   Spec

	plugin.UnimplementedSource
}

var _ plugin.Client = (*Client)(nil)

func New(_ context.Context, logger zerolog.Logger, specBytes []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var spec Spec
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	return &Client{
		logger: logger.With().Str("module", "test").Logger(),
		spec:   spec,
	}, nil
}

func (*Client) Read(context.Context, *schema.Table, chan<- arrow.Record) error {
	return nil
}

//revive:disable We need to range over the channel to clear it, but revive thinks it can be removed
func (c *Client) Write(_ context.Context, messages <-chan message.WriteMessage) error {
	if c.spec.ErrorOnWrite {
		return errors.New("error_on_write is true")
	}
	for m := range messages {
		if m, ok := m.(*message.WriteInsert); ok {
			m.Record.Release()
		}
	}
	return nil
}

func (*Client) Close(context.Context) error {
	return nil
}
