// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"

func Armsaas() []*Table {
	tables := []*Table{
		{
			NewFunc:   armsaas.NewApplicationsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas",
			URL:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/applications",
			Namespace: "Microsoft.SaaS",
			Multiplex: `client.SubscriptionResourceGroupMultiplexRegisteredNamespace("Microsoft.SaaS")`,
		},
		{
			NewFunc:   armsaas.NewResourcesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas",
			URL:       "/providers/Microsoft.SaaS/saasresources",
			Namespace: "Microsoft.SaaS",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.SaaS")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armsaas())
}
