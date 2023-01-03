package recipes

import (
	"github.com/stripe/stripe-go/v74"
)

func init() {
	AllResources = append(AllResources, []*Resource{
		{
			DataStruct: &stripe.Invoice{},
			PKColumns:  []string{"id"},
			Service:    "invoices",
		},
		{
			DataStruct: &stripe.InvoiceItem{},
			PKColumns:  []string{"id"},
			Service:    "invoices",
			SkipMocks:  true,
		},
	}...)
}
