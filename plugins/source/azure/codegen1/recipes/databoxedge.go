// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"

func Armdataboxedge() []*Table {
	tables := []*Table{
		{
			NewFunc:        armdataboxedge.NewAvailableSKUsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DataBoxEdge/availableSkus",
			Namespace:      "microsoft.databoxedge",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_databoxedge)`,
			Pager:          `NewListPager`,
			ResponseStruct: "AvailableSKUsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdataboxedge())
}
