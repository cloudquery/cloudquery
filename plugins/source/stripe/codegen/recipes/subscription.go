package recipes

import (
	"github.com/stripe/stripe-go/v74"
)

func init() {
	AllResources = append(AllResources, []*Resource{
		{
			DataStruct: &stripe.Subscription{},
			PKColumns:  []string{"id"},
		},
	}...)
}
