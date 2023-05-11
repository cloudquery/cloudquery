package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger      zerolog.Logger
	Endpoint    string
	OfferCodes  []string
	RegionCodes []string
}

func (*Client) ID() string {
	return "awspricing"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}
	pluginSpec.SetDefaults()

	return &Client{
		Logger:      logger,
		OfferCodes:  pluginSpec.OfferCodes,
		RegionCodes: pluginSpec.RegionCodes,
	}, nil
}
