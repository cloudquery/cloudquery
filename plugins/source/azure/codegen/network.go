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
					azureStruct:    &network.ExpressRoutePort{},
					templateParams: []string{"List"},
				},
				{
					azureStruct:    &network.RouteFilter{},
					templateParams: []string{"List"},
				},
				{
					azureStruct:    &network.ExpressRouteCircuit{},
					templateParams: []string{"ListAll"},
				},
				{
					azureStruct:    &network.Interface{},
					templateParams: []string{"ListAll"},
				},
				{
					azureStruct:    &network.PublicIPAddress{},
					templateParams: []string{"ListAll"},
				},
				{
					azureStruct:    &network.RouteTable{},
					templateParams: []string{"ListAll"},
				},
				{
					azureStruct:    &network.SecurityGroup{},
					templateParams: []string{"ListAll"},
				},
				{
					azureStruct:    &network.VirtualNetwork{},
					templateParams: []string{"ListAll"},
				},
			},
		},
		{
			templates: []template{
				{
					source:            "resource_list_by_subscription.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{},
				},
				{
					source:            "resource_list_by_subscription_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct: &network.ExpressRouteGateway{},
				},
			},
		},
		{
			templates: []template{
				{
					source:            "resource_list_no_page.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{},
				},
				{
					source:            "resource_list_no_page_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:    &network.Watcher{},
					templateParams: []string{"ListAll"},
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
