package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

func MariaDBResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "servers",
			Struct: &armmariadb.Server{},
			ResponseStruct: &armmariadb.ServersClientListResponse{},
			Client: &armmariadb.ServersClient{},
			ListFunc: (&armmariadb.ServersClient{}).NewListPager,
			NewFunc: armmariadb.NewServersClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "mariadb/armmariadb"
		r.Service = "armmariadb"
		r.Template = "list"
	}

	return resources
}