package client

import (
	"errors"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	heroku "github.com/heroku/heroku-go/v5"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger hclog.Logger
	Heroku HerokuService
	Team   string
	Teams  []string
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, diag.Diagnostics) {
	providerConfig := config.(*Config)
	// validate provider config
	if providerConfig.Token == "" {
		return nil, diag.FromError(errors.New("missing access token in configuration"), diag.ACCESS)
	}

	// TODO: support teams filter
	// TODO: validate provider config

	heroku.DefaultTransport.BearerToken = providerConfig.Token
	client := heroku.DefaultClient
	client.Transport = paginator{transport: client.Transport}
	h := heroku.NewService(client)
	return &Client{
		logger: logger,
		Heroku: h,
		Teams:  providerConfig.Teams,
	}, nil
}
