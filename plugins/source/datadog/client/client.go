package client

import (
	"context"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger zerolog.Logger

	Accounts []Account
	// this is set by the table client multiplexer
	multiplexedAccount Account
	DDServices         DatadogServices
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.multiplexedAccount.Name
}

func (c *Client) BuildContextV1(ctx context.Context) context.Context {
	return context.WithValue(
		ctx,
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: c.multiplexedAccount.APIKey,
			},
			"appKeyAuth": {
				Key: c.multiplexedAccount.AppKey,
			},
		},
	)
}

func (c *Client) BuildContextV2(ctx context.Context) context.Context {
	return context.WithValue(
		ctx,
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: c.multiplexedAccount.APIKey,
			},
			"appKeyAuth": {
				Key: c.multiplexedAccount.AppKey,
			},
		},
	)
}

func (c *Client) withAccount(account Account) schema.ClientMeta {
	return &Client{
		logger:             c.logger.With().Str("id", account.Name).Logger(),
		Accounts:           c.Accounts,
		multiplexedAccount: account,
		DDServices:         c.DDServices,
	}
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	cfSpec := &Spec{}
	if err := s.UnmarshalSpec(cfSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal datadog spec: %w", err)
	}

	if len(cfSpec.Accounts) == 0 {
		return nil, errors.New("no datadog accounts configured")
	}
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)

	client := Client{
		logger:     logger,
		Accounts:   cfSpec.Accounts,
		DDServices: NewDatadogServices(apiClient),
	}

	// Return the initialized client, and it will be passed to your resources
	return &client, nil
}
