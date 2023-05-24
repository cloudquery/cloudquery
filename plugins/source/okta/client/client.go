package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
	"github.com/rs/zerolog"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger zerolog.Logger
	spec   specs.Source

	*okta.APIClient
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.spec.Name
}

func New(logger zerolog.Logger, s specs.Source, okt *okta.APIClient) *Client {
	return &Client{
		APIClient: okt,

		logger: logger,
		spec:   s,
	}
}

func Configure(_ context.Context, logger zerolog.Logger, srcSpec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	spec := &Spec{}
	if err := srcSpec.UnmarshalSpec(spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal okta spec: %w", err)
	}

	l := logger.With().Str("module", "okta-source").Logger()
	spec.setDefaults(&l)
	if err := spec.validate(); err != nil {
		return nil, err
	}

	cf := okta.NewConfiguration(
		okta.WithOrgUrl(spec.Domain),
		okta.WithToken(spec.Token),
		okta.WithCache(true),
		okta.WithRateLimitMaxBackOff(int64(spec.RateLimit.MaxBackoff/time.Second)), // this param takes int64 of seconds
		okta.WithRateLimitMaxRetries(spec.RateLimit.MaxRetries),
	)
	cf.Debug = spec.Debug
	c := okta.NewAPIClient(cf)

	return New(l, srcSpec, c), nil
}
