// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"

func Armdataboxedge() []*Table {
	tables := []*Table{
		{
			NewFunc:   armdataboxedge.NewAvailableSKUsClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.DataBoxEdge/availableSkus",
			Namespace: "Microsoft.DataBoxEdge",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.DataBoxEdge")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdataboxedge())
}
