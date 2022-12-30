package recipes

import (
	"github.com/stripe/stripe-go/v74"
)

func RefundResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &stripe.Refund{},
			PKColumns:  []string{"id"},
			Service:    "refunds",
		},
	}
}
