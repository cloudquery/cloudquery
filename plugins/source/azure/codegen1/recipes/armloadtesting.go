// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting"

func Armloadtesting() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armloadtesting.NewLoadTestsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting",
		},
		{
			NewFunc: armloadtesting.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting",
		},
		{
			NewFunc: armloadtesting.NewQuotasClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armloadtesting())
}