// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"

func Armdesktopvirtualization() []Table {
	tables := []Table{
		{
			Name:           "host_pool",
			Struct:         &armdesktopvirtualization.HostPool{},
			ResponseStruct: &armdesktopvirtualization.HostPoolsClientListResponse{},
			Client:         &armdesktopvirtualization.HostPoolsClient{},
			ListFunc:       (&armdesktopvirtualization.HostPoolsClient{}).NewListPager,
			NewFunc:        armdesktopvirtualization.NewHostPoolsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DesktopVirtualization/hostPools",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.DesktopVirtualization")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armdesktopvirtualization"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdesktopvirtualization()...)
}
