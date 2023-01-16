package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/xanzy/go-gitlab"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger  zerolog.Logger
	Gitlab  *gitlab.Client
	spec    specs.Source
	BaseURL string
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.spec.Name
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	gitlabSpec := &Spec{}
	if err := s.UnmarshalSpec(gitlabSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gitlab spec: %w", err)
	}
	if err := gitlabSpec.Validate(); err != nil {
		return nil, err
	}

	opts := []gitlab.ClientOptionFunc{}
	if gitlabSpec.BaseURL != "" {
		opts = append(opts, gitlab.WithBaseURL(gitlabSpec.BaseURL))
	}

	c, err := gitlab.NewClient(gitlabSpec.Token, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		logger: logger,
		Gitlab: c,
		spec:   s,
	}, nil
}
