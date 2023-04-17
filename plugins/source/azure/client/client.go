package client

import (
	"context"
	"fmt"
	"strings"
	"sync"

	// Import all autorest modules
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/log"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"github.com/rs/zerolog"
	"github.com/thoas/go-funk"
	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"
)

type Client struct {
	subscriptions []string

	// SubscriptionsObjects is to cache full objects returned from ListSubscriptions on initialisation
	SubscriptionsObjects []*armsubscription.Subscription

	// ResourceGroups is to cache full objects returned from ListResourceGroups on initialisation,
	// as a map from subscription ID to list of resource groups.
	ResourceGroups map[string][]*armresources.ResourceGroup

	logger               zerolog.Logger
	registeredNamespaces map[string]map[string]bool
	// this is set by table client multiplexer
	SubscriptionId string
	// this is set by table client multiplexer (SubscriptionResourceGroupMultiplexRegisteredNamespace)
	ResourceGroup string
	Creds         azcore.TokenCredential
	Options       *arm.ClientOptions

	pluginSpec *Spec
}

func (c *Client) discoverSubscriptions(ctx context.Context) error {
	c.subscriptions = make([]string, 0)
	subscriptionClient, err := armsubscription.NewSubscriptionsClient(c.Creds, c.Options)
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

func (c *Client) getResourceGroupsForSubscription(ctx context.Context, subscriptionId string) ([]*armresources.ResourceGroup, error) {
	cl, err := armresources.NewResourceGroupsClient(subscriptionId, c.Creds, c.Options)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource group client: %w", err)
	}
	var groups []*armresources.ResourceGroup
	pager := cl.NewListPager(&armresources.ResourceGroupsClientListOptions{})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list resource groups: %w", err)
		}
		if len(page.Value) == 0 {
			continue
		}
		groups = append(groups, page.Value...)
	}

	return groups, nil
}

func (c *Client) getRegisteredProvidersForSubscription(ctx context.Context, subscriptionId string) ([]*armresources.Provider, error) {
	providerClient, err := armresources.NewProvidersClient(subscriptionId, c.Creds, c.Options)
	if err != nil {
		return nil, fmt.Errorf("failed to create provider client: %w", err)
	}
	var providers []*armresources.Provider
	providerPager := providerClient.NewListPager(nil)
	for providerPager.More() {
		providerPage, err := providerPager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list providers: %w", err)
		}
		if len(providerPage.Value) == 0 {
			continue
		}
		for _, p := range providerPage.Value {
			if p.RegistrationState != nil && *p.RegistrationState == "Registered" {
				providers = append(providers, p)
			}
		}
	}
	return providers, nil
}

func (c *Client) discoverResourceGroups(ctx context.Context) error {
	c.ResourceGroups = make(map[string][]*armresources.ResourceGroup, len(c.subscriptions))
	c.registeredNamespaces = make(map[string]map[string]bool, len(c.subscriptions))

	groupsLock, namespacesLock := sync.Mutex{}, sync.Mutex{}

	errorGroup, gtx := errgroup.WithContext(ctx)
	errorGroup.SetLimit(c.pluginSpec.DiscoveryConcurrency)
	for _, subID := range c.subscriptions {
		subID := subID
		errorGroup.Go(func() error {
			groups, err := c.getResourceGroupsForSubscription(gtx, subID)
			if err != nil {
				return err
			}
			groupsLock.Lock()
			defer groupsLock.Unlock()
			c.ResourceGroups[subID] = groups

			return nil
		})

		errorGroup.Go(func() error {
			providers, err := c.getRegisteredProvidersForSubscription(gtx, subID)
			if err != nil {
				return err
			}

			namespacesLock.Lock()
			defer namespacesLock.Unlock()
			c.registeredNamespaces[subID] = make(map[string]bool)
			for _, p := range providers {
				c.registeredNamespaces[subID][strings.ToLower(*p.Namespace)] = true
			}

			return nil
		})
	}
	return errorGroup.Wait()
}

func getCloudConfigFromSpec(specCloud string) (cloud.Configuration, error) {
	var specCloudToConfig = map[string]cloud.Configuration{
		"AzurePublic":     cloud.AzurePublic,
		"AzureGovernment": cloud.AzureGovernment,
		"AzureChina":      cloud.AzureChina,
	}

	if v, ok := specCloudToConfig[specCloud]; ok {
		return v, nil
	}

	return cloud.Configuration{}, fmt.Errorf("unknown Azure cloud name %q. Supported values are %q", specCloud, maps.Keys(specCloudToConfig))
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	var spec Spec
	var err error
	if err := s.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
	}

	spec.SetDefaults()

	uniqueSubscriptions := funk.Uniq(spec.Subscriptions).([]string)
	c := &Client{
		logger:        logger,
		subscriptions: uniqueSubscriptions,
		pluginSpec:    &spec,
	}

	if spec.CloudName != "" {
		cloudConfig, err := getCloudConfigFromSpec(spec.CloudName)
		if err != nil {
			return nil, err
		}
		c.Options = &arm.ClientOptions{ClientOptions: azcore.ClientOptions{Cloud: cloudConfig}}
	}

	// NewDefaultAzureCredential builds a chain of credentials, and reports errors via the log listener
	// This is currently the way we have to get the errors and report them to the user
	// Any credential that has errors is ignored and the next one in the chain is tried when authenticating
	// So it's useful to report all the errors
	// It's logged as information as we don't know which credential chain the user intended to use
	log.SetEvents(azidentity.EventAuthentication)
	log.SetListener(func(e log.Event, s string) {
		if strings.HasPrefix(s, "NewDefaultAzureCredential failed") {
			c.Logger().Info().Str("azure-sdk-for-go", "azidentity").Msg(s)
		}
	})

	var credsOptions *azidentity.DefaultAzureCredentialOptions
	if c.Options != nil {
		credsOptions = &azidentity.DefaultAzureCredentialOptions{ClientOptions: c.Options.ClientOptions}
	}

	c.Creds, err = azidentity.NewDefaultAzureCredential(credsOptions)
	if err != nil {
		return nil, err
	}

	log.SetListener(nil)

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

	if err := c.discoverResourceGroups(ctx); err != nil {
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
