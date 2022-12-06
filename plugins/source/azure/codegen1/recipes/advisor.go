// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"

func Armadvisor() []*Table {
	tables := []*Table{
		{
			NewFunc:   armadvisor.NewRecommendationsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/recommendations",
			Namespace: "Microsoft.Advisor",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Advisor")`,
		},
		{
			NewFunc:   armadvisor.NewRecommendationMetadataClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor",
			URL:       "/providers/Microsoft.Advisor/metadata",
			Namespace: "Microsoft.Advisor",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Advisor")`,
		},
		{
			NewFunc:   armadvisor.NewSuppressionsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/suppressions",
			Namespace: "Microsoft.Advisor",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Advisor")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armadvisor())
}
