package client

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client/services"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/xanzy/go-gitlab"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger  zerolog.Logger
	Gitlab  Services
	spec    specs.Source
	BaseURL string
}
type Services struct {
	Users    services.UsersClient
	Groups   services.GroupsClient
	Projects services.ProjectsClient
	Settings services.SettingsClient
	Releases services.ReleasesClient
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.spec.Name
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	gitlabSpec := &Spec{}
	if err := s.UnmarshalSpec(gitlabSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gitlab spec: %w", err)
	}

	gitlabToken := gitlabSpec.Token
	if gitlabToken == "" {
		gitlabToken, _ = os.LookupEnv("GITLAB_API_TOKEN")
	}
	if gitlabToken == "" {
		return nil, errors.New("missing GITLAB_API_TOKEN, either set it as an environment variable or pass it in the spec")
	}

	opts := []gitlab.ClientOptionFunc{}
	if gitlabSpec.BaseURL != "" {
		opts = append(opts, gitlab.WithBaseURL(gitlabSpec.BaseURL))
	}

	c, err := gitlab.NewClient(gitlabToken, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		logger: logger,
		Gitlab: NewServices(c),
		spec:   s,
	}, nil
}

// Take gitlab.Client as an argument and return an initialized Services struct
func NewServices(c *gitlab.Client) Services {
	return Services{
		Users:    c.Users,
		Groups:   c.Groups,
		Projects: c.Projects,
		Settings: c.Settings,
		Releases: c.Releases,
	}
}
