// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules"

func Armhardwaresecuritymodules() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armhardwaresecuritymodules.NewDedicatedHsmClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armhardwaresecuritymodules())
}