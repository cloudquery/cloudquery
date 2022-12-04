// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"

func Armquota() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armquota.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota",
		},
		{
			NewFunc: armquota.NewRequestStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota",
		},
		{
			NewFunc: armquota.NewUsagesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota",
		},
		{
			NewFunc: armquota.NewOperationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armquota())
}