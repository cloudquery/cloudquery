package client

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
  "github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/rs/zerolog"
)

type Client struct {
  Logger zerolog.Logger
}

func (c *Client) ID() string {
  // TODO: Change to either your plugin name or a unique dynamic identifier
	return "ID"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}
  // TODO: Add your client initialization here

  return &Client{
    Logger: logger,
  }, nil
}
