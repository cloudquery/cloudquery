// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"

func Armdatafactory() []Table {
	tables := []Table{
		{
			Name:           "factory",
			Struct:         &armdatafactory.Factory{},
			ResponseStruct: &armdatafactory.FactoriesClientListResponse{},
			Client:         &armdatafactory.FactoriesClient{},
			ListFunc:       (&armdatafactory.FactoriesClient{}).NewListPager,
			NewFunc:        armdatafactory.NewFactoriesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DataFactory/factories",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.DataFactory")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armdatafactory"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdatafactory()...)
}
