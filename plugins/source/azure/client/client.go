package client

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	// this is set by table client multiplexer
	SubscriptionID string

	services map[string]*Services // key = subscription ID
	logger   zerolog.Logger
}

func (c *Client) ID() string {
	return c.SubscriptionID
}

func (c *Client) ScopeSubscription() string {
	return "/subscriptions/" + c.SubscriptionID
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

// withSubscription allows multiplexer to create a new client with given SubscriptionID
func (c *Client) withSubscription(subscriptionID string) *Client {
	return &Client{
		services:       c.services,
		logger:         c.logger.With().Str("subscription_id", subscriptionID).Logger(),
		SubscriptionID: subscriptionID,
	}
}

func (c *Client) Services() *Services {
	return c.services[c.SubscriptionID]
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	var spec Spec
	if err := s.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Azure spec: %w", err)
	}

	// New-SDK auth; chained credentials with ordering:
	//  1. Environment
	//  2. ManagedIdentity
	//  3. AzureCLI
	credentials, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	subscriptions, err := getSubscriptions(ctx, &logger, &spec, credentials)
	if err != nil {
		return nil, err
	}

	servicesMap := make(map[string]*Services)
	for _, subscriptionId := range subscriptions {
		subscriptionServices, err := initServices(subscriptionId, credentials, nil) // skip options for now
		if err != nil {
			return nil, err
		}
		servicesMap[subscriptionId] = subscriptionServices
	}
	return &Client{
		logger:   logger,
		services: servicesMap,
	}, nil
}
