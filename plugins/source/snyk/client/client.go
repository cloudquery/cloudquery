package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/internal/legacy"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
	"github.com/rs/zerolog"
)

const (
	defaultMaxRetries = 5
	defaultBackoff    = 60 * time.Second
)

type Client struct {
	*snyk.Client
	LegacyClient *legacy.Client

	OrganizationID string
	organizations  []string

	logger zerolog.Logger

	maxRetries int
	backoff    time.Duration // backoff duration between retries (jitter will be added)
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

	legacyClient := legacy.NewClient(snykSpec.APIKey)
	c := &Client{
		Client:        client,
		LegacyClient:  legacyClient,
		logger:        logger,
		organizations: snykSpec.Organizations,
		maxRetries:    defaultMaxRetries,
		backoff:       defaultBackoff,
	}

	return c, c.initOrganizations(ctx)
}
