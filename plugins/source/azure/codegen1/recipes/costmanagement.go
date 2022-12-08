// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"

func Armcostmanagement() []*Table {
	tables := []*Table{
		{
			NewFunc:   armcostmanagement.NewViewsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement",
			URL:       "/providers/Microsoft.CostManagement/views",
			Namespace: "Microsoft.CostManagement",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_CostManagement)`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcostmanagement())
}
