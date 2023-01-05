package recipes

import "github.com/stripe/stripe-go/v74"

var AllResources = []*Resource{
	{
		DataStruct: &stripe.Account{},
	},
	{
		DataStruct: &stripe.Customer{},
	},
	{
		DataStruct: &stripe.Dispute{},
	},
	{
		DataStruct: &stripe.Invoice{},
	},
	{
		DataStruct:  &stripe.InvoiceItem{},
		SkipMocks:   true,
		Service:     "invoices",
		Description: "https://stripe.com/docs/api/invoiceitems",
	},
	{
		DataStruct: &stripe.Product{},
		SkipMocks:  true,
	},
	{
		DataStruct: &stripe.Refund{},
	},
	{
		DataStruct: &stripe.Subscription{},
	},
}
