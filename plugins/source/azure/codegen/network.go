package codegen

import (
	"reflect"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/plugin-sdk/codegen"
)

var resourcesWithListFunction = []Resource{}

var resourcesWithListAllFunction = []Resource{
	{
		AzureStruct:     &network.ExpressRouteCircuit{},
		AzureSubService: "express_route_circuits",
		ListFunction:    "ListAll",
	},
}

func NetworkResources() []Resource {
	resources := []Resource{}
	resources = append(resources, resourcesWithListFunction...)
	resources = append(resources, resourcesWithListAllFunction...)
	for i := range resources {
		resources[i].AzureService = "network"
		resources[i].DefaultColumns = []codegen.ColumnDefinition{SubscriptionIdColumn}
		resources[i].AzureStructName = reflect.TypeOf(resources[i].AzureStruct).Elem().Name()
		resources[i].SkipFields = []string{}
		resources[i].Imports = []string{}
		resources[i].CreateTableOptions.PrimaryKeys = []string{"subscription_id", "id"}
		resources[i].Templates = []string{"resource_list"}
		initResourceTable(&resources[i])
	}

	return resources
}
