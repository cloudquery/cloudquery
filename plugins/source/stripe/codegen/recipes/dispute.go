package recipes

import (
	"github.com/stripe/stripe-go/v74"
)

func DisputeResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &stripe.Dispute{},
			PKColumns:  []string{"id"},
			Service:    "disputes",
		},
	}
}
