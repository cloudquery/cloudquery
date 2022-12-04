// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"

func Armmaps() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmaps.NewAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps",
		},
		{
			NewFunc: armmaps.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps",
		},
		{
			NewFunc: armmaps.NewCreatorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmaps())
}