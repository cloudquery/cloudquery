package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
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

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var facebookSpec Spec

	if err := s.UnmarshalSpec(&facebookSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	if err := facebookSpec.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate plugin spec: %w", err)
	}

	facebookClient := rest.NewFacebookClient( /* httpClient */ nil, facebookSpec.AdAccountId, facebookSpec.AccessToken)

	return &Client{
		FacebookClient: facebookClient,
		Spec:           &facebookSpec,
		Logger:         logger,
	}, nil
}
