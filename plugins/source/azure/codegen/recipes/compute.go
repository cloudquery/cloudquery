package recipes

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Compute() []Resource {
	var virtualMachineRelations = []resourceDefinition{
		{
			azureStruct:  &compute.VirtualMachineInstanceView{},
			listFunction: "InstanceView",
			listFunctionArgsInit: []string{`virtualMachine := parent.Item.(compute.VirtualMachine)
			resource, err := client.ParseResourceID(*virtualMachine.ID)
			if err != nil {
				return err
			}`},
			listFunctionArgs: []string{"resource.ResourceGroup", "*virtualMachine.Name"},
			listHandler: `if err != nil {
				return err
			}
			res <- response`,
			subServiceOverride:       "InstanceViews",
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`},
			mockListResult:           mockDirectResponse,
		},
		{
			azureStruct:  &compute.VirtualMachineExtension{},
			listFunction: "List",
			listFunctionArgsInit: []string{`virtualMachine := parent.Item.(compute.VirtualMachine)
			resource, err := client.ParseResourceID(*virtualMachine.ID)
			if err != nil {
				return err
			}`},
			listFunctionArgs:         []string{"resource.ResourceGroup", "*virtualMachine.Name", `""`},
			listHandler:              valueHandler,
			skipFields:               []string{"Type"},
			customColumns:            []codegen.ColumnDefinition{{Name: "type", Type: schema.TypeString, Resolver: "schema.PathResolver(`Type`)"}},
			mockListFunctionArgsInit: []string{""},
			mockListFunctionArgs:     []string{`"test"`, `"test"`, `""`},
			mockListResult:           "VirtualMachineExtensionsListResult",
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
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"},
				},
			},
			definitions: []resourceDefinition{
				{
					azureStruct:    &compute.Disk{},
					listFunction:   "List",
					mockListResult: "DiskList",
				},
				{
					azureStruct:    &compute.VirtualMachineScaleSet{},
					mockListResult: "VirtualMachineScaleSetListWithLinkResult",
					skipFields:     []string{"DoNotRunExtensionsOnOverprovisionedVMs"},
					customColumns:  []codegen.ColumnDefinition{{Name: "do_not_run_extensions_on_overprovisioned_vms", Type: schema.TypeBool, Resolver: "schema.PathResolver(`DoNotRunExtensionsOnOverprovisionedVMs`)"}},
				},
				{
					azureStruct:      &compute.VirtualMachine{},
					listFunctionArgs: []string{`"false"`},
					relations:        virtualMachineRelations,
				},
			},
		},
		{
			templates: []template{
				{
					source:            "resource_list.go.tpl",
					destinationSuffix: ".go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"},
				},
				{
					source:            "resource_list_value_mock_test.go.tpl",
					destinationSuffix: "_mock_test.go",
					imports:           []string{"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"},
				},
			},
			definitions: virtualMachineRelations,
		},
	}

	initParents(resourcesByTemplates)
	return generateResources(resourcesByTemplates)
}
