package client

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/cloudquery/httpcache"
	"github.com/cloudquery/httpcache/diskcache"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/gofri/go-github-ratelimit/github_ratelimit"
	"github.com/google/go-github/v59/github"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
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

	Spec Spec
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
		apiCall := ""
		if callbackContext.Request != nil {
			apiCall = callbackContext.Request.URL.String()
		}
		logger.Warn().Msgf("GitHub secondary rate limit detected for API call: %s. Sleeping until %s", apiCall, callbackContext.SleepUntil.Format(time.RFC3339))
	}
}

func New(ctx context.Context, logger zerolog.Logger, spec Spec) (schema.ClientMeta, error) {
	ghServices := map[string]GithubServices{}
	for _, auth := range spec.AppAuth {
		var (
			itr *ghinstallation.Transport
			err error
		)
		appId, err := strconv.ParseInt(auth.AppID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse AppID %v: %w", auth.AppID, err)
		}
		installationId, err := strconv.ParseInt(auth.InstallationID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse InstallationID %v: %w", auth.InstallationID, err)
		}
		transport := http.DefaultTransport
		if spec.LocalCachePath != "" {
			cache := diskcache.New(spec.LocalCachePath)
			transport = httpcache.NewTransport(cache)
		}
		if auth.PrivateKeyPath != "" {
			itr, err = ghinstallation.NewKeyFromFile(transport, appId, installationId, auth.PrivateKeyPath)
		} else {
			itr, err = ghinstallation.New(transport, appId, installationId, []byte(auth.PrivateKey))
		}
		if err != nil {
			return nil, fmt.Errorf("failed to create GitHub client for org %v: %w", auth.Org, err)
		}
		if spec.EnterpriseSettings != nil {
			itr.BaseURL = spec.EnterpriseSettings.BaseURL
		}
		ghc, err := githubClientForHTTPClient(itr, logger, spec.EnterpriseSettings)
		if err != nil {
			return nil, fmt.Errorf("failed to create GitHub client for org %v: %w", auth.Org, err)
		}
		ghServices[auth.Org] = servicesForClient(ghc)
	}

	var defaultServices GithubServices
	if spec.AccessToken != "" {
		var cl *http.Client
		if spec.LocalCachePath != "" {
			cache := diskcache.New(spec.LocalCachePath)
			cl = httpcache.NewTransport(cache).Client()
		}
		httpClient := github.NewClient(cl).WithAuthToken(spec.AccessToken)
		ghc, err := githubClientForHTTPClient(httpClient.Client().Transport, logger, spec.EnterpriseSettings)
		if err != nil {
			return nil, fmt.Errorf("failed to create GitHub client for access token: %w", err)
		}
		defaultServices = servicesForClient(ghc)
		_, _, err = defaultServices.Users.Get(ctx, "")
		if err != nil {
			return nil, fmt.Errorf("failed to authenticate with GitHub using access token: %w", err)
		}

	} else {
		defaultServices = ghServices[spec.AppAuth[0].Org]
	}
	ghServices[""] = defaultServices

	c := &Client{
		logger:      logger,
		orgServices: ghServices,
		orgs:        spec.Orgs,
		repos:       spec.Repos,
		Spec:        spec,
	}
	c.logger.Info().Msg("Discovering repositories")
	orgRepositories, err := c.discoverRepositories(ctx, spec.DiscoveryConcurrency, spec.Orgs, spec.Repos, spec.IncludeArchivedRepos)
	if err != nil {
		return nil, fmt.Errorf("failed to discover repositories: %w", err)
	}
	c.orgRepositories = orgRepositories
	return c, nil
}

func servicesForClient(c *github.Client) GithubServices {
	return GithubServices{
		Actions:         c.Actions,
		Billing:         c.Billing,
		Dependabot:      c.Dependabot,
		Issues:          c.Issues,
		Organizations:   c.Organizations,
		Repositories:    c.Repositories,
		Teams:           c.Teams,
		DependencyGraph: c.DependencyGraph,
		Users:           c.Users,
	}
}

func (c *Client) removeArchivedRepos(repos []*github.Repository) []*github.Repository {
	filtered := []*github.Repository{}
	for _, repo := range repos {
		if repo.GetArchived() {
			c.logger.Debug().Msgf("Skipping archived repository %q", repo.GetFullName())
			continue
		}
		filtered = append(filtered, repo)
	}
	return filtered
}

func (c *Client) discoverRepositories(ctx context.Context, discoveryConcurrency int, orgs []string, repos []string, includeArchivedRepos bool) (map[string][]*github.Repository, error) {
	orgRepos := make(map[string][]*github.Repository)
	orgReposLock := sync.Mutex{}
	errorGroup, gtx := errgroup.WithContext(ctx)
	errorGroup.SetLimit(discoveryConcurrency)

	for _, org := range orgs {
		org := org
		opts := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 100}}
		services := c.servicesForOrg(org)
		errorGroup.Go(func() error {
			orgRepositories := []*github.Repository{}
			for {
				repos, resp, err := services.Repositories.ListByOrg(gtx, org, opts)
				if err != nil {
					return err
				}
				if !includeArchivedRepos {
					repos = c.removeArchivedRepos(repos)
				}
				orgRepositories = append(orgRepositories, repos...)

				if resp.NextPage == 0 {
					break
				}
				opts.Page = resp.NextPage
			}

			orgReposLock.Lock()
			defer orgReposLock.Unlock()
			orgRepos[org] = orgRepositories

			return nil
		})
	}

	if err := errorGroup.Wait(); err != nil {
		return nil, err
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

func githubClientForHTTPClient(httpTransport http.RoundTripper, logger zerolog.Logger, enterpriseSettings *EnterpriseSettings) (*github.Client, error) {
	rateLimiter, err := github_ratelimit.NewRateLimitWaiterClient(httpTransport, github_ratelimit.WithLimitDetectedCallback(limitDetectedCallback(logger)))
	if err != nil {
		return nil, fmt.Errorf("failed to create rate limiter: %w", err)
	}

	if enterpriseSettings != nil {
		return github.NewClient(rateLimiter).WithEnterpriseURLs(enterpriseSettings.BaseURL, enterpriseSettings.UploadURL)
	}

	return github.NewClient(rateLimiter), nil
}

func splitRepo(repo string) []string {
	return strings.Split(repo, "/")
}
