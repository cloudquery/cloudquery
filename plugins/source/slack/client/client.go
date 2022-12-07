package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client/services"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/slack-go/slack"
)

type Client struct {
	logger     zerolog.Logger
	spec       specs.Source
	Slack      services.SlackClient
	AllTeamIDs []string
	TeamID     string
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.spec.Name
}

func (c *Client) withTeamID(teamID string) schema.ClientMeta {
	return &Client{
		logger:     c.logger.With().Str("team_id", teamID).Logger(),
		AllTeamIDs: c.AllTeamIDs,
		spec:       c.spec,
		Slack:      c.Slack,
		TeamID:     teamID,
	}
}

func Configure(_ context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	var config Spec
	err := s.UnmarshalSpec(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	var opts []slack.Option
	if config.Debug {
		opts = append(opts, slack.OptionDebug(true))
	}
	client := slack.New(config.Token, opts...)
	client.Teams()
	return &Client{
		logger: logger,
		spec:   s,
		Slack:  client,
	}, nil
}
