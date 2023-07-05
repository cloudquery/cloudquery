package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

type Client struct {
	FacebookClient *rest.FacebookClient

	Spec *Spec

	Logger zerolog.Logger
}

func (c *Client) ID() string {
	return c.Spec.AdAccountId
}

func New(ctx context.Context, logger zerolog.Logger, s Spec) (schema.ClientMeta, error) {
	if err := s.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate plugin spec: %w", err)
	}

	facebookClient := rest.NewFacebookClient( /* httpClient */ nil, s.AdAccountId, s.AccessToken)

	return &Client{
		FacebookClient: facebookClient,
		Spec:           &s,
		Logger:         logger,
	}, nil
}
