package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

func MySQLResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "servers",
			Struct: &armmysql.Server{},
			ResponseStruct: &armmysql.ServersClientListResponse{},
			Client: &armmysql.ServersClient{},
			ListFunc: (&armmysql.ServersClient{}).NewListPager,
			NewFunc: armmysql.NewServersClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "mysql/armmysql"
		r.Service = "armmysql"
		r.Template = "list"
	}

	return resources
}