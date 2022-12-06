// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"

func Armapplicationinsights() []Table {
	tables := []Table{
		{
			Name:           "web_test",
			Struct:         &armapplicationinsights.WebTest{},
			ResponseStruct: &armapplicationinsights.WebTestsClientListResponse{},
			Client:         &armapplicationinsights.WebTestsClient{},
			ListFunc:       (&armapplicationinsights.WebTestsClient{}).NewListPager,
			NewFunc:        armapplicationinsights.NewWebTestsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/webtests",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Insights")`,
		},
		{
			Name:           "component",
			Struct:         &armapplicationinsights.Component{},
			ResponseStruct: &armapplicationinsights.ComponentsClientListResponse{},
			Client:         &armapplicationinsights.ComponentsClient{},
			ListFunc:       (&armapplicationinsights.ComponentsClient{}).NewListPager,
			NewFunc:        armapplicationinsights.NewComponentsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/components",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Insights")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armapplicationinsights"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armapplicationinsights()...)
}
