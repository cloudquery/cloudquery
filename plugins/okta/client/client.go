package client

import (
	"context"
	"errors"
	"os"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

const exampleDomain = "https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com"

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger hclog.Logger
	Okta   *okta.Client
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, diag.Diagnostics) {
	providerConfig := config.(*Config)
	oktaToken, ok := os.LookupEnv("OKTA_API_TOKEN")
	if !ok {
		if providerConfig.Token == "" {
			return nil, diag.FromError(errors.New("missing OKTA_API_TOKEN, either set it as an environment variable or pass it in the configuration"), diag.USER)
		}

		oktaToken = providerConfig.Token
	}

	if providerConfig.Domain == "" || providerConfig.Domain == exampleDomain {
		return nil, diag.FromError(errors.New(`failed to configure provider, please set your okta "domain" in config.hcl`), diag.USER)
	}

	_, c, err := okta.NewClient(context.Background(), okta.WithOrgUrl(providerConfig.Domain), okta.WithToken(oktaToken), okta.WithCache(true))
	if err != nil {
		return nil, classifyError(err, diag.INTERNAL)
	}
	client := Client{
		logger: logger,
		Okta:   c,
	}
	// Return the initialized client and it will be passed to your resources
	return &client, nil
}
