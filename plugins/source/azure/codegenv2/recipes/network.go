package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

func NetworkResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "virtual_networks",
			Struct: &armnetwork.VirtualNetwork{},
			ResponseStruct: &armnetwork.VirtualNetworksClientListAllResponse{},
			Client: &armnetwork.VirtualNetworksClient{},
			ListFunc: (&armnetwork.VirtualNetworksClient{}).NewListAllPager,
			NewFunc: armnetwork.NewVirtualNetworksClient,
			OutputField: "Value",
		},
		{
			SubService: "security_groups",
			Struct: &armnetwork.SecurityGroup{},
			ResponseStruct: &armnetwork.SecurityGroupsClientListAllResponse{},
			Client: &armnetwork.SecurityGroupsClient{},
			ListFunc: (&armnetwork.SecurityGroupsClient{}).NewListAllPager,
			NewFunc: armnetwork.NewSecurityGroupsClient,
			OutputField: "Value",
		},
		{
			SubService: "interfaces",
			Struct: &armnetwork.Interface{},
			ResponseStruct: &armnetwork.InterfacesClientListAllResponse{},
			Client: &armnetwork.InterfacesClient{},
			ListFunc: (&armnetwork.InterfacesClient{}).NewListAllPager,
			NewFunc: armnetwork.NewInterfacesClient,
			OutputField: "Value",
		},
		{
			SubService: "watchers",
			Struct: &armnetwork.Watcher{},
			ResponseStruct: &armnetwork.WatchersClientListAllResponse{},
			Client: &armnetwork.WatchersClient{},
			ListFunc: (&armnetwork.WatchersClient{}).NewListAllPager,
			NewFunc: armnetwork.NewWatchersClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "network/armnetwork"
		r.Service = "armnetwork"
		r.Template = "list"
	}

	return resources
}