package client

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
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

func New(logger zerolog.Logger, spec Spec) (schema.ClientMeta, error) {
	return &Client{
		Logger:      logger,
		OfferCodes:  spec.OfferCodes,
		RegionCodes: spec.RegionCodes,
	}, nil
}
