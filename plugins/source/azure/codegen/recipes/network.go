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
					relations:   []string{"virtualNetworkGateways()"},
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
				{
					azureStruct:      &network.VirtualNetworkGateway{},
					listFunction:     "List",
					listFunctionArgs: []string{"resourceDetails.ResourceGroup"},
					listFunctionArgsInit: []string{"network := parent.Item.(network.VirtualNetwork)", `resourceDetails, err := client.ParseResourceID(*network.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					relations:  []string{"virtualNetworkGatewayConnections()"},
					isRelation: true,
				},
				{
					azureStruct:      &network.VirtualNetworkGatewayConnection{},
					listFunction:     "ListConnections",
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*gateway.Name"},
					listFunctionArgsInit: []string{"gateway := parent.Item.(network.VirtualNetworkGateway)", `resourceDetails, err := client.ParseResourceID(*gateway.ID)
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
