package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/google/go-github/v48/github"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger zerolog.Logger

	// CHANGEME:  Usually you store here your 3rd party clients and use them in the fetcher
	Github GithubServices

	Org string

	Orgs []string
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.Org
}

func (c Client) WithOrg(org string) schema.ClientMeta {
	return &Client{
		logger: c.logger.With().Str("org", org).Logger(),
		Github: c.Github,
		Org:    org,
		Orgs:   c.Orgs,
	}
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	var spec Spec
	err := s.UnmarshalSpec(&spec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal GitHub spec: %w", err)
	}

	// validate plugin config
	if spec.AccessToken == "" {
		return nil, fmt.Errorf("missing personal access token in configuration")
	}
	if len(spec.Orgs) == 0 {
		return nil, fmt.Errorf("no organizations defined in configuration")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: spec.AccessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := github.NewClient(tc)

	// Init your client and 3rd party clients using the user's configuration
	// passed by the SDK.
	return &Client{
		logger: logger,
		Github: GithubServices{
			Actions:       c.Actions,
			Billing:       c.Billing,
			Dependabot:    c.Dependabot,
			Issues:        c.Issues,
			Organizations: c.Organizations,
			Repositories:  c.Repositories,
			Teams:         c.Teams,
		},
		Orgs: spec.Orgs,
	}, nil
}
