package client

import (
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	spec   Spec

	TeamID   string
	TeamIDs  []string
	Services *vercel.Client
	Backend  state.Client
}

func New(logger zerolog.Logger, spec Spec, services *vercel.Client, teamIDs []string, bk state.Client) *Client {
	return &Client{
		logger:   logger,
		spec:     spec,
		TeamIDs:  teamIDs,
		Services: services,
		Backend:  bk,
	}
}

func (c Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c Client) ID() string {
	n := []string{"vercel"}
	if c.TeamID != "" {
		n = append(n, c.TeamID)
	}
	return strings.Join(n, "_")
}

func (c Client) WithTeamID(teamID string) schema.ClientMeta {
	return &Client{
		logger:   c.logger.With().Str("team_id", teamID).Logger(),
		spec:     c.spec,
		TeamID:   teamID,
		TeamIDs:  c.TeamIDs,
		Services: c.Services.WithTeamID(teamID),
	}
}
