package recipes

import (
	"github.com/stripe/stripe-go/v74"
)

func CustomerResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &stripe.Customer{},
			PKColumns:  []string{"id"},
			Service:    "customers",
		},
	}
}
