// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance"

func init() {
	tables := []Table{
		{
			Service:        "armcontainerinstance",
			Name:           "container_groups",
			Struct:         &armcontainerinstance.ContainerGroup{},
			ResponseStruct: &armcontainerinstance.ContainerGroupsClientListResponse{},
			Client:         &armcontainerinstance.ContainerGroupsClient{},
			ListFunc:       (&armcontainerinstance.ContainerGroupsClient{}).NewListPager,
			NewFunc:        armcontainerinstance.NewContainerGroupsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/containerGroups",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ContainerInstance)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
