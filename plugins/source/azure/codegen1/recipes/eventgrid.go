// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"

func Armeventgrid() []*Table {
	tables := []*Table{
		{
			NewFunc:   armeventgrid.NewTopicTypesClient,
			PkgPath:   "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid",
			URL:       "/providers/Microsoft.EventGrid/topicTypes",
			Namespace: "Microsoft.EventGrid",
			Multiplex: `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.EventGrid")`,
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armeventgrid())
}
