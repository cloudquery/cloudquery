package recipes

import (
	"github.com/stripe/stripe-go/v74"
)

func init() {
	AllResources = append(AllResources, []*Resource{
		{
			DataStruct: &stripe.Product{},
			PKColumns:  []string{"id"},
			Service:    "products",
			SkipMocks:  true,
		},
	}...)
}
