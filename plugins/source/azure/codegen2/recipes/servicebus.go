// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"

func Armservicebus() []Table {
	tables := []Table{
		{
			Service:        "armservicebus",
			Name:           "namespaces",
			Struct:         &armservicebus.SBNamespace{},
			ResponseStruct: &armservicebus.NamespacesClientListResponse{},
			Client:         &armservicebus.NamespacesClient{},
			ListFunc:       (&armservicebus.NamespacesClient{}).NewListPager,
			NewFunc:        armservicebus.NewNamespacesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/namespaces",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ServiceBus)`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, Armservicebus()...)
}
