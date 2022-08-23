package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
)

func AuthorizationResources() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:          &authorization.RoleAssignment{},
					listFunction:         "List",
					listFunctionArgs:     []string{"\"\""},
					mockListFunctionArgs: []string{"\"\""},
				},
				{
					azureStruct:          &authorization.RoleDefinition{},
					listFunction:         "List",
					listFunctionArgs:     []string{"client.ScopeSubscription(meta.(*client.Client).SubscriptionId)", "\"\""},
					mockListFunctionArgs: []string{"gomock.Any()", "\"\""},
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
