// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"

func Armadvisor() []Table {
	tables := []Table{
		{
			Name:           "resource_recommendation_base",
			Struct:         &armadvisor.ResourceRecommendationBase{},
			ResponseStruct: &armadvisor.RecommendationsClientListResponse{},
			Client:         &armadvisor.RecommendationsClient{},
			ListFunc:       (&armadvisor.RecommendationsClient{}).NewListPager,
			NewFunc:        armadvisor.NewRecommendationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/recommendations",
		},
		{
			Name:           "metadata_entity",
			Struct:         &armadvisor.MetadataEntity{},
			ResponseStruct: &armadvisor.RecommendationMetadataClientListResponse{},
			Client:         &armadvisor.RecommendationMetadataClient{},
			ListFunc:       (&armadvisor.RecommendationMetadataClient{}).NewListPager,
			NewFunc:        armadvisor.NewRecommendationMetadataClient,
			URL:            "/providers/Microsoft.Advisor/metadata",
		},
		{
			Name:           "suppression_contract",
			Struct:         &armadvisor.SuppressionContract{},
			ResponseStruct: &armadvisor.SuppressionsClientListResponse{},
			Client:         &armadvisor.SuppressionsClient{},
			ListFunc:       (&armadvisor.SuppressionsClient{}).NewListPager,
			NewFunc:        armadvisor.NewSuppressionsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/suppressions",
		},
	}

	for i := range tables {
		tables[i].Service = "armadvisor"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armadvisor()...)
}
