// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsiot/armwindowsiot"

func Armwindowsiot() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armwindowsiot.NewServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsiot/armwindowsiot",
		},
		{
			NewFunc: armwindowsiot.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsiot/armwindowsiot",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armwindowsiot())
}