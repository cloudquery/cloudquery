package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
	"github.com/rs/zerolog"
)

const (
	defaultMaxRetries = 10
	defaultBackoff    = 65 * time.Second
)

type Client struct {
	*snyk.Client

	Spec           Spec
	OrganizationID string
	Organizations  []snyk.Organization

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
	m := l.logger.Debug()
	if len(args)%2 != 0 {
		for i := 0; i < len(args); i++ {
			m = m.Interface(fmt.Sprintf("arg_%02d", i), args[i])
		}
		m.Msg("Log from Snyk SDK")
		return
	}
	for i := 0; i < len(args); i += 2 {
		k, ok := args[i].(string)
		if !ok {
			m = m.Interface(fmt.Sprintf("arg_%02d", i), args[i])
			m = m.Interface(fmt.Sprintf("arg_%02d", i+1), args[i+1])
			continue
		}
		m = m.Interface(k, args[i+1])
	}
	m.Msg("Log from Snyk SDK")
}

func New(ctx context.Context, logger zerolog.Logger, spec Spec) (schema.ClientMeta, error) {
	httpClient := http.DefaultClient
	httpClient.Timeout = 1 * time.Minute
	options := []snyk.ClientOption{
		snyk.WithHTTPClient(httpClient),
		snyk.WithUserAgent("cloudquery:source-snyk"),
		snyk.WithLogger(&SnykLogger{logger: logger}),
		snyk.WithLogRequests(true), // these will be filtered out by the logger if not in debug mode
	}
	if len(spec.EndpointURL) > 0 {
		options = append(options, snyk.WithBaseURL(spec.EndpointURL))
	}

	client := snyk.NewClient(spec.APIKey, options...)

	maxRetries := defaultMaxRetries
	if spec.Retries > 0 {
		maxRetries = spec.Retries
	}
	backoff := defaultBackoff
	if spec.RetryDelaySeconds > 0 {
		backoff = time.Duration(spec.RetryDelaySeconds) * time.Second
	}
	c := &Client{
		Client:     client,
		Spec:       spec,
		logger:     logger,
		maxRetries: maxRetries,
		backoff:    backoff,
	}

	return c, c.initOrganizations(ctx)
}
