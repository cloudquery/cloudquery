// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"

func Armapplicationinsights() []Table {
	tables := []Table{
		{
			Service:        "armapplicationinsights",
			Name:           "components",
			Struct:         &armapplicationinsights.Component{},
			ResponseStruct: &armapplicationinsights.ComponentsClientListResponse{},
			Client:         &armapplicationinsights.ComponentsClient{},
			ListFunc:       (&armapplicationinsights.ComponentsClient{}).NewListPager,
			NewFunc:        armapplicationinsights.NewComponentsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/components",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Insights)`,
		},
		{
			Service:        "armapplicationinsights",
			Name:           "web_tests",
			Struct:         &armapplicationinsights.WebTest{},
			ResponseStruct: &armapplicationinsights.WebTestsClientListResponse{},
			Client:         &armapplicationinsights.WebTestsClient{},
			ListFunc:       (&armapplicationinsights.WebTestsClient{}).NewListPager,
			NewFunc:        armapplicationinsights.NewWebTestsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/webtests",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Insights)`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, Armapplicationinsights()...)
}
