package recipes

import (
	analytics "github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/analytics/mgmt/account"
	store "github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/store/mgmt/account"
)

func Datalake() []Resource {
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/analytics/mgmt/account",
					},
				},
				{
					source:            "resource_list_value_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/analytics/mgmt/account",
					},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:      &analytics.DataLakeAnalyticsAccount{},
					listFunction:     "List",
					listFunctionArgs: []string{`""`, `nil`, `nil`, `""`, `""`, `nil`},
					getFunction:      "Get",
					getFunctionArgsInit: []string{"account := r.Item.(account.DataLakeAnalyticsAccountBasic)", "resourceDetails, err := client.ParseResourceID(*account.ID)", `if err != nil {
						errors.WithStack(err)
					}`},
					getFunctionArgs:    []string{"resourceDetails.ResourceGroup", "*account.Name"},
					mockValueType:      "DataLakeAnalyticsAccountBasic",
					mockDefinitionType: `DataLakeAnalyticsAccountBasic`,
				},
			},
			serviceNameOverride: "DataLake",
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/store/mgmt/account",
					},
				},
				{
					source:            "resource_list_value_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/store/mgmt/account",
					},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:      &store.DataLakeStoreAccount{},
					listFunction:     "List",
					listFunctionArgs: []string{`""`, `nil`, `nil`, `""`, `""`, `nil`},
					getFunction:      "Get",
					getFunctionArgsInit: []string{"account := r.Item.(account.DataLakeStoreAccount)", "resourceDetails, err := client.ParseResourceID(*account.ID)", `if err != nil {
						errors.WithStack(err)
					}`},
					getFunctionArgs:    []string{"resourceDetails.ResourceGroup", "*account.Name"},
					mockValueType:      "DataLakeStoreAccountBasic",
					mockDefinitionType: "DataLakeStoreAccountBasic",
				},
			},
			serviceNameOverride: "DataLake",
		},
	}

	return generateResources(resourcesByTemplates)
}
