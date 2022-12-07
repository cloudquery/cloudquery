package client

import (
	"context"
	"errors"
	"fmt"
	"strings"

	// Import all autorest modules
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	subscriptions       []string
	logger              zerolog.Logger
	registeredNamespace map[string]map[string]bool
	resourceGroups      map[string][]*armresources.GenericResourceExpanded
	debugMode bool
	// this is set by table client multiplexer
	SubscriptionId string
	ResourceGroup  string
	Creds          azcore.TokenCredential
	Options        *arm.ClientOptions
}

func getSubscriptions(spec *Spec, creds azcore.TokenCredential) ([]string, error) {
	subscriptions := make([]string, 0)
	if len(spec.Subscriptions) == 0 {
		ctx := context.Background()
		c, err := armsubscription.NewSubscriptionsClient(creds, nil)
		if err != nil {
			return nil, err
		}
		pager := c.NewListPager(nil)
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				return nil, err
			}
			for _, sub := range page.Value {
				subscriptions = append(subscriptions, strings.TrimPrefix(*sub.ID, "/subscriptions/"))
			}
		}
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

	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	subscriptions, err := getSubscriptions(&spec, creds)
	if err != nil {
		return nil, err
	}

	resourceGroups := make(map[string][]*armresources.GenericResourceExpanded, len(subscriptions))
	filter := "resourceType eq 'Microsoft.Resources/resourceGroups'"
	registeredNamespace := make(map[string]map[string]bool, len(subscriptions))
	for _, subID := range subscriptions {
		registeredNamespace[subID] = make(map[string]bool)
		// fmt.Println("sub", sub)
		cl, err := armresources.NewClient(subID, creds, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create resource group client: %w", err)
		}
		pager := cl.NewListPager(&armresources.ClientListOptions{
			Filter: &filter,
		})
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to list resource groups: %w", err)
			}
			if len(page.Value) == 0 {
				continue
			}
			resourceGroups[subID] = append(resourceGroups[subID], page.Value...)
		}

		providerClient, err := armresources.NewProvidersClient(subID, creds, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create provider client: %w", err)
		}
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
					registeredNamespace[subID][*p.Namespace] = true
				}
			}
		}
	}

	return &Client{
		logger:              logger,
		registeredNamespace: registeredNamespace,
		subscriptions:       subscriptions,
		resourceGroups:      resourceGroups,
		Creds:               creds,
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
		logger:         c.logger.With().Str("subscription_id", subscriptionId).Logger(),
		SubscriptionId: subscriptionId,
		Creds:          c.Creds,
		Options:        c.Options,
	}
}

// withSubscription allows multiplexer to create a new client with given subscriptionId
func (c Client) withResourceGroup(name string) *Client {
	return &Client{
		subscriptions:  c.subscriptions,
		logger:         c.logger.With().Str("resource_group", name).Logger(),
		SubscriptionId: c.SubscriptionId,
		ResourceGroup:  name,
		Creds:          c.Creds,
		Options:        c.Options,
	}
}
