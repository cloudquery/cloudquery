package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
)

func EventHub() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &eventhub.EHNamespace{},
					listFunction:       "List",
					subServiceOverride: "Namespaces",
					relations:          []string{"networkRuleSets()"},
				},
			},
			serviceNameOverride: "EventHub",
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &eventhub.NetworkRuleSet{},
					listFunction: "GetNetworkRuleSet",
					listFunctionArgsInit: []string{`namespace := parent.Item.(eventhub.EHNamespace)
					resource, err := client.ParseResourceID(*namespace.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					listFunctionArgs: []string{"resource.ResourceGroup", "*namespace.Name"},
					listHandler: `if err != nil {
						return errors.WithStack(err)
					}
					res <- response`,
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
					mockListResult:           mockDirectResponse,
				},
			},
			serviceNameOverride: "EventHub",
		},
	}

	return generateResources(resourcesByTemplates)
}
