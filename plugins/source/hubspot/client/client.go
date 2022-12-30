package client

import (
	"context"
	"fmt"

	"github.com/clarkmcc/go-hubspot"
	"github.com/clarkmcc/go-hubspot/generated/v3/companies"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	authorizer hubspot.Authorizer
	id         string
	host       string

	Companies *companies.APIClient

	logger zerolog.Logger
}

var _ schema.ClientMeta = (*Client)(nil)

func (c *Client) ID() string {
	if len(c.host) > 0 {
		return c.id + "(" + c.host + ")"
	}
	return c.id
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) AuthContext(ctx context.Context) context.Context {
	return hubspot.WithAuthorizer(ctx, c.authorizer)
}

func Configure(_ context.Context, logger zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
	hsSpec := new(Spec)
	err := spec.UnmarshalSpec(hsSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	if len(hsSpec.Token) == 0 {
		return nil, fmt.Errorf("missing token in configuration")
	}

	client := &Client{
		authorizer: hubspot.NewTokenAuthorizer(hsSpec.Token),
		host:       hsSpec.Host,
		id:         "",
		logger:     logger,
	}

	// companies
	{
		cfg := companies.NewConfiguration()
		if len(hsSpec.Host) > 0 {
			cfg.Host = hsSpec.Host
		}
		client.Companies = companies.NewAPIClient(cfg)
	}

	return client, nil
}
