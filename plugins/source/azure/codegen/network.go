package codegen

import (
	"path"
	"reflect"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
)

var resourcesWithListFunction = []Resource{}

var resourcesWithListAllFunction = []Resource{
	{
		AzureStruct:  &network.ExpressRouteCircuit{},
		ListFunction: "ListAll",
	},
}

func NetworkResources() []Resource {
	resources := []Resource{}
	resources = append(resources, resourcesWithListFunction...)
	resources = append(resources, resourcesWithListAllFunction...)
	for i := range resources {
		elementTypeParts := strings.Split(reflect.TypeOf(resources[i].AzureStruct).Elem().String(), ".")
		resources[i].AzurePackageName = elementTypeParts[0]
		resources[i].AzureStructName = elementTypeParts[1]
		resources[i].AzureService = strcase.ToCamel(resources[i].AzurePackageName)
		resources[i].AzureSubService = resources[i].AzureStructName + "s"
		resources[i].DefaultColumns = []codegen.ColumnDefinition{SubscriptionIdColumn, IdColumn}
		resources[i].SkipFields = []string{"ID"}
		resources[i].Imports = []string{}
		resources[i].MockImports = []string{"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"}
		resources[i].CreateTableOptions.PrimaryKeys = []string{"subscription_id", "id"}
		resources[i].Templates = []Template{
			{Source: "resource_list.go.tpl", Destination: path.Join(resources[i].AzurePackageName, strcase.ToSnake(resources[i].AzureStructName)+".go")},
			{Source: "resource_list_mock_test.go.tpl", Destination: path.Join(resources[i].AzurePackageName, strcase.ToSnake(resources[i].AzureStructName)+"_mock_test.go")},
		}
		initResourceTable(&resources[i])
	}

	return resources
}
