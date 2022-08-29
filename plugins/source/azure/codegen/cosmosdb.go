package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
)

func CosmosDB() []Resource {
	var resourcesByTemplates = []byTemplates{
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
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb",
					},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:        &documentdb.DatabaseAccountGetResults{},
					listFunction:       "List",
					listHandler:        valueHandler,
					subServiceOverride: "Accounts",
					mockListResult:     "DatabaseAccountsListResult",
				},
			},
			serviceNameOverride: "CosmosDB",
		},
	}

	return generateResources(resourcesByTemplates)
}
