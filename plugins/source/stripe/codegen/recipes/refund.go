package recipes

import (
	"github.com/stripe/stripe-go/v74"
)

func init() {
	AllResources = append(AllResources, []*Resource{
		{
			DataStruct: &stripe.Refund{},
			PKColumns:  []string{"id"},
			Service:    "refunds",
		},
	}...)
}
