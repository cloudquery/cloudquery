// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"

func Armsupport() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armsupport.NewServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport",
		},
		{
			NewFunc: armsupport.NewTicketsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport",
		},
		{
			NewFunc: armsupport.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport",
		},
		{
			NewFunc: armsupport.NewCommunicationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport",
		},
		{
			NewFunc: armsupport.NewProblemClassificationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armsupport())
}