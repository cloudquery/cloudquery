package client

import (
	"context"
	"errors"
	"fmt"

	// Import all autorest modules
	"github.com/Azure/azure-sdk-for-go/profiles/latest/subscription/mgmt/subscription"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	subscriptions []string
	logger        zerolog.Logger
	// this is set by table client multiplexer
	SubscriptionId string
	services       map[string]*services.Services
}

func getAuth(logger *zerolog.Logger) (autorest.Authorizer, *azidentity.DefaultAzureCredential, error) {
	// Old-SDK auth with ordering:
	//  1. Environment
	//  2. AzureCLI
	logger.Info().Msg("Trying to authenticate via environment variables")
	azureAuth, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		logger.Info().Msg("Trying to authenticate via CLI")
		azureAuth, err = auth.NewAuthorizerFromCLI()
		if err != nil {
			return nil, nil, err
		}
	}

	// New-SDK auth; chained credentials with ordering:
	//  1. Environment
	//  2. ManagedIdentity
	//  3. AzureCLI
	azCred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return azureAuth, nil, err
	}

	return azureAuth, azCred, nil
}

func getSubscriptions(logger *zerolog.Logger, spec *Spec, azureAuth autorest.Authorizer) ([]string, error) {
	subscriptions := make([]string, 0)
	if len(spec.Subscriptions) == 0 {
		ctx := context.Background()
		svc := subscription.NewSubscriptionsClient()
		svc.Authorizer = azureAuth
		res, err := svc.List(ctx)
		if err != nil {
			return nil, err
		}
		for res.NotDone() {
			for _, sub := range res.Values() {
				switch sub.State {
				case subscription.Disabled:
					logger.Info().Msgf("Not fetching from subscription because it is disabled %s - %s", "subscription", *sub.SubscriptionID)
				case subscription.Deleted:
					logger.Info().Msgf("Not fetching from subscription because it is deleted %s - %s", "subscription", *sub.SubscriptionID)
				default:
					subscriptions = append(subscriptions, *sub.SubscriptionID)
				}
			}
			err := res.NextWithContext(ctx)
			if err != nil {
				return nil, err
			}
		}
		logger.Info().Msgf("No subscriptions specified, going to using all available ones %s %s", "subscriptions", subscriptions)
	} else {
		subscriptions = spec.Subscriptions
	}

	if len(subscriptions) == 0 {
		return nil, errors.New("could not find any subscription")
	}

	return subscriptions, nil
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	var spec Spec
	if err := s.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
	}

	azureAuth, azCred, err := getAuth(&logger)
	if err != nil {
		return nil, err
	}

	subscriptions, err := getSubscriptions(&logger, &spec, azureAuth)
	if err != nil {
		return nil, err
	}

	servicesMap := make(map[string]*services.Services)
	for _, subscriptionId := range subscriptions {
		subscriptionServices, err := services.InitServices(subscriptionId, azureAuth, azCred)
		if err != nil {
			return nil, err
		}
		servicesMap[subscriptionId] = &subscriptionServices
	}
	return &Client{
		logger:        logger,
		subscriptions: subscriptions,
		services:      servicesMap,
	}, nil
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.SubscriptionId
}

// withSubscription allows multiplexer to create a new client with given subscriptionId
func (c Client) withSubscription(subscriptionId string) *Client {
	return &Client{
		subscriptions:  c.subscriptions,
		services:       c.services,
		logger:         c.logger.With().Str("subscription_id", subscriptionId).Logger(),
		SubscriptionId: subscriptionId,
	}
}

func (c Client) Services() *services.Services {
	return c.services[c.SubscriptionId]
}
