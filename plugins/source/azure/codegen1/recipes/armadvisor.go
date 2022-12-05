// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"

func Armadvisor() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armadvisor.NewConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor",
			URL: "",
		},
		{
			NewFunc: armadvisor.NewRecommendationMetadataClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor",
			URL: "/providers/Microsoft.Advisor/metadata",
		},
		{
			NewFunc: armadvisor.NewRecommendationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/recommendations",
		},
		{
			NewFunc: armadvisor.NewSuppressionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/suppressions",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armadvisor())
}