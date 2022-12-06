// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder"

func Armvirtualmachineimagebuilder() []Table {
	tables := []Table{
		{
			Name:           "image_template",
			Struct:         &armvirtualmachineimagebuilder.ImageTemplate{},
			ResponseStruct: &armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListResponse{},
			Client:         &armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClient{},
			ListFunc:       (&armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClient{}).NewListPager,
			NewFunc:        armvirtualmachineimagebuilder.NewVirtualMachineImageTemplatesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.VirtualMachineImages/imageTemplates",
		},
	}

	for i := range tables {
		tables[i].Service = "armvirtualmachineimagebuilder"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armvirtualmachineimagebuilder()...)
}
