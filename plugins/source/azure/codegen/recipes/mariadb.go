package recipes

import "github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb"

func MariaDB() []Resource {
	var serverRelations = []resourceDefinition{
		{
			azureStruct:      &mariadb.Configuration{},
			listFunction:     "ListByServer",
			listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
			listFunctionArgsInit: []string{"server := parent.Item.(mariadb.Server)", `resourceDetails, err := client.ParseResourceID(*server.ID)
			if err != nil {
				return errors.WithStack(err)
			}`},
			listHandler:              valueHandler,
			isRelation:               true,
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
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
					source:            "resource_list_value_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb"},
				},
			},
			definitions: append([]resourceDefinition{
				{
					azureStruct:  &mariadb.Server{},
					listFunction: "List",
					listHandler:  valueHandler,
					relations:    serverRelations,
				},
			}, serverRelations...),
			serviceNameOverride: "MariaDB",
		},
	}

	return generateResources(resourcesByTemplates)
}
