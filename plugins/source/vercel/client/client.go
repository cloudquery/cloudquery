package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	logger     zerolog.Logger
	sourceSpec specs.Source
	veSpec     Spec

	TeamID   string
	TeamIDs  []string
	Services *vercel.Client
}

func New(logger zerolog.Logger, sourceSpec specs.Source, veSpec Spec, services *vercel.Client, teamIDs []string) Client {
	return Client{
		logger:     logger,
		sourceSpec: sourceSpec,
		veSpec:     veSpec,
		Services:   services,
		TeamIDs:    teamIDs,
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	n := []string{c.sourceSpec.Name}
	if c.TeamID != "" {
		n = append(n, c.TeamID)
	}
	return strings.Join(n, "_")
}

func (c *Client) WithTeamID(teamID string) schema.ClientMeta {
	return &Client{
		logger:     c.logger.With().Str("team_id", teamID).Logger(),
		sourceSpec: c.sourceSpec,
		veSpec:     c.veSpec,
		TeamID:     teamID,
		TeamIDs:    c.TeamIDs,
		Services:   c.Services.WithTeamID(teamID),
	}
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	veSpec := &Spec{}
	if err := s.UnmarshalSpec(veSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal vercel spec: %w", err)
	}

	services, err := getServiceClient(logger.With().Str("source", "stripe-client").Logger(), veSpec, "")
	if err != nil {
		return nil, err
	}
	if len(veSpec.TeamIDs) == 0 {
		veSpec.TeamIDs, err = getTeamIDs(ctx, services)
		if err != nil {
			return nil, fmt.Errorf("failed to discover team ids: %w", err)
		}
	}

	cl := New(logger, s, *veSpec, services, veSpec.TeamIDs)
	return &cl, nil
}

func getServiceClient(logger zerolog.Logger, spec *Spec, teamID string) (*vercel.Client, error) {
	if spec.AccessToken == "" {
		return nil, errors.New("no access token provided")
	}
	if spec.EndpointURL == "" {
		spec.EndpointURL = "https://api.vercel.com"
	}
	if spec.Timeout < 1 {
		spec.Timeout = 5
	}
	if spec.PageSize < 1 {
		spec.PageSize = 100
	}
	if spec.MaxRetries < 1 {
		spec.MaxRetries = 10
	}
	if spec.MaxWait < 1 {
		spec.MaxWait = 300
	}

	return vercel.New(
		logger,
		&http.Client{
			Timeout: time.Duration(spec.Timeout) * time.Second,
		},
		spec.EndpointURL,
		spec.AccessToken,
		teamID,
		spec.MaxRetries,
		spec.MaxWait,
		spec.PageSize,
	), nil
}

func getTeamIDs(ctx context.Context, svc *vercel.Client) ([]string, error) {
	var pg vercel.Paginator
	var teams []string

	for {
		list, p, err := svc.ListTeams(ctx, &pg)
		if err != nil {
			return nil, err
		}
		for _, t := range list {
			teams = append(teams, t.ID)
		}

		if p.Next == nil {
			break
		}
		pg.Next = p.Next
	}

	return teams, nil
}
