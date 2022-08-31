package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/rs/zerolog"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger zerolog.Logger
	Heroku *heroku.Service
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func Configure(ctx context.Context, p *plugins.SourcePlugin, s specs.Source) (schema.ClientMeta, error) {
	var herokuSpec Spec
	if err := s.UnmarshalSpec(&herokuSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal heroku spec: %w", err)
	}

	// validate Heroku spec
	if herokuSpec.Token == "" {
		return nil, errors.New("missing access token in configuration")
	}

	heroku.DefaultTransport.BearerToken = herokuSpec.Token
	client := heroku.DefaultClient
	client.Transport = Paginator{transport: client.Transport}
	h := heroku.NewService(client)
	return &Client{
		Heroku: h,
	}, nil
}
