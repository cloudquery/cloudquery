package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
	"github.com/rs/zerolog"
)

type Client struct {
	*snyk.Client

	OrganizationID string
	organizations  []string

	logger zerolog.Logger
}

var _ schema.ClientMeta = (*Client)(nil)

func (c *Client) ID() string {
	return c.OrganizationID
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	snykSpec := new(Spec)
	err := spec.UnmarshalSpec(snykSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}

	client, err := snykSpec.getClient(spec.Version)
	if err != nil {
		return nil, fmt.Errorf("failed to create Snyk client: %w", err)
	}

	c := &Client{
		Client:        client,
		logger:        logger,
		organizations: snykSpec.Organizations,
	}

	return c, c.initOrganizations(ctx)
}
