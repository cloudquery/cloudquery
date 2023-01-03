package recipes

import (
	"github.com/stripe/stripe-go/v74"
)

func AccountResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &stripe.Account{},
			PKColumns:  []string{"id"},
			Service:    "accounts",
		},
	}
}
