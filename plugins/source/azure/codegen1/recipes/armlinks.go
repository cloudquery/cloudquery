// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks"

func Armlinks() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armlinks.NewResourceLinksClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armlinks())
}