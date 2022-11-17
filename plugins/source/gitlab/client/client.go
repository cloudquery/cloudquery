package client

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/xanzy/go-gitlab"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger zerolog.Logger
	Gitlab *gitlab.Client
	spec   specs.Source
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

	gitlabToken, ok := os.LookupEnv("GITLAB_API_TOKEN")
	if !ok {
		if gitlabSpec.Token == "" {
			return nil, errors.New("missing GITLAB_API_TOKEN, either set it as an environment variable or pass it in the configuration")
		}

		gitlabToken = gitlabSpec.Token
	}

	c, err := gitlab.NewClient(gitlabToken)
	if err != nil {
		return nil, err
	}

	return &Client{
		logger: logger,
		Gitlab: c,
		spec:   s,
	}, nil
}
