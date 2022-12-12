// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"

func Armpowerbidedicated() []*Table {
	tables := []*Table{
		{
			NewFunc:        armpowerbidedicated.NewCapacitiesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.PowerBIDedicated/capacities",
			Namespace:      "Microsoft.PowerBIDedicated",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_PowerBIDedicated)`,
			Pager:          `NewListPager`,
			ResponseStruct: "CapacitiesClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armpowerbidedicated())
}
