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
	"github.com/thoas/go-funk"
)

type Client struct {
	subscriptions []string
	// This is to cache full objects returned from ListSubscriptions on initialisation
	SubscriptionsObjects []*armsubscription.Subscription
	logger               zerolog.Logger
	registeredNamespaces map[string]map[string]bool
	resourceGroups       map[string][]string
	// this is set by table client multiplexer
	SubscriptionId string
	// this is set by table client multiplexer (SubscriptionResourceGroupMultiplexRegisteredNamespace)
	ResourceGroup string
	Creds         azcore.TokenCredential
	Options       *arm.ClientOptions
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
		// we record all returned values, even disabled
		c.SubscriptionsObjects = append(c.SubscriptionsObjects, page.Value...)
		for _, sub := range page.Value {
			if *sub.State == armsubscription.SubscriptionStateEnabled {
				c.subscriptions = append(c.subscriptions, strings.TrimPrefix(*sub.ID, "/subscriptions/"))
			}
		}
	}

	return nil
}

func getResourceGroupNames(resourceGroups []*armresources.ResourceGroup) []string {
	names := make([]string, len(resourceGroups))
	for i, rg := range resourceGroups {
		names[i] = *rg.Name
	}
	return names
}

func (c *Client) disocverResourceGroups(ctx context.Context) error {
	c.resourceGroups = make(map[string][]string, len(c.subscriptions))
	c.registeredNamespaces = make(map[string]map[string]bool, len(c.subscriptions))

	for _, subID := range c.subscriptions {
		c.registeredNamespaces[subID] = make(map[string]bool)
		cl, err := armresources.NewResourceGroupsClient(subID, c.Creds, nil)
		if err != nil {
			return fmt.Errorf("failed to create resource group client: %w", err)
		}
		pager := cl.NewListPager(&armresources.ResourceGroupsClientListOptions{})
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				return fmt.Errorf("failed to list resource groups: %w", err)
			}
			if len(page.Value) == 0 {
				continue
			}
			c.resourceGroups[subID] = append(c.resourceGroups[subID], getResourceGroupNames(page.Value)...)
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

	uniqueSubscriptions := funk.Uniq(spec.Subscriptions).([]string)
	c := &Client{
		logger:        logger,
		subscriptions: uniqueSubscriptions,
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
	if c.ResourceGroup != "" {
		return fmt.Sprintf("subscriptions/%s/resourceGroups/%s", c.SubscriptionId, c.ResourceGroup)
	}
	return fmt.Sprintf("subscriptions/%s", c.SubscriptionId)
}

// withSubscription allows multiplexer to create a new client with given subscriptionId
func (c *Client) withSubscription(subscriptionId string) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("subscription_id", subscriptionId).Logger()
	newC.SubscriptionId = subscriptionId
	return &newC
}

func (c *Client) withResourceGroup(resourceGroup string) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("resource_group", resourceGroup).Logger()
	newC.ResourceGroup = resourceGroup
	return &newC
}
