package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/slack-go/slack"
)

type Client struct {
	logger zerolog.Logger
	spec   specs.Source
	Slack  *slack.Client
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.spec.Name
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
	var config Spec
	err := spec.UnmarshalSpec(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	var opts []slack.Option
	if config.Debug {
		opts = append(opts, slack.OptionDebug(true))
	}
	client := slack.New(config.Token, opts...)
	return &Client{
		logger: logger,
		spec:   spec,
		Slack:  client,
	}, nil
}
