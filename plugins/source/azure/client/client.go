package client

import (
	"context"
	"errors"
	"fmt"

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
	subscriptions  []string
	logger         zerolog.Logger
	resourceGroups map[string][]*armresources.GenericResourceExpanded
	// this is set by table client multiplexer
	SubscriptionId string
	ResourceGroup  string
	Creds          azcore.TokenCredential
	Options        *arm.ClientOptions
}

func getSubscriptions(logger *zerolog.Logger, spec *Spec, creds azcore.TokenCredential) ([]string, error) {
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
				subscriptions = append(subscriptions, *sub.ID)
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

	subscriptions, err := getSubscriptions(&logger, &spec, creds)
	if err != nil {
		return nil, err
	}
	resourceGroups := make(map[string][]*armresources.GenericResourceExpanded, len(subscriptions))
	fiilter := "$filter=Microsoft.Resources/resourceGroups"
	for _, sub := range subscriptions {
		cl, err := armresources.NewClient(sub, creds, nil)
		if err != nil {
			return nil, err
		}
		pager := cl.NewListPager(&armresources.ClientListOptions{
			Filter: &fiilter,
		})
		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				return nil, err
			}
			resourceGroups[sub] = append(resourceGroups[sub], page.Value...)
		}
	}

	return &Client{
		logger:         logger,
		subscriptions:  subscriptions,
		resourceGroups: resourceGroups,
		Creds:          creds,
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
