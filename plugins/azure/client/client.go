package client

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	_ "github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	subscriptions []string
	services      map[string]*services.Services
	logger        hclog.Logger
	// this is set by table client multiplexer
	SubscriptionId string
}

func NewAzureClient(log hclog.Logger, subscriptionId []string) *Client {
	return &Client{
		logger:        log,
		subscriptions: subscriptionId,
		services:      make(map[string]*services.Services),
	}
}

func (c Client) Logger() hclog.Logger {
	return c.logger
}

// SetSubscriptionServices adds a set of services to an azure subscription
func (c *Client) SetSubscriptionServices(subscriptionId string, services services.Services) {
	c.services[subscriptionId] = &services

}

// Services returns all available services of client, returns subscriptionId's specific services
func (c Client) Services() *services.Services {
	return c.services[c.SubscriptionId]
}

// withSubscription allows multiplexer to create a new client with given subscriptionId
func (c Client) withSubscription(subscriptionId string) *Client {
	return &Client{
		subscriptions:  c.subscriptions,
		services:       c.services,
		logger:         c.logger.With("subscription_id", subscriptionId),
		SubscriptionId: subscriptionId,
	}
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)

	azureAuth, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return nil, err
	}
	client := NewAzureClient(logger, providerConfig.Subscriptions)

	if len(providerConfig.Subscriptions) == 0 {
		ctx := context.Background()
		svc := subscription.NewSubscriptionsClient()
		svc.Authorizer = azureAuth
		res, err := svc.List(ctx)
		if err != nil {
			return nil, err
		}
		subscriptions := make([]string, 0)
		for res.NotDone() {
			for _, sub := range res.Values() {
				subscriptions = append(subscriptions, *sub.SubscriptionID)
			}
			err := res.NextWithContext(ctx)
			if err != nil {
				return nil, err
			}
		}
		client.subscriptions = subscriptions
		logger.Info("No subscriptions specified going to using all available ones", "subscriptions", subscriptions)
	}
	for _, sub := range client.subscriptions {
		client.SetSubscriptionServices(sub, services.InitServices(sub, azureAuth))
	}

	// Return the initialized client and it will be passed to your resources
	return client, nil
}
