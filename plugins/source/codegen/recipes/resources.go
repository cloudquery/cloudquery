package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
)

func Resources() []Resource {
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
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources",
						"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links",
						"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy",
					},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:      &resources.Group{},
					listFunction:     "List",
					listFunctionArgs: []string{`""`, `nil`},
				},
				{
					azureStruct:        &links.ResourceLink{},
					skipFields:         []string{"Type"},
					listFunction:       "ListAtSubscription",
					listFunctionArgs:   []string{`""`},
					subServiceOverride: "Links",
					mockListResult:     "ResourceLinkResult",
				},
				{
					azureStruct:          &policy.Assignment{},
					listFunction:         "List",
					listFunctionArgs:     []string{"meta.(*client.Client).SubscriptionId", `""`, `nil`},
					mockListFunctionArgs: []string{"gomock.Any()", `""`, `nil`},
					subServiceOverride:   "PolicyAssignments",
				},
			},
			serviceNameOverride: "Resources",
		},
	}

	return generateResources(resourcesByTemplates)
}
