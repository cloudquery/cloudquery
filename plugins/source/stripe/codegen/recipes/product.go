package recipes

import (
	"github.com/stripe/stripe-go/v74"
)

func ProductResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &stripe.Product{},
			PKColumns:  []string{"id"},
			Service:    "products",
		},
	}
}
