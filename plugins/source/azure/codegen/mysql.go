package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
)

func MySQLResources() []Resource {
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
				},
			},
			serviceNameOverride: "MySQL",
		},
	}

	return generateResources(resourcesByTemplates)
}
