// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"

func Armadvisor() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armadvisor.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor",
		},
		{
			NewFunc: armadvisor.NewRecommendationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor",
		},
		{
			NewFunc: armadvisor.NewSuppressionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor",
		},
		{
			NewFunc: armadvisor.NewConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor",
		},
		{
			NewFunc: armadvisor.NewRecommendationMetadataClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armadvisor())
}