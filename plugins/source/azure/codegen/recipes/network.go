package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func Network() []Resource {
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
					azureStruct: &network.VirtualNetwork{},
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
					mockListResult: "ExpressRouteGatewayList",
				},
				{
					azureStruct: &network.Watcher{},
					listHandler: valueHandler,
					relations:   []string{"flowLogs()"},
				},
			},
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:      &network.FlowLog{},
					listFunction:     "List",
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*watcher.Name"},
					listFunctionArgsInit: []string{"watcher := parent.Item.(network.Watcher)", `resourceDetails, err := client.ParseResourceID(*watcher.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					isRelation: true,
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
