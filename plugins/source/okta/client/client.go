package client

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/rs/zerolog"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger zerolog.Logger
	Okta   *okta.Client
	spec   specs.Source
}

const exampleDomain = "https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com"

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.spec.Name
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	oktaSpec := &Spec{}
	if err := s.UnmarshalSpec(oktaSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal okta spec: %w", err)
	}

	oktaToken, ok := os.LookupEnv("OKTA_API_TOKEN")
	if !ok {
		if oktaSpec.Token == "" {
			return nil, errors.New("missing OKTA_API_TOKEN, either set it as an environment variable or pass it in the configuration")
		}

		oktaToken = oktaSpec.Token
	}

	if oktaSpec.Domain == "" || oktaSpec.Domain == exampleDomain {
		return nil, errors.New(`failed to configure provider, please set your okta "domain" in okta.yml`)
	}

	_, c, err := okta.NewClient(context.Background(), okta.WithOrgUrl(oktaSpec.Domain), okta.WithToken(oktaToken), okta.WithCache(true))
	if err != nil {
		return nil, err
	}

	return &Client{
		logger: logger,
		Okta:   c,
		spec:   s,
	}, nil
}
