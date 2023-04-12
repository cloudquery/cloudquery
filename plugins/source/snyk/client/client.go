package client

import (
	"context"
	"fmt"
	"time"

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

type SnykLogger struct {
	logger zerolog.Logger
}

func (l *SnykLogger) Log(args ...any) {
	if len(args) == 1 {
		l.logger.Debug().Interface("msg", args[0]).Msgf("Log from Snyk SDK")
		return
	}
	if len(args)%2 != 0 {
		l.logger.Debug().Interface("args", args).Msgf("Log from Snyk SDK")
		return
	}
	m := l.logger.Debug()
	for i := 0; i < len(args); i += 2 {
		k, ok := args[i].(string)
		if !ok {

		}
		if i+1 < len(args) {
			m = m.Interface(k, args[i+1])
		}
	}
	m.Msg("Log from Snyk SDK")
}

func Configure(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	snykSpec := new(Spec)
	err := spec.UnmarshalSpec(snykSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	err = snykSpec.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	snykLogger := SnykLogger{
		logger: logger,
	}
	options := []snyk.ClientOption{
		snyk.WithUserAgent("cloudquery/snyk/" + spec.Version),
		snyk.WithLogger(&snykLogger),
		snyk.WithLogRequests(snykSpec.Debug),
	}
	if len(snykSpec.EndpointURL) > 0 {
		options = append(options, snyk.WithBaseURL(snykSpec.EndpointURL))
	}

	client := snyk.NewClient(snykSpec.APIKey, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Snyk client: %w", err)
	}

	c := &Client{
		Client:        client,
		logger:        logger,
		organizations: snykSpec.Organizations,
		maxRetries:    defaultMaxRetries,
		backoff:       defaultBackoff,
	}

	return c, c.initOrganizations(ctx)
}
