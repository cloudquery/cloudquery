// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"

func Armguestconfiguration() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armguestconfiguration.NewHCRPAssignmentReportsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration",
		},
		{
			NewFunc: armguestconfiguration.NewAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration",
		},
		{
			NewFunc: armguestconfiguration.NewAssignmentsVMSSClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration",
		},
		{
			NewFunc: armguestconfiguration.NewHCRPAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration",
		},
		{
			NewFunc: armguestconfiguration.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration",
		},
		{
			NewFunc: armguestconfiguration.NewAssignmentReportsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration",
		},
		{
			NewFunc: armguestconfiguration.NewAssignmentReportsVMSSClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armguestconfiguration())
}