package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func NetworkResources() []Resource {
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{},
				},
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &network.ExpressRoutePort{},
					listFunction: "List",
				},
				{
					azureStruct:  &network.RouteFilter{},
					listFunction: "List",
				},
				{
					azureStruct: &network.ExpressRouteCircuit{},
				},
				{
					azureStruct: &network.Interface{},
				},
				{
					azureStruct: &network.PublicIPAddress{},
				},
				{
					azureStruct: &network.RouteTable{},
				},
				{
					azureStruct: &network.SecurityGroup{},
				},
			},
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{},
				},
				{
					source:            "resource_list_value_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:    &network.ExpressRouteGateway{},
					listFunction:   "ListBySubscription",
					listHandler:    valueHandler,
					mockListResult: "List",
				},
				{
					azureStruct: &network.Watcher{},
					listHandler: valueHandler,
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
