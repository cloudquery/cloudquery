package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	keyvault71 "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	hsmKeyValue "github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
)

func KeyValue() []Resource {
	var vaultRelations = []resourceDefinition{
		{
			azureStruct:              &keyvault71.KeyItem{},
			listFunction:             "GetKeys",
			listFunctionArgs:         []string{"*vault.Properties.VaultURI", "&maxResults"},
			listFunctionArgsInit:     []string{"vault := parent.Item.(keyvault.Vault)", "maxResults := int32(25)"},
			subServiceOverride:       "Keys",
			mockListFunctionArgsInit: []string{"maxResults := int32(25)"},
			mockListFunctionArgs:     []string{`"test"`, `&maxResults`},
			mockListResult:           "KeyListResult",
		},
		{
			azureStruct:              &keyvault71.SecretItem{},
			listFunction:             "GetSecrets",
			listFunctionArgs:         []string{"*vault.Properties.VaultURI", "&maxResults"},
			listFunctionArgsInit:     []string{"vault := parent.Item.(keyvault.Vault)", "maxResults := int32(25)"},
			subServiceOverride:       "Secrets",
			mockListFunctionArgsInit: []string{"maxResults := int32(25)"},
			mockListFunctionArgs:     []string{`"test"`, `&maxResults`},
			mockListResult:           "SecretListResult",
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:          &hsmKeyValue.ManagedHsm{},
					listFunction:         "ListBySubscription",
					listFunctionArgs:     []string{"&maxResults"},
					listFunctionArgsInit: []string{"maxResults := int32(100)"},
				},
			},
			serviceNameOverride: "KeyVault",
		},
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:              &keyvault.Vault{},
					listFunction:             "ListBySubscription",
					listFunctionArgs:         []string{"&maxResults"},
					listFunctionArgsInit:     []string{"maxResults := int32(1000)"},
					relations:                vaultRelations,
					mockListFunctionArgsInit: []string{"maxResults := int32(1000)", `vaultURI := "test"`, `data.Properties.VaultURI = &vaultURI`},
				},
			},
			serviceNameOverride: "KeyVault",
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"},
				},
				{
					source:            "resource_list_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"},
				},
			},
			definitions:         vaultRelations,
			serviceNameOverride: "KeyVault",
		},
	}

	initParents(resourcesByTemplates)
	return generateResources(resourcesByTemplates)
}
