// Code generated by codegen; DO NOT EDIT.
package packages

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization"

func Armauthorization() []*Table {
	tables := []*Table{
		{
			NewFunc: armauthorization.NewProviderOperationsMetadataClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization",
			URL: "/providers/Microsoft.Authorization/providerOperations",
		},
		{
			NewFunc: armauthorization.NewRoleAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/roleAssignments",
		},
		{
			NewFunc: armauthorization.NewClassicAdministratorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/classicAdministrators",
		},
		{
			NewFunc: armauthorization.NewRoleDefinitionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization",
			URL: "/{scope}/providers/Microsoft.Authorization/roleDefinitions",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armauthorization())
}