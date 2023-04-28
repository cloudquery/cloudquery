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
		DataStruct:  &stripe.ApplePayDomain{},
		Description: "https://stripe.com/docs/api",
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
		DataStruct:  &stripe.BillingPortalConfiguration{},
		Service:     "billing_portal",
		Description: "https://stripe.com/docs/api/customer_portal/configuration",
	},
	{
		DataStruct:     &stripe.Charge{},
		IgnoreInTests:  []string{"Destination", "Dispute", "Level3", "Source"},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:    &stripe.CheckoutSession{},
		Service:       "checkout",
		Description:   "https://stripe.com/docs/api/checkout/sessions",
		ExtraChildren: []string{"CheckoutSessionLineItems()"},
		/*
			Children: []*Resource{
				{
					DataStruct:  &stripe.LineItem{},
					ListParams:  `Session: stripe.String(p.ID),`,
					Description: "https://stripe.com/docs/api/checkout/sessions/line_items",
				},
			},
		*/
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
		Description:    "https://stripe.com/docs/api/identity/verification_reports",
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
		Description:    "https://stripe.com/docs/api/issuing/authorizations",
	},
	{
		DataStruct:     &stripe.IssuingCardholder{},
		Service:        "issuing",
		StateParamName: createdStateParam,
		Description:    "https://stripe.com/docs/api/issuing/cardholders",
	},
	{
		DataStruct:     &stripe.IssuingCard{},
		Service:        "issuing",
		StateParamName: createdStateParam,
		Description:    "https://stripe.com/docs/api/issuing/cards",
	},
	{
		DataStruct:     &stripe.IssuingDispute{},
		Service:        "issuing",
		StateParamName: createdStateParam,
		Description:    "https://stripe.com/docs/api/issuing/disputes",
	},
	{
		DataStruct:     &stripe.IssuingTransaction{},
		Service:        "issuing",
		StateParamName: createdStateParam,
		Description:    "https://stripe.com/docs/api/issuing/transactions",
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
		ExpandFields:   []string{"data.currency_options", "data.tiers"},
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
		DataStruct:  &stripe.RadarEarlyFraudWarning{},
		Service:     "radar",
		Description: "https://stripe.com/docs/api/radar/early_fraud_warnings",
	},
	{
		DataStruct:     &stripe.Refund{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:     &stripe.ReportingReportRun{},
		Service:        "reporting",
		StateParamName: createdStateParam,
		Description:    "https://stripe.com/docs/api/reporting/report_run",
	},
	{
		DataStruct:  &stripe.ReportingReportType{},
		Service:     "reporting",
		Description: "https://stripe.com/docs/api/reporting/report_type",
	},
	{
		DataStruct:     &stripe.Review{},
		StateParamName: createdStateParam,
		Description:    "https://stripe.com/docs/api/radar/reviews",
	},
	{
		DataStruct:     &stripe.ShippingRate{},
		StateParamName: createdStateParam,
	},
	{
		DataStruct:  &stripe.SigmaScheduledQueryRun{},
		Service:     "sigma",
		Description: "https://stripe.com/docs/api/sigma/scheduled_queries",
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
		DataStruct:  &stripe.TerminalConfiguration{},
		Service:     "terminal",
		Description: "https://stripe.com/docs/api/terminal/configuration",
	},
	{
		DataStruct:  &stripe.TerminalLocation{},
		Service:     "terminal",
		Description: "https://stripe.com/docs/api/terminal/locations",
	},
	{
		DataStruct:  &stripe.TerminalReader{},
		Service:     "terminal",
		Description: "https://stripe.com/docs/api/terminal/readers",
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
		Description:    "https://stripe.com/docs/api/treasury/financial_accounts",
		Children: []*Resource{
			{
				DataStruct:  &stripe.TreasuryCreditReversal{},
				ListParams:  `FinancialAccount: stripe.String(p.ID),`,
				Description: "https://stripe.com/docs/api/treasury/credit_reversals",
			},
			{
				DataStruct:  &stripe.TreasuryDebitReversal{},
				ListParams:  `FinancialAccount: stripe.String(p.ID),`,
				Description: "https://stripe.com/docs/api/treasury/debit_reversals",
			},
			{
				DataStruct:  &stripe.TreasuryInboundTransfer{},
				ListParams:  `FinancialAccount: stripe.String(p.ID),`,
				Description: "https://stripe.com/docs/api/treasury/inbound_transfers",
			},
			{
				DataStruct:  &stripe.TreasuryOutboundPayment{},
				ListParams:  `FinancialAccount: stripe.String(p.ID),`,
				Description: "https://stripe.com/docs/api/treasury/outbound_payments",
			},
			{
				DataStruct:  &stripe.TreasuryOutboundTransfer{},
				ListParams:  `FinancialAccount: stripe.String(p.ID),`,
				Description: "https://stripe.com/docs/api/treasury/outbound_transfers",
			},
			{
				DataStruct:  &stripe.TreasuryReceivedCredit{},
				ListParams:  `FinancialAccount: stripe.String(p.ID),`,
				Description: "https://stripe.com/docs/api/treasury/received_credits",
			},
			{
				DataStruct:  &stripe.TreasuryReceivedDebit{},
				ListParams:  `FinancialAccount: stripe.String(p.ID),`,
				Description: "https://stripe.com/docs/api/treasury/received_debits",
			},
			{
				DataStruct:     &stripe.TreasuryTransactionEntry{},
				ListParams:     `FinancialAccount: stripe.String(p.ID),`,
				StateParamName: createdStateParam,
				Description:    "https://stripe.com/docs/api/treasury/transaction_entries",
			},
			{
				DataStruct:     &stripe.TreasuryTransaction{},
				ListParams:     `FinancialAccount: stripe.String(p.ID),`,
				StateParamName: createdStateParam,
				Description:    "https://stripe.com/docs/api/treasury/transactions",
			},
		},
	},
	{
		DataStruct: &stripe.WebhookEndpoint{},
	},
}
