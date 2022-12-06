// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"

func Armdns() []*Table {
	tables := []*Table{
		{
			NewFunc:   armdns.NewZonesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnszones",
			Namespace: "Microsoft.Network",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Network")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdns())
}
