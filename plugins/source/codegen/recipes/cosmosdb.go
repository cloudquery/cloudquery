package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
)

func CosmosDB() []Resource {
	var accountRelations = []resourceDefinition{
		{
			azureStruct:        &documentdb.MongoDBDatabaseGetResults{},
			listFunction:       "ListMongoDBDatabases",
			listHandler:        valueHandler,
			subServiceOverride: "MongoDBDatabases",
			listFunctionArgsInit: []string{`account := parent.Item.(documentdb.DatabaseAccountGetResults)
			resource, err := client.ParseResourceID(*account.ID)
			if err != nil {
				return err
			}`},
			listFunctionArgs:         []string{"resource.ResourceGroup", "*account.Name"},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
			mockListResult:           "MongoDBDatabaseListResult",
		},
		{
			azureStruct:        &documentdb.SQLDatabaseGetResults{},
			listFunction:       "ListSQLDatabases",
			listHandler:        valueHandler,
			subServiceOverride: "SQLDatabases",
			listFunctionArgsInit: []string{`account := parent.Item.(documentdb.DatabaseAccountGetResults)
			resource, err := client.ParseResourceID(*account.ID)
			if err != nil {
				return err
			}`},
			listFunctionArgs:         []string{"resource.ResourceGroup", "*account.Name"},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
			mockListResult:           "SQLDatabaseListResult",
		},
	}
	var resourcesByTemplates = []byTemplates{
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports: []string{
						"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb",
					},
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
					relations:          accountRelations,
				},
			},
			serviceNameOverride: "CosmosDB",
		},
	}

	initParents(resourcesByTemplates)

	resourcesByTemplates[0].definitions = append(resourcesByTemplates[0].definitions, accountRelations...)

	return generateResources(resourcesByTemplates)
}
