package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Authorization() []Resource {
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
					azureStruct:      &authorization.RoleAssignment{},
					listFunction:     "List",
					listFunctionArgs: []string{`""`},
				},
				{
					azureStruct:          &authorization.RoleDefinition{},
					listFunction:         "List",
					listFunctionArgs:     []string{"client.ScopeSubscription(meta.(*client.Client).SubscriptionId)", `""`},
					mockListFunctionArgs: []string{"gomock.Any()", `""`},
					skipFields:           []string{"RoleType"},
					customColumns:        []codegen.ColumnDefinition{{Name: "role_type", Type: schema.TypeString, Resolver: `schema.PathResolver("RoleType")`}},
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
