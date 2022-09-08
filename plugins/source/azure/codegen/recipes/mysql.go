package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
)

func MySQL() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &mysql.Server{},
					listFunction: "List",
					listHandler:  valueHandler,
					relations:    []string{"configurations()"},
				},
			},
			serviceNameOverride: "MySQL",
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:      &mysql.Configuration{},
					listFunction:     "ListByServer",
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
					listFunctionArgsInit: []string{"server := parent.Item.(mysql.Server)", `resourceDetails, err := client.ParseResourceID(*server.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					listHandler: valueHandler,
					isRelation:  true,
				},
			},
			serviceNameOverride: "MySQL",
		},
	}

	return generateResources(resourcesByTemplates)
}
