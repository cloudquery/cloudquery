// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"

func init() {
	tables := []Table{
		{
			Service:        "armdesktopvirtualization",
			Name:           "host_pools",
			Struct:         &armdesktopvirtualization.HostPool{},
			ResponseStruct: &armdesktopvirtualization.HostPoolsClientListResponse{},
			Client:         &armdesktopvirtualization.HostPoolsClient{},
			ListFunc:       (&armdesktopvirtualization.HostPoolsClient{}).NewListPager,
			NewFunc:        armdesktopvirtualization.NewHostPoolsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DesktopVirtualization/hostPools",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DesktopVirtualization)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
