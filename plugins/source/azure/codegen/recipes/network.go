package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
)

func Network() []Resource {
	var watcherRelations = []resourceDefinition{
		{
			azureStruct:      &network.FlowLog{},
			listFunction:     "List",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*watcher.Name"},
			listFunctionArgsInit: []string{"watcher := parent.Item.(network.Watcher)", `resourceDetails, err := client.ParseResourceID(*watcher.ID)
			if err != nil {
				return err
			}`},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
		},
	}
	var gatewayRelations = []resourceDefinition{
		{
			azureStruct:      &network.VirtualNetworkGatewayConnectionListEntity{},
			listFunction:     "ListConnections",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*gateway.Name"},
			listFunctionArgsInit: []string{"gateway := parent.Item.(network.VirtualNetworkGateway)", `resourceDetails, err := client.ParseResourceID(*gateway.ID)
			if err != nil {
				return err
			}`},
			subServiceOverride:       "VirtualNetworkGatewayConnections",
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
			mockListResult:           "VirtualNetworkGatewayListConnectionsResult",
		},
	}
	var networkRelations = []resourceDefinition{
		{
			azureStruct:      &network.VirtualNetworkGateway{},
			listFunction:     "List",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup"},
			listFunctionArgsInit: []string{"network := parent.Item.(network.VirtualNetwork)", `resourceDetails, err := client.ParseResourceID(*network.ID)
			if err != nil {
				return err
			}`},
			relations:                gatewayRelations,
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`},
		},
	}
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"},
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
					relations:   networkRelations,
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
					relations:   watcherRelations,
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
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"},
				},
			},
		},
	}

	initParents(resourcesByTemplates)

	resourcesByTemplates[0].definitions = append(resourcesByTemplates[0].definitions, append(append(watcherRelations, networkRelations...), gatewayRelations...)...)

	return generateResources(resourcesByTemplates)
}
