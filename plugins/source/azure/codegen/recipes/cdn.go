package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
)

func CDN() []Resource {
	var ruleSetRelations = []resourceDefinition{
		{
			azureStruct:  &cdn.Rule{},
			listFunction: "ListByRuleSet",
			listFunctionArgsInit: []string{`profile := parent.Parent.Item.(cdn.Profile)
			resource, err := client.ParseResourceID(*profile.ID)
			if err != nil {
				return err
			}`, `ruleSet := parent.Item.(cdn.RuleSet)`},
			listFunctionArgs:         []string{"resource.ResourceGroup", "*profile.Name", "*ruleSet.Name"},
			mockListFunctionArgsInit: []string{`data.Actions = &[]cdn.BasicDeliveryRuleAction{}`, `data.Conditions = &[]cdn.BasicDeliveryRuleCondition{}`},
			mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
		},
	}
	var endpointRelations = []resourceDefinition{
		{
			azureStruct:  &cdn.CustomDomain{},
			listFunction: "ListByEndpoint",
			listFunctionArgsInit: []string{`profile := parent.Parent.Item.(cdn.Profile)
			resource, err := client.ParseResourceID(*profile.ID)
			if err != nil {
				return err
			}`, `endpoint := parent.Item.(cdn.Endpoint)`},
			listFunctionArgs:         []string{"resource.ResourceGroup", "*profile.Name", "*endpoint.Name"},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
		},
		{
			azureStruct:  &cdn.Route{},
			listFunction: "ListByEndpoint",
			listFunctionArgsInit: []string{`profile := parent.Parent.Item.(cdn.Profile)
			resource, err := client.ParseResourceID(*profile.ID)
			if err != nil {
				return err
			}`, `endpoint := parent.Item.(cdn.Endpoint)`},
			listFunctionArgs:         []string{"resource.ResourceGroup", "*profile.Name", "*endpoint.Name"},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
		},
	}
	var profileRelations = []resourceDefinition{
		{
			azureStruct:  &cdn.Endpoint{},
			listFunction: "ListByProfile",
			listFunctionArgsInit: []string{`profile := parent.Item.(cdn.Profile)
			resource, err := client.ParseResourceID(*profile.ID)
			if err != nil {
				return err
			}`},
			listFunctionArgs:         []string{"resource.ResourceGroup", "*profile.Name"},
			relations:                endpointRelations,
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
		},
		{
			azureStruct:  &cdn.RuleSet{},
			listFunction: "ListByProfile",
			listFunctionArgsInit: []string{`profile := parent.Item.(cdn.Profile)
			resource, err := client.ParseResourceID(*profile.ID)
			if err != nil {
				return err
			}`},
			listFunctionArgs:         []string{"resource.ResourceGroup", "*profile.Name"},
			relations:                ruleSetRelations,
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
		},
		{
			azureStruct:  &cdn.SecurityPolicy{},
			listFunction: "ListByProfile",
			listFunctionArgsInit: []string{`profile := parent.Item.(cdn.Profile)
			resource, err := client.ParseResourceID(*profile.ID)
			if err != nil {
				return err
			}`},
			listFunctionArgs:         []string{"resource.ResourceGroup", "*profile.Name"},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
		},
	}

	var topLevelResources = []resourceDefinition{
		{
			azureStruct:  &cdn.Profile{},
			listFunction: "List",
			relations:    profileRelations,
		},
	}

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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"},
				},
			},
			definitions:         topLevelResources,
			serviceNameOverride: "CDN",
		},
	}

	initParents(resourcesByTemplates)

	resourcesByTemplates[0].definitions = append(resourcesByTemplates[0].definitions, profileRelations...)
	resourcesByTemplates[0].definitions = append(resourcesByTemplates[0].definitions, endpointRelations...)
	resourcesByTemplates[0].definitions = append(resourcesByTemplates[0].definitions, ruleSetRelations...)

	return generateResources(resourcesByTemplates)
}
