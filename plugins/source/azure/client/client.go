package client

import (
	"context"
	"fmt"
	"strings"

	// Import all autorest modules
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	subscriptions []string
	// This is to cache full objects returned from ListSubscriptions on initialisation
	SubscriptionsObjects []*armsubscription.Subscription
	logger               zerolog.Logger
	registeredNamespaces map[string]map[string]bool
	resourceGroups       map[string][]*armresources.GenericResourceExpanded
	// this is set by table client multiplexer
	SubscriptionId string
	Creds          azcore.TokenCredential
	Options        *arm.ClientOptions
}

func (c *Client) discoverSubscriptions(ctx context.Context) error {
	c.subscriptions = make([]string, 0)
	subscriptionClient, err := armsubscription.NewSubscriptionsClient(c.Creds, nil)
	if err != nil {
		return err
	}
	pager := subscriptionClient.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, sub := range page.Value {
			if *sub.State == armsubscription.SubscriptionStateEnabled {
				c.subscriptions = append(c.subscriptions, strings.TrimPrefix(*sub.ID, "/subscriptions/"))
				c.SubscriptionsObjects = append(c.SubscriptionsObjects, sub)
			}
		}
	}

	return nil
}

func (c *Client) disocverResourceGroups(ctx context.Context) error {
	c.resourceGroups = make(map[string][]*armresources.GenericResourceExpanded, len(c.subscriptions))
	filter := "resourceType eq 'Microsoft.Resources/resourceGroups'"
	c.registeredNamespaces = make(map[string]map[string]bool, len(c.subscriptions))

	for _, subID := range c.subscriptions {
		c.registeredNamespaces[subID] = make(map[string]bool)
		cl, err := armresources.NewClient(subID, c.Creds, nil)
		if err != nil {
			return fmt.Errorf("failed to create resource group client: %w", err)
		}
		pager := cl.NewListPager(&armresources.ClientListOptions{
			Filter: &filter,
		})
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				return fmt.Errorf("failed to list resource groups: %w", err)
			}
			if len(page.Value) == 0 {
				continue
			}
			c.resourceGroups[subID] = append(c.resourceGroups[subID], page.Value...)
		}

		providerClient, err := armresources.NewProvidersClient(subID, c.Creds, nil)
		if err != nil {
			return fmt.Errorf("failed to create provider client: %w", err)
		}
		providerPager := providerClient.NewListPager(nil)
		for providerPager.More() {
			providerPage, err := providerPager.NextPage(ctx)
			if err != nil {
				return fmt.Errorf("failed to list providers: %w", err)
			}
			if len(providerPage.Value) == 0 {
				continue
			}
			for _, p := range providerPage.Value {
				if p.RegistrationState != nil && *p.RegistrationState == "Registered" {
					c.registeredNamespaces[subID][strings.ToLower(*p.Namespace)] = true
				}
			}
		}
	}
	return nil
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	var spec Spec
	var err error
	if err := s.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
	}

	c := &Client{
		logger:        logger,
		subscriptions: spec.Subscriptions,
	}

	c.Creds, err = azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	// if subscription are not specified discover subscriptions with default credentials
	if len(c.subscriptions) == 0 {
		err = c.discoverSubscriptions(ctx)
		if err != nil {
			return nil, err
		}
	}

	if len(c.subscriptions) == 0 {
		return nil, fmt.Errorf("no subscriptions found")
	}

	if err := c.disocverResourceGroups(ctx); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.SubscriptionId
}

// withSubscription allows multiplexer to create a new client with given subscriptionId
func (c *Client) withSubscription(subscriptionId string) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("subscription_id", subscriptionId).Logger()
	newC.SubscriptionId = subscriptionId
	return &newC
}
