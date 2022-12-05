// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"

func Armmsi() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmsi.NewFederatedIdentityCredentialsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials",
		},
		{
			NewFunc: armmsi.NewSystemAssignedIdentitiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi",
			URL: "",
		},
		{
			NewFunc: armmsi.NewUserAssignedIdentitiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmsi())
}