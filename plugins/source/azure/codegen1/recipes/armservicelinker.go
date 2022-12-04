// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker"

func Armservicelinker() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armservicelinker.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker",
		},
		{
			NewFunc: armservicelinker.NewLinkerClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armservicelinker())
}