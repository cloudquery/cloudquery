package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client/services"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/slack-go/slack"
)

const (
	defaultMaxRetries = 3
	defaultBackoff    = 1 * time.Second
)

type Client struct {
	logger zerolog.Logger
	spec   specs.Source
	Slack  services.SlackClient
	Teams  []slack.Team
	TeamID string
	BotID  string

	maxRetries int
	backoff    time.Duration // backoff duration between retries (jitter will be added)
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
		Teams:      c.Teams,
		spec:       c.spec,
		Slack:      c.Slack,
		TeamID:     teamID,
		maxRetries: c.maxRetries,
		backoff:    c.backoff,
	}
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
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
	teams, err := listTeams(ctx, client)
	if err != nil {
		return nil, err
	}
	logger.Debug().Int("num_teams", len(teams)).Msg("got teams")
	return &Client{
		logger:     logger,
		spec:       s,
		Slack:      client,
		Teams:      teams,
		maxRetries: defaultMaxRetries,
		backoff:    defaultBackoff,
	}, nil
}

func listTeams(ctx context.Context, client *slack.Client) ([]slack.Team, error) {
	params := slack.ListTeamsParameters{
		Limit: 1000,
	}
	var allTeams []slack.Team
	for {
		teams, cursor, err := client.ListTeamsContext(ctx, params)
		if err != nil {
			return nil, fmt.Errorf("failed to list teams: %w", err)
		}
		allTeams = append(allTeams, teams...)
		if cursor == "" {
			break
		}
		params.Cursor = cursor
	}
	return allTeams, nil
}
