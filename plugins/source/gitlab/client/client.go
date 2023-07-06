package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/xanzy/go-gitlab"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger         zerolog.Logger
	Gitlab         *gitlab.Client
	spec           Spec
	BaseURL        string
	MinAccessLevel *gitlab.AccessLevelValue
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (*Client) ID() string {
	return "gitlab"
}

func Configure(ctx context.Context, logger zerolog.Logger, s Spec) (schema.ClientMeta, error) {
	if err := s.Validate(); err != nil {
		return nil, err
	}

	var minAccessLevel *gitlab.AccessLevelValue
	opts := []gitlab.ClientOptionFunc{}
	if s.BaseURL != "" {
		opts = append(opts, gitlab.WithBaseURL(s.BaseURL))
	} else {
		// on GitLab SaaS we don't want to sync the whole of GitLab, so we sync based on access level
		// TODO: use gitlab.MinimalAccessPermissions once supported. Related https://gitlab.com/gitlab-org/gitlab/-/issues/296089#note_496386453
		minAccessLevel = gitlab.AccessLevel(gitlab.GuestPermissions)
	}

	c, err := gitlab.NewClient(s.Token, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		logger:         logger,
		Gitlab:         c,
		spec:           s,
		MinAccessLevel: minAccessLevel,
	}, nil
}
