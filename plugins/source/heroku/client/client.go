package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/rs/zerolog"
)

// Client is what you need to create and initialize in Configure.
// It will be passed for each resource fetcher.
type Client struct {
	logger zerolog.Logger
	Heroku *heroku.Service
	spec   specs.Source
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.spec.Name
}

func Configure(ctx context.Context, l zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
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
		logger: l,
		Heroku: h,
		spec:   s,
	}, nil
}
