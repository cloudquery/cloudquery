// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementpartner/armmanagementpartner"

func Armmanagementpartner() []*Table {
	tables := []*Table{
		{
			NewFunc:   armmanagementpartner.NewOperationClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementpartner/armmanagementpartner",
			URL:       "/providers/Microsoft.ManagementPartner/operations",
			Namespace: "Microsoft.ManagementPartner",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ManagementPartner)`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armmanagementpartner())
}
