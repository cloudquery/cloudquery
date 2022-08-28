package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
)

func Security() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:          &security.Assessment{},
					listFunction:         "List",
					listFunctionArgs:     []string{`fmt.Sprintf("/subscriptions/%s", meta.(*client.Client).SubscriptionId)`},
					mockListFunctionArgs: []string{`"/subscriptions/test_sub"`},
					mockListResult:       "AssessmentList",
					mockFieldsToIgnore:   []string{"ResourceDetails"},
				},
				{
					azureStruct:    &security.AutoProvisioningSetting{},
					listFunction:   "List",
					mockListResult: "AutoProvisioningSettingList",
				},
				{
					azureStruct:    &security.Contact{},
					listFunction:   "List",
					mockListResult: "ContactList",
				},
				{
					azureStruct:    &security.JitNetworkAccessPolicy{},
					listFunction:   "List",
					mockListResult: "JitNetworkAccessPoliciesList",
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:    &security.Pricing{},
					listFunction:   "List",
					listHandler:    valueHandler,
					mockListResult: "PricingList",
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
