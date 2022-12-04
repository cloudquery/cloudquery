// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customproviders/armcustomproviders"

func Armcustomproviders() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armcustomproviders.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customproviders/armcustomproviders",
		},
		{
			NewFunc: armcustomproviders.NewAssociationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customproviders/armcustomproviders",
		},
		{
			NewFunc: armcustomproviders.NewCustomResourceProviderClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customproviders/armcustomproviders",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armcustomproviders())
}