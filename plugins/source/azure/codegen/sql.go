package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func SQL() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &sql.Server{},
					listFunction: "List",
				},
				{
					azureStruct:  &sql.ManagedInstance{},
					listFunction: "List",
				},
			},
			serviceNameOverride: "SQL",
		},
	}

	return generateResources(resourcesByTemplates)
}
