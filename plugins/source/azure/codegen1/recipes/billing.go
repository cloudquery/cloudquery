// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"

func Armbilling() []*Table {
	tables := []*Table{
		{
			NewFunc: armbilling.NewEnrollmentAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling",
			URL: "/providers/Microsoft.Billing/enrollmentAccounts",
		},
		{
			NewFunc: armbilling.NewAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling",
			URL: "/providers/Microsoft.Billing/billingAccounts",
		},
		{
			NewFunc: armbilling.NewPeriodsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armbilling())
}