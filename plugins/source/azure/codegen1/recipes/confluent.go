// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"

func Armconfluent() []*Table {
	tables := []*Table{
		{
			NewFunc: armconfluent.NewMarketplaceAgreementsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Confluent/agreements",
		},
		{
			NewFunc: armconfluent.NewOrganizationOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent",
			URL: "/providers/Microsoft.Confluent/operations",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armconfluent())
}