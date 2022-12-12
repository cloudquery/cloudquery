// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"

func init() {
	tables := []Table{
		{
			Service:        "armrelay",
			Name:           "namespaces",
			Struct:         &armrelay.Namespace{},
			ResponseStruct: &armrelay.NamespacesClientListResponse{},
			Client:         &armrelay.NamespacesClient{},
			ListFunc:       (&armrelay.NamespacesClient{}).NewListPager,
			NewFunc:        armrelay.NewNamespacesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Relay/namespaces",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Relay)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
