package recipes

import "github.com/stripe/stripe-go/v74"

var AllResources = []*Resource{
	{
		DataStruct: &stripe.Account{},
		Children: []*Resource{
			{
				DataStruct: &stripe.Capability{},
				ListParams: `Account: stripe.String(p.ID),`,
			},
		},
	},
	{
		DataStruct: &stripe.ApplePayDomain{},
	},
	{
		DataStruct: &stripe.ApplicationFee{},
		Children: []*Resource{
			{
				DataStruct: &stripe.FeeRefund{},
				ListParams: `ID: stripe.String(p.ID),`,
			},
		},
	},
	{
		DataStruct: &stripe.Balance{},
		Single:     true,
		Service:    "balance",
	},
	{
		DataStruct: &stripe.BalanceTransaction{},
		Service:    "balance",
	},
	{
		DataStruct: &stripe.BillingPortalConfiguration{},
		Service:    "billing_portal",
	},
	{
		DataStruct:    &stripe.Charge{},
		IgnoreInTests: []string{"Destination", "Dispute", "Level3", "Source"},
	},
	{
		DataStruct: &stripe.CheckoutSession{},
		Service:    "checkout",
	},
	{
		DataStruct: &stripe.CountrySpec{},
	},
	{
		DataStruct: &stripe.Coupon{},
	},
	{
		DataStruct: &stripe.CreditNote{},
	},
	{
		DataStruct:    &stripe.Customer{},
		IgnoreInTests: []string{"DefaultSource"},
	},
	{
		DataStruct: &stripe.Dispute{},
	},
	{
		DataStruct: &stripe.FileLink{},
	},
	{
		DataStruct: &stripe.File{},
	},
	{
		DataStruct: &stripe.IdentityVerificationReport{},
		Service:    "identity",
	},
	{
		DataStruct:    &stripe.Invoice{},
		IgnoreInTests: []string{"DefaultSource"},
	},
	{
		DataStruct:    &stripe.InvoiceItem{},
		Service:       "invoices",
		Description:   "https://stripe.com/docs/api/invoiceitems",
		IgnoreInTests: []string{"Plan"},
	},
	{
		DataStruct: &stripe.IssuingAuthorization{},
		Service:    "issuing",
	},
	{
		DataStruct: &stripe.IssuingCardholder{},
		Service:    "issuing",
	},
	{
		DataStruct: &stripe.IssuingCard{},
		Service:    "issuing",
	},
	{
		DataStruct: &stripe.IssuingDispute{},
		Service:    "issuing",
	},
	{
		DataStruct: &stripe.IssuingTransaction{},
		Service:    "issuing",
	},
	{
		DataStruct:    &stripe.PaymentIntent{},
		Service:       "payment",
		IgnoreInTests: []string{"Source"},
	},
	{
		DataStruct: &stripe.PaymentLink{},
		Service:    "payment",
	},
	{
		DataStruct: &stripe.PaymentMethod{},
		Service:    "payment",
	},
	{
		DataStruct: &stripe.Payout{},
	},
	{
		DataStruct: &stripe.Plan{},
	},
	{
		DataStruct: &stripe.Price{},
	},
	{
		DataStruct:    &stripe.Product{},
		IgnoreInTests: []string{"Attributes", "DeactivateOn"},
	},
	{
		DataStruct: &stripe.PromotionCode{},
	},
	{
		DataStruct: &stripe.Quote{},
	},
	{
		DataStruct: &stripe.RadarEarlyFraudWarning{},
		Service:    "radar",
	},
	{
		DataStruct: &stripe.Refund{},
	},
	{
		DataStruct: &stripe.ReportingReportRun{},
		Service:    "reporting",
	},
	{
		DataStruct: &stripe.ReportingReportType{},
		Service:    "reporting",
	},
	{
		DataStruct: &stripe.Review{},
	},
	{
		DataStruct: &stripe.ShippingRate{},
	},
	{
		DataStruct: &stripe.SigmaScheduledQueryRun{},
		Service:    "sigma",
	},
	{
		DataStruct:    &stripe.Subscription{},
		Service:       "subscription",
		IgnoreInTests: []string{"DefaultSource"},
	},
	{
		DataStruct: &stripe.SubscriptionSchedule{},
		Service:    "subscription",
	},
	{
		DataStruct: &stripe.TaxCode{},
		Service:    "tax",
	},
	{
		DataStruct: &stripe.TaxRate{},
		Service:    "tax",
	},
	{
		DataStruct: &stripe.TerminalConfiguration{},
		Service:    "terminal",
	},
	{
		DataStruct: &stripe.TerminalLocation{},
		Service:    "terminal",
	},
	{
		DataStruct: &stripe.TerminalReader{},
		Service:    "terminal",
	},
	{
		DataStruct:    &stripe.Topup{},
		IgnoreInTests: []string{"Source"},
	},
	{
		DataStruct: &stripe.Transfer{},
	},
	{
		DataStruct: &stripe.TreasuryFinancialAccount{},
		Service:    "treasury",
		Children: []*Resource{
			{
				DataStruct: &stripe.TreasuryCreditReversal{},
				ListParams: `FinancialAccount: stripe.String(p.ID),`,
			},
			{
				DataStruct: &stripe.TreasuryDebitReversal{},
				ListParams: `FinancialAccount: stripe.String(p.ID),`,
			},
			{
				DataStruct: &stripe.TreasuryInboundTransfer{},
				ListParams: `FinancialAccount: stripe.String(p.ID),`,
			},
			{
				DataStruct: &stripe.TreasuryOutboundPayment{},
				ListParams: `FinancialAccount: stripe.String(p.ID),`,
			},
			{
				DataStruct: &stripe.TreasuryOutboundTransfer{},
				ListParams: `FinancialAccount: stripe.String(p.ID),`,
			},
			{
				DataStruct: &stripe.TreasuryReceivedCredit{},
				ListParams: `FinancialAccount: stripe.String(p.ID),`,
			},
			{
				DataStruct: &stripe.TreasuryReceivedDebit{},
				ListParams: `FinancialAccount: stripe.String(p.ID),`,
			},
			{
				DataStruct: &stripe.TreasuryTransactionEntry{},
				ListParams: `FinancialAccount: stripe.String(p.ID),`,
			},
			{
				DataStruct: &stripe.TreasuryTransaction{},
				ListParams: `FinancialAccount: stripe.String(p.ID),`,
			},
		},
	},
	{
		DataStruct: &stripe.WebhookEndpoint{},
	},
}
