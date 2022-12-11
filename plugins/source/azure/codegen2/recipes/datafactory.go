// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"

func init() {
	tables := []Table{
		{
			Service:        "armdatafactory",
			Name:           "factories",
			Struct:         &armdatafactory.Factory{},
			ResponseStruct: &armdatafactory.FactoriesClientListResponse{},
			Client:         &armdatafactory.FactoriesClient{},
			ListFunc:       (&armdatafactory.FactoriesClient{}).NewListPager,
			NewFunc:        armdatafactory.NewFactoriesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DataFactory/factories",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DataFactory)`,
		},
	}
	Tables = append(Tables, tables...)
}
