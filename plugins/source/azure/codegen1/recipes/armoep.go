// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oep/armoep"

func Armoep() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armoep.NewEnergyServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oep/armoep",
		},
		{
			NewFunc: armoep.NewLocationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oep/armoep",
		},
		{
			NewFunc: armoep.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oep/armoep",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armoep())
}