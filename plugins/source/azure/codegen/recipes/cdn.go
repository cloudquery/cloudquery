package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
)

func CDN() []Resource {
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
			definitions: []resourceDefinition{
				{
					azureStruct:  &cdn.Profile{},
					listFunction: "List",
					relations:    []string{"endpoints()", "ruleSets()", "securityPolicies()"},
				},
				{
					azureStruct:  &cdn.Endpoint{},
					listFunction: "ListByProfile",
					listFunctionArgsInit: []string{`profile := parent.Item.(cdn.Profile)
					resource, err := client.ParseResourceID(*profile.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					listFunctionArgs:         []string{"resource.ResourceGroup", "*profile.Name"},
					relations:                []string{"customDomains()", "routes()"},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
				{
					azureStruct:  &cdn.RuleSet{},
					listFunction: "ListByProfile",
					listFunctionArgsInit: []string{`profile := parent.Item.(cdn.Profile)
					resource, err := client.ParseResourceID(*profile.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					listFunctionArgs:         []string{"resource.ResourceGroup", "*profile.Name"},
					relations:                []string{"rules()"},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
				{
					azureStruct:  &cdn.SecurityPolicy{},
					listFunction: "ListByProfile",
					listFunctionArgsInit: []string{`profile := parent.Item.(cdn.Profile)
					resource, err := client.ParseResourceID(*profile.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					listFunctionArgs:         []string{"resource.ResourceGroup", "*profile.Name"},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
				{
					azureStruct:  &cdn.CustomDomain{},
					listFunction: "ListByEndpoint",
					listFunctionArgsInit: []string{`profile := parent.Parent.Item.(cdn.Profile)
					resource, err := client.ParseResourceID(*profile.ID)
					if err != nil {
						return errors.WithStack(err)
					}`, `endpoint := parent.Item.(cdn.Endpoint)`},
					listFunctionArgs:         []string{"resource.ResourceGroup", "*profile.Name", "*endpoint.Name"},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
				},
				{
					azureStruct:  &cdn.Route{},
					listFunction: "ListByEndpoint",
					listFunctionArgsInit: []string{`profile := parent.Parent.Item.(cdn.Profile)
					resource, err := client.ParseResourceID(*profile.ID)
					if err != nil {
						return errors.WithStack(err)
					}`, `endpoint := parent.Item.(cdn.Endpoint)`},
					listFunctionArgs:         []string{"resource.ResourceGroup", "*profile.Name", "*endpoint.Name"},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
				},
				{
					azureStruct:  &cdn.Rule{},
					listFunction: "ListByRuleSet",
					listFunctionArgsInit: []string{`profile := parent.Parent.Item.(cdn.Profile)
					resource, err := client.ParseResourceID(*profile.ID)
					if err != nil {
						return errors.WithStack(err)
					}`, `ruleSet := parent.Item.(cdn.RuleSet)`},
					listFunctionArgs:         []string{"resource.ResourceGroup", "*profile.Name", "*ruleSet.Name"},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`, `"test"`},
				},
			},
			serviceNameOverride: "CDN",
		},
	}

	return generateResources(resourcesByTemplates)
}
