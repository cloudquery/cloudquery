// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"

func Armmanagedservices() []*Table {
	tables := []*Table{
		{
			NewFunc:   armmanagedservices.NewMarketplaceRegistrationDefinitionsWithoutScopeClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices",
			URL:       "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions",
			Namespace: "Microsoft.ManagedServices",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.ManagedServices")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armmanagedservices())
}
