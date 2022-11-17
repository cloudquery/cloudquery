package client

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/rs/zerolog"
)

func getSubscriptions(ctx context.Context, logger *zerolog.Logger, spec *Spec, credentials azcore.TokenCredential) ([]string, error) {
	// limited by spec
	if len(spec.Subscriptions) > 0 {
		return spec.Subscriptions, nil
	}

	var subscriptions []string
	svc, err := armsubscription.NewSubscriptionsClient(credentials, nil)
	if err != nil {
		return nil, err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		res, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, sub := range res.Value {
			switch *sub.State {
			case armsubscription.SubscriptionStateDisabled:
				logger.Info().Msgf("Not fetching from subscription because it is disabled %s - %s", "subscription", *sub.SubscriptionID)
			case armsubscription.SubscriptionStateDeleted:
				logger.Info().Msgf("Not fetching from subscription because it is deleted %s - %s", "subscription", *sub.SubscriptionID)
			default:
				subscriptions = append(subscriptions, *sub.SubscriptionID)
			}
		}
	}
	logger.Info().Msgf("No subscriptions specified, going to using all available ones %s %s", "subscriptions", subscriptions)

	if len(subscriptions) == 0 {
		return nil, errors.New("could not find any subscription")
	}

	return subscriptions, nil
}
