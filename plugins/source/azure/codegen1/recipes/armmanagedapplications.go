// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications"

func Armmanagedapplications() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmanagedapplications.NewApplicationDefinitionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications",
		},
		{
			NewFunc: armmanagedapplications.NewApplicationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications",
		},
		{
			NewFunc: armmanagedapplications.NewJitRequestsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications",
		},
		{
			NewFunc: armmanagedapplications.NewApplicationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmanagedapplications())
}