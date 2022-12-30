package recipes

import (
	"github.com/stripe/stripe-go/v74"
)

func SubscriptionResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &stripe.Subscription{},
			PKColumns:  []string{"id"},
			Service:    "subscriptions",
		},
	}
}
