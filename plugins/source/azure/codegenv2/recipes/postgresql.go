package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

func PostgreSQLResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "servers",
			Struct: &armpostgresql.Server{},
			ResponseStruct: &armpostgresql.ServersClientListResponse{},
			Client: &armpostgresql.ServersClient{},
			ListFunc: (&armpostgresql.ServersClient{}).NewListPager,
			NewFunc: armpostgresql.NewServersClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "postgresql/armpostgresql"
		r.Service = "armpostgresql"
		r.Template = "list"
	}

	return resources
}