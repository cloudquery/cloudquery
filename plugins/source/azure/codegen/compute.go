package codegen

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
)

func ComputeResources() []Resource {
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:    &compute.Disk{},
					listFunction:   "List",
					mockListResult: "List",
				},
				{
					azureStruct: &compute.VirtualMachineScaleSet{},
				},
				{
					azureStruct:      &compute.VirtualMachine{},
					listFunctionArgs: []string{`"false"`},
				},
			},
		},
	}

	return generateResources(resourcesByTemplates)
}
