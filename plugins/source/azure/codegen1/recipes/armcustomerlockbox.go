// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"

func Armcustomerlockbox() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armcustomerlockbox.NewGetClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox",
		},
		{
			NewFunc: armcustomerlockbox.NewPostClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox",
		},
		{
			NewFunc: armcustomerlockbox.NewRequestsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox",
		},
		{
			NewFunc: armcustomerlockbox.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armcustomerlockbox())
}