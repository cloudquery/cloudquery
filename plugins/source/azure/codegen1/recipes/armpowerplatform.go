// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform"

func Armpowerplatform() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armpowerplatform.NewAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform",
		},
		{
			NewFunc: armpowerplatform.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform",
		},
		{
			NewFunc: armpowerplatform.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform",
		},
		{
			NewFunc: armpowerplatform.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform",
		},
		{
			NewFunc: armpowerplatform.NewEnterprisePoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armpowerplatform())
}