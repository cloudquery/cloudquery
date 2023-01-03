package recipes

import (
	"github.com/stripe/stripe-go/v74"
)

func init() {
	AllResources = append(AllResources, []*Resource{
		{
			DataStruct: &stripe.Dispute{},
			PKColumns:  []string{"id"},
			Service:    "disputes",
		},
	}...)
}
