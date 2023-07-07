package client

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	// Import all autorest modules
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/log"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/thoas/go-funk"
	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"
)

type Client struct {
	subscriptions []string

	// SubscriptionsObjects is to cache full objects returned from ListSubscriptions on initialisation
	SubscriptionsObjects []*armsubscriptions.Subscription

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

	pluginSpec      *Spec
	BillingAccounts []*armbilling.Account
	BillingAccount  *armbilling.Account
	BillingProfile  *armbilling.Profile
	BillingPeriods  map[string][]*armbilling.Period
	BillingPeriod   *armbilling.Period

	storageAccountKeys *sync.Map
}

func (c *Client) discoverSubscriptions(ctx context.Context) error {
	c.subscriptions = make([]string, 0)
	subscriptionClient, err := armsubscriptions.NewClient(c.Creds, c.Options)
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
			if *sub.State == armsubscriptions.SubscriptionStateEnabled {
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

func (c *Client) discoverBillingAccounts(ctx context.Context) error {
	accounts := make([]*armbilling.Account, 0)
	svc, err := armbilling.NewAccountsClient(c.Creds, c.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(&armbilling.AccountsClientListOptions{Expand: to.Ptr("soldTo,billingProfiles,billingProfiles/invoiceSections")})
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		accounts = append(accounts, p.Value...)
	}
	c.BillingAccounts = accounts
	return nil
}

func (c *Client) discoverBillingPeriods(ctx context.Context) error {
	billingPeriods := make(map[string][]*armbilling.Period, len(c.subscriptions))
	errorGroup, gtx := errgroup.WithContext(ctx)
	errorGroup.SetLimit(c.pluginSpec.DiscoveryConcurrency)

	periodsLock := sync.Mutex{}

	for _, subID := range c.subscriptions {
		subID := subID
		errorGroup.Go(func() error {
			periods := make([]*armbilling.Period, 0)
			svc, err := armbilling.NewPeriodsClient(subID, c.Creds, c.Options)
			if err != nil {
				return err
			}
			pager := svc.NewListPager(nil)
			for pager.More() {
				p, err := pager.NextPage(gtx)
				if err != nil {
					return err
				}
				periods = append(periods, p.Value...)
			}

			periodsLock.Lock()
			defer periodsLock.Unlock()
			billingPeriods[subID] = periods

			return nil
		})
	}
	err := errorGroup.Wait()
	if err != nil {
		return err
	}
	c.BillingPeriods = billingPeriods
	return nil
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

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (schema.ClientMeta, error) {
	s.SetDefaults()
	uniqueSubscriptions := funk.Uniq(s.Subscriptions).([]string)
	c := &Client{
		logger:             logger,
		subscriptions:      uniqueSubscriptions,
		pluginSpec:         s,
		storageAccountKeys: &sync.Map{},
	}

	if s.CloudName != "" {
		cloudConfig, err := getCloudConfigFromSpec(s.CloudName)
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
	var err error
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
	// User specified subscriptions, that CloudQuery should skip syncing
	c.subscriptions = funk.LeftJoinString(c.subscriptions, s.SkipSubscriptions)

	if len(c.subscriptions) == 0 {
		return nil, fmt.Errorf("no subscriptions found")
	}

	if err := c.discoverResourceGroups(ctx); err != nil {
		return nil, err
	}

	if err := c.discoverBillingAccounts(ctx); err != nil {
		c.logger.Warn().Err(err).Msg("failed to discover billing accounts (skipping)")
	}

	if err := c.discoverBillingPeriods(ctx); err != nil {
		c.logger.Warn().Err(err).Msg("failed to discover billing periods (skipping)")
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
	if c.BillingProfile != nil {
		return fmt.Sprintf("billingAccounts/%s/billingProfiles/%s", *c.BillingAccount.Name, *c.BillingProfile.Name)
	}
	if c.BillingAccount != nil {
		return fmt.Sprintf("billingAccounts/%s", *c.BillingAccount.Name)
	}
	if c.BillingPeriod != nil {
		return fmt.Sprintf("subscriptions/%s/billingPeriods/%s", c.SubscriptionId, *c.BillingPeriod.Name)
	}
	return fmt.Sprintf("subscriptions/%s", c.SubscriptionId)
}

func (c *Client) Duplicate() *Client {
	newClient := *c
	return &newClient
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

func (c *Client) withBillingAccount(billingAccount *armbilling.Account) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("billing_account", *billingAccount.ID).Logger()
	newC.BillingAccount = billingAccount
	return &newC
}

func (c *Client) withBillingProfile(billingProfile *armbilling.Profile) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("billing_profile", *billingProfile.ID).Logger()
	newC.BillingProfile = billingProfile
	return &newC
}

func (c *Client) withBillingPeriod(billingPeriod *armbilling.Period) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("billing_period", *billingPeriod.ID).Logger()
	newC.BillingPeriod = billingPeriod
	return &newC
}

var ErrNoStorageKeysFound = errors.New("no storage keys found")

func (c *Client) GetStorageAccountKey(ctx context.Context, acc *armstorage.Account) (string, error) {
	key, err := loadOrStore(c.storageAccountKeys, *acc.Name, func() (any, error) {
		svc, err := armstorage.NewAccountsClient(c.SubscriptionId, c.Creds, c.Options)
		if err != nil {
			return nil, err
		}

		group, err := ParseResourceGroup(*acc.ID)
		if err != nil {
			return nil, err
		}

		keysResponse, err := svc.ListKeys(ctx, group, *acc.Name, nil)
		if err != nil {
			return nil, err
		}

		if len(keysResponse.Keys) == 0 {
			return nil, ErrNoStorageKeysFound
		}
		return *keysResponse.Keys[0].Value, nil
	})
	if key == nil {
		return "", err
	}
	return key.(string), err
}
