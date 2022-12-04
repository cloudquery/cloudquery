// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality"

func Armmixedreality() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmixedreality.NewObjectAnchorsAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality",
		},
		{
			NewFunc: armmixedreality.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality",
		},
		{
			NewFunc: armmixedreality.NewClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality",
		},
		{
			NewFunc: armmixedreality.NewRemoteRenderingAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality",
		},
		{
			NewFunc: armmixedreality.NewSpatialAnchorsAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmixedreality())
}