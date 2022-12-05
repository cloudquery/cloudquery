// Code generated by codegen; DO NOT EDIT.
package packages

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"

func Armbotservice() []*Table {
	tables := []*Table{
		{
			NewFunc: armbotservice.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/privateEndpointConnections",
		},
		{
			NewFunc: armbotservice.NewBotsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.BotService/botServices",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armbotservice())
}