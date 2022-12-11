// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"

func init() {
	tables := []Table{
		{
			Service:        "armbilling",
			Name:           "accounts",
			Struct:         &armbilling.Account{},
			ResponseStruct: &armbilling.AccountsClientListResponse{},
			Client:         &armbilling.AccountsClient{},
			ListFunc:       (&armbilling.AccountsClient{}).NewListPager,
			NewFunc:        armbilling.NewAccountsClient,
			URL:            "/providers/Microsoft.Billing/billingAccounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Billing)`,
		},
		{
			Service:        "armbilling",
			Name:           "enrollment_accounts",
			Struct:         &armbilling.EnrollmentAccountSummary{},
			ResponseStruct: &armbilling.EnrollmentAccountsClientListResponse{},
			Client:         &armbilling.EnrollmentAccountsClient{},
			ListFunc:       (&armbilling.EnrollmentAccountsClient{}).NewListPager,
			NewFunc:        armbilling.NewEnrollmentAccountsClient,
			URL:            "/providers/Microsoft.Billing/enrollmentAccounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Billing)`,
		},
		{
			Service:        "armbilling",
			Name:           "periods",
			Struct:         &armbilling.Period{},
			ResponseStruct: &armbilling.PeriodsClientListResponse{},
			Client:         &armbilling.PeriodsClient{},
			ListFunc:       (&armbilling.PeriodsClient{}).NewListPager,
			NewFunc:        armbilling.NewPeriodsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Billing)`,
		},
	}
	Tables = append(Tables, tables...)
}
