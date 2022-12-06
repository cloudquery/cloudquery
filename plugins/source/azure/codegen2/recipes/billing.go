// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"

func Armbilling() []Table {
	tables := []Table{
		{
      Name: "enrollment_account_summary",
      Struct: &armbilling.EnrollmentAccountSummary{},
      ResponseStruct: &armbilling.EnrollmentAccountsClientListResponse{},
      Client: &armbilling.EnrollmentAccountsClient{},
      ListFunc: (&armbilling.EnrollmentAccountsClient{}).NewListPager,
			NewFunc: armbilling.NewEnrollmentAccountsClient,
			URL: "/providers/Microsoft.Billing/enrollmentAccounts",
		},
		{
      Name: "period",
      Struct: &armbilling.Period{},
      ResponseStruct: &armbilling.PeriodsClientListResponse{},
      Client: &armbilling.PeriodsClient{},
      ListFunc: (&armbilling.PeriodsClient{}).NewListPager,
			NewFunc: armbilling.NewPeriodsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods",
		},
		{
      Name: "account",
      Struct: &armbilling.Account{},
      ResponseStruct: &armbilling.AccountsClientListResponse{},
      Client: &armbilling.AccountsClient{},
      ListFunc: (&armbilling.AccountsClient{}).NewListPager,
			NewFunc: armbilling.NewAccountsClient,
			URL: "/providers/Microsoft.Billing/billingAccounts",
		},
	}

	for i := range tables {
		tables[i].Service = "armbilling"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armbilling()...)
}