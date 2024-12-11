package client

import (
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/client/services"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/hermanschaaf/hackernews"
	"github.com/rs/zerolog"
)

const (
	defaultMaxRetries = 5
	defaultBackoff    = 10 * time.Second
)

type Client struct {
	logger     zerolog.Logger
	Spec       Spec
	HackerNews services.HackernewsClient
	Backend    state.Client
	maxRetries int
	backoff    time.Duration // backoff duration between retries (jitter will be added)
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (*Client) ID() string {
	return "hackernews"
}

func New(logger zerolog.Logger, spec Spec, hnClient *hackernews.Client, backend state.Client) (schema.ClientMeta, error) {
	return &Client{
		logger:     logger,
		Spec:       spec,
		HackerNews: hnClient,
		Backend:    backend,
		maxRetries: defaultMaxRetries,
		backoff:    defaultBackoff,
	}, nil
}
