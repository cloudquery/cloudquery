package recipes

import (
	"github.com/stripe/stripe-go/v74"
)

func init() {
	AllResources = append(AllResources, []*Resource{
		{
			DataStruct: &stripe.Invoice{},
			PKColumns:  []string{"id"},
		},
		{
			DataStruct:  &stripe.InvoiceItem{},
			PKColumns:   []string{"id"},
			SkipMocks:   true,
			Service:     "invoices",
			Description: "https://stripe.com/docs/api/invoiceitem",
		},
	}...)
}
