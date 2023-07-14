package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/beatlabs/github-auth/app/inst"
	"github.com/beatlabs/github-auth/key"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/gofri/go-github-ratelimit/github_ratelimit"
	"github.com/google/go-github/v49/github"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger zerolog.Logger

	orgServices map[string]GithubServices
	Github      GithubServices

	Org        string
	Repository *github.Repository

	orgs            []string
	orgRepositories map[string][]*github.Repository
	repos           []string
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

func (c *Client) servicesForOrg(org string) GithubServices {
	if _, ok := c.orgServices[org]; ok {
		return c.orgServices[org]
	}
	return c.orgServices[""]
}

func (c *Client) WithOrg(org string) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("org", org).Logger()
	newC.Org = org
	newC.Github = c.servicesForOrg(org)
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

func New(ctx context.Context, logger zerolog.Logger, spec Spec) (schema.ClientMeta, error) {
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate GitHub spec: %w", err)
	}
	spec.SetDefaults()

	ghServices := map[string]GithubServices{}
	for _, auth := range spec.AppAuth {
		k, err := key.FromFile(auth.PrivateKeyPath)
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key: %w", err)
		}
		i, err := inst.NewConfig(auth.AppID, auth.InstallationID, k)
		if err != nil {
			return nil, fmt.Errorf("failed to create app config: %w", err)
		}
		httpClient := i.Client(ctx)
		ghc, err := githubClientForHTTPClient(httpClient, logger, spec.EnterpriseSettings)
		if err != nil {
			return nil, fmt.Errorf("failed to create GitHub client for org %v: %w", auth.Org, err)
		}
		ghServices[auth.Org] = servicesForClient(ghc)
	}

	var defaultServices GithubServices
	if spec.AccessToken != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: spec.AccessToken})
		httpClient := oauth2.NewClient(ctx, ts)
		ghc, err := githubClientForHTTPClient(httpClient, logger, spec.EnterpriseSettings)
		if err != nil {
			return nil, fmt.Errorf("failed to create GitHub client for access token: %w", err)
		}
		defaultServices = servicesForClient(ghc)
	} else {
		defaultServices = ghServices[spec.AppAuth[0].Org]
	}
	ghServices[""] = defaultServices

	c := &Client{
		logger:      logger,
		orgServices: ghServices,
		orgs:        spec.Orgs,
		repos:       spec.Repos,
	}
	c.logger.Info().Msg("Discovering repositories")
	orgRepositories, err := c.discoverRepositories(ctx, spec.Orgs, spec.Repos)
	if err != nil {
		return nil, fmt.Errorf("failed to discover repositories: %w", err)
	}
	c.orgRepositories = orgRepositories
	return c, nil
}

func servicesForClient(c *github.Client) GithubServices {
	return GithubServices{
		Actions:       c.Actions,
		Billing:       c.Billing,
		Dependabot:    c.Dependabot,
		Issues:        c.Issues,
		Organizations: c.Organizations,
		Repositories:  c.Repositories,
		Teams:         c.Teams,
	}
}

func (c *Client) discoverRepositories(ctx context.Context, orgs []string, repos []string) (map[string][]*github.Repository, error) {
	opts := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 100}}

	orgRepos := make(map[string][]*github.Repository)
	for _, org := range orgs {
		services := c.servicesForOrg(org)
		for {
			repos, resp, err := services.Repositories.ListByOrg(ctx, org, opts)
			if err != nil {
				return nil, err
			}
			orgRepos[org] = append(orgRepos[org], repos...)

			if resp.NextPage == 0 {
				break
			}
			opts.Page = resp.NextPage
		}
	}
	seenOrgs := make(map[string]struct{})
	for _, repo := range repos {
		repoSplit := splitRepo(repo)
		if len(repoSplit) != 2 {
			return nil, fmt.Errorf("invalid repository: %s", repo)
		}
		org, name := repoSplit[0], repoSplit[1]
		services := c.servicesForOrg(org)
		r, _, err := services.Repositories.Get(ctx, org, name)
		if err != nil {
			return nil, err
		}
		if _, seen := seenOrgs[org]; !seen {
			// if org is also in orgs list, we will only sync repos in repos list
			seenOrgs[org] = struct{}{}
			orgRepos[org] = []*github.Repository{}
		}
		orgRepos[org] = append(orgRepos[org], r)
	}

	return orgRepos, nil
}

func githubClientForHTTPClient(httpClient *http.Client, logger zerolog.Logger, enterpriseSettings *EnterpriseSettings) (*github.Client, error) {
	rateLimiter, err := github_ratelimit.NewRateLimitWaiterClient(httpClient.Transport, github_ratelimit.WithLimitDetectedCallback(limitDetectedCallback(logger)))
	if err != nil {
		return nil, fmt.Errorf("failed to create rate limiter: %w", err)
	}

	if enterpriseSettings != nil {
		return github.NewEnterpriseClient(enterpriseSettings.BaseURL, enterpriseSettings.UploadURL, rateLimiter)
	}

	return github.NewClient(rateLimiter), nil
}

func splitRepo(repo string) []string {
	return strings.Split(repo, "/")
}
