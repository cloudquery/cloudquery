// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"

func Armprivatedns() []*Table {
	tables := []*Table{
		{
			NewFunc:   armprivatedns.NewPrivateZonesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns",
			URL:       "/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateDnsZones",
			Namespace: "Microsoft.Network",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.Network")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armprivatedns())
}
