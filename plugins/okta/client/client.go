package client

import (
	"context"
	"fmt"
	"os"

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

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)
	oktaToken, ok := os.LookupEnv("OKTA_API_TOKEN")
	if !ok {
		if providerConfig.Token != "" {
			oktaToken = providerConfig.Token
		} else {
			return nil, fmt.Errorf("missing OKTA_API_TOKEN, either set it as an environment variable or pass it in the configuration")
		}
	}

	if providerConfig.Domain == "" || providerConfig.Domain == exampleDomain {
		return nil, fmt.Errorf(`failed to configure provider, please set your okta "domain" in config.hcl`)
	}

	_, c, err := okta.NewClient(context.Background(), okta.WithOrgUrl(providerConfig.Domain), okta.WithToken(oktaToken), okta.WithCache(true))
	if err != nil {
		return nil, err
	}
	client := Client{
		logger: logger,
		Okta:   c,
	}
	// Return the initialized client and it will be passed to your resources
	return &client, nil
}
