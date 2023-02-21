package recipes

import "github.com/stripe/stripe-go/v74"

const createdStateParam = "Created"

var AllResources = []*Resource{
	{
		DataStruct:     &stripe.Account{},
		StateParamName: createdStateParam,
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
		DataStruct:     &stripe.ApplicationFee{},
		StateParamName: createdStateParam,
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
		DataStruct:     &stripe.BalanceTransaction{},
		Service:        "balance",
		StateParamName: createdStateParam,
	},
	{
		DataStruct: &stripe.BillingPortalConfiguration{},
		Service:    "billing_portal",
	},
	{
		DataStruct:     &stripe.Charge{},
		IgnoreInTests:  []string{"Destination", "Dispute", "Level3", "Source"},
		StateParamName: createdStateParam,
	},
	{
		DataStruct: &stripe.CheckoutSession{},
		Service:    "checkout",
	},
	{
		DataStruct: &stripe.CountrySpec{},
	},
	{
		DataStruct:     &stripe.Coupon{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct: &stripe.CreditNote{},
	},
	{
		DataStruct:     &stripe.Customer{},
		IgnoreInTests:  []string{"DefaultSource"},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.Dispute{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.Event{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.FileLink{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.File{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.IdentityVerificationReport{},
		Service:        "identity",
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.Invoice{},
		IgnoreInTests:  []string{"DefaultSource"},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:    &stripe.InvoiceItem{},
		Service:       "invoices",
		Description:   "https://stripe.com/docs/api/invoiceitems",
		IgnoreInTests: []string{"Plan"},
	},
	{
		DataStruct:     &stripe.IssuingAuthorization{},
		Service:        "issuing",
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.IssuingCardholder{},
		Service:        "issuing",
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.IssuingCard{},
		Service:        "issuing",
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.IssuingDispute{},
		Service:        "issuing",
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.IssuingTransaction{},
		Service:        "issuing",
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.PaymentIntent{},
		Service:        "payment",
		IgnoreInTests:  []string{"Source"},
		StateParamName: createdStateParam,
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
		DataStruct:     &stripe.Payout{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.Plan{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.Price{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.Product{},
		IgnoreInTests:  []string{"Attributes", "DeactivateOn"},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.PromotionCode{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct: &stripe.Quote{},
	},
	{
		DataStruct: &stripe.RadarEarlyFraudWarning{},
		Service:    "radar",
	},
	{
		DataStruct:     &stripe.Refund{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.ReportingReportRun{},
		Service:        "reporting",
		StateParamName: createdStateParam,
	},
	{
		DataStruct: &stripe.ReportingReportType{},
		Service:    "reporting",
	},
	{
		DataStruct:     &stripe.Review{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.ShippingRate{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct: &stripe.SigmaScheduledQueryRun{},
		Service:    "sigma",
	},
	{
		DataStruct:     &stripe.Subscription{},
		Service:        "subscription",
		IgnoreInTests:  []string{"DefaultSource"},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.SubscriptionSchedule{},
		Service:        "subscription",
		StateParamName: createdStateParam,
	},
	{
		DataStruct: &stripe.TaxCode{},
		Service:    "tax",
	},
	{
		DataStruct:     &stripe.TaxRate{},
		Service:        "tax",
		StateParamName: createdStateParam,
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
		DataStruct:     &stripe.Topup{},
		IgnoreInTests:  []string{"Source"},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.Transfer{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.TreasuryFinancialAccount{},
		Service:        "treasury",
		StateParamName: createdStateParam,
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
				DataStruct:     &stripe.TreasuryTransactionEntry{},
				ListParams:     `FinancialAccount: stripe.String(p.ID),`,
				StateParamName: createdStateParam,
			},
			{
				DataStruct:     &stripe.TreasuryTransaction{},
				ListParams:     `FinancialAccount: stripe.String(p.ID),`,
				StateParamName: createdStateParam,
			},
		},
	},
	{
		DataStruct: &stripe.WebhookEndpoint{},
	},
}
