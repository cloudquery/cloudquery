// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"

func Armdatafactory() []*Table {
	tables := []*Table{
		{
			NewFunc:        armdatafactory.NewFactoriesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DataFactory/factories",
			Namespace:      "Microsoft.DataFactory",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DataFactory)`,
			Pager:          `NewListPager`,
			ResponseStruct: "FactoriesClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdatafactory())
}
