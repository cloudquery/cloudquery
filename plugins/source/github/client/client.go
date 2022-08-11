package client

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/google/go-github/v45/github"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/oauth2"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger hclog.Logger

	// CHANGEME:  Usually you store here your 3rd party clients and use them in the fetcher
	Github GithubServices

	Org string

	Orgs []string
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func (c Client) WithOrg(org string) schema.ClientMeta {
	return &Client{
		logger: c.logger.With("org", org),
		Github: c.Github,
		Org:    org,
		Orgs:   c.Orgs,
	}
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, diag.Diagnostics) {
	providerConfig := config.(*Config)
	_ = providerConfig
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: providerConfig.AccessToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)

	// Init your client and 3rd party clients using the user's configuration
	// passed by the SDK providerConfig
	return &Client{
		logger: logger,
		Github: GithubServices{
			Teams:         client.Teams,
			Billing:       client.Billing,
			Repositories:  client.Repositories,
			Organizations: client.Organizations,
			Issues:        client.Issues,
		},
		Orgs: providerConfig.Orgs,
	}, nil
}
