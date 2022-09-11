package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
)

func ServiceBus() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &servicebus.SBNamespace{},
					listFunction:       "List",
					subServiceOverride: "Namespaces",
					relations:          []string{"topics()"},
				},
				{
					azureStruct:        &servicebus.SBTopic{},
					listFunction:       "ListByNamespace",
					subServiceOverride: "Topics",
					listFunctionArgs:   []string{"resourceDetails.ResourceGroup", "*namespace.Name", "nil", "nil"},
					listFunctionArgsInit: []string{"namespace := parent.Item.(servicebus.SBNamespace)", `resourceDetails, err := client.ParseResourceID(*namespace.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					relations:                []string{"authorizationRules()"},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`, `nil`, `nil`},
				},
				{
					azureStruct:        &servicebus.SBAuthorizationRule{},
					listFunction:       "ListAuthorizationRules",
					subServiceOverride: "AuthorizationRules",
					listFunctionArgs:   []string{"resourceDetails.ResourceGroup", "*namespace.Name", "*topic.Name"},
					listFunctionArgsInit: []string{
						"namespace := parent.Parent.Item.(servicebus.SBNamespace)",
						"topic := parent.Item.(servicebus.SBTopic)",
						`resourceDetails, err := client.ParseResourceID(*topic.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					relations:                []string{"accessKeys()"},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
				},
				{
					azureStruct:        &servicebus.AccessKeys{},
					listFunction:       "ListKeys",
					subServiceOverride: "AccessKeys",
					listFunctionArgs:   []string{"resourceDetails.ResourceGroup", "*namespace.Name", "*topic.Name", "*rule.Name"},
					listFunctionArgsInit: []string{
						"namespace := parent.Parent.Parent.Item.(servicebus.SBNamespace)",
						"topic := parent.Parent.Item.(servicebus.SBTopic)",
						"rule := parent.Item.(servicebus.SBAuthorizationRule)",
						`resourceDetails, err := client.ParseResourceID(*rule.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					listHandler: `if err != nil {
						return errors.WithStack(err)
					}
					res <- response`,
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`, `"test"`},
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
