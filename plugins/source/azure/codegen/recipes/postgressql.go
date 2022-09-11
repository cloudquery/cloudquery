package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
)

func PostgresSQL() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:  &postgresql.Server{},
					listFunction: "List",
					listHandler:  valueHandler,
					relations:    []string{"configurations()", "firewallRules()"},
				},
				{
					azureStruct:      &postgresql.Configuration{},
					listFunction:     "ListByServer",
					listHandler:      valueHandler,
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
					listFunctionArgsInit: []string{"server := parent.Item.(postgresql.Server)", `resourceDetails, err := client.ParseResourceID(*server.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
				{
					azureStruct:      &postgresql.FirewallRule{},
					listFunction:     "ListByServer",
					listHandler:      valueHandler,
					listFunctionArgs: []string{"resourceDetails.ResourceGroup", "*server.Name"},
					listFunctionArgsInit: []string{"server := parent.Item.(postgresql.Server)", `resourceDetails, err := client.ParseResourceID(*server.ID)
					if err != nil {
						return errors.WithStack(err)
					}`},
					isRelation:               true,
					mockListFunctionArgsInit: []string{""},
					mockListFunctionArgs:     []string{`"test"`, `"test"`},
				},
			},
			serviceNameOverride: "PostgreSQL",
		},
	}

	return generateResources(resourcesByTemplates)
}
