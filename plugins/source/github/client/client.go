package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/gofri/go-github-ratelimit/github_ratelimit"
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

	Org        string
	Repository *github.Repository

	orgs            []string
	orgRepositories map[string][]*github.Repository
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	if c.Repository != nil {
		return fmt.Sprintf("org: %s repo: %s/%s", c.Org, c.Repository.Owner.GetLogin(), c.Repository.GetName())
	}
	return fmt.Sprintf("org:%s", c.Org)
}

func (c *Client) WithOrg(org string) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("org", org).Logger()
	newC.Org = org
	return &newC
}

func (c *Client) WithRepository(repository *github.Repository) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("repository", fmt.Sprintf("%s/%s", repository.Owner.GetLogin(), repository.GetName())).Logger()
	newC.Repository = repository
	return &newC
}

func limitDetectedCallback(logger zerolog.Logger) github_ratelimit.OnLimitDetected {
	return func(callbackContext *github_ratelimit.CallbackContext) {
		logger.Warn().Msgf("GitHub secondary rate limit detected. Sleeping until %s", callbackContext.SleepUntil.Format(time.RFC3339))
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
	rateLimiter, err := github_ratelimit.NewRateLimitWaiterClient(tc.Transport, github_ratelimit.WithLimitDetectedCallback(limitDetectedCallback(logger)))
	if err != nil {
		return nil, fmt.Errorf("failed to create rate limiter: %w", err)
	}
	c := github.NewClient(rateLimiter)

	logger.Info().Msg("Discovering organizations repositories")
	orgRepositories, err := discoverRepositories(ctx, c, spec.Orgs)
	if err != nil {
		return nil, fmt.Errorf("failed to discover organizations repositories: %w", err)
	}

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
		orgs:            spec.Orgs,
		orgRepositories: orgRepositories,
	}, nil
}

func discoverRepositories(ctx context.Context, client *github.Client, orgs []string) (map[string][]*github.Repository, error) {
	opts := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 100}}

	orgRepos := make(map[string][]*github.Repository)
	for _, org := range orgs {
		for {
			repos, resp, err := client.Repositories.ListByOrg(ctx, org, opts)
			if err != nil {
				return nil, err
			}
			orgRepos[org] = append(orgRepos[org], repos...)
			opts.Page = resp.NextPage
			if opts.Page == resp.LastPage {
				break
			}
		}
	}

	return orgRepos, nil
}
