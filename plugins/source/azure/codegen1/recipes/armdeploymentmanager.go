// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"

func Armdeploymentmanager() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdeploymentmanager.NewArtifactSourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager",
		},
		{
			NewFunc: armdeploymentmanager.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager",
		},
		{
			NewFunc: armdeploymentmanager.NewRolloutsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager",
		},
		{
			NewFunc: armdeploymentmanager.NewServiceTopologiesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager",
		},
		{
			NewFunc: armdeploymentmanager.NewStepsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager",
		},
		{
			NewFunc: armdeploymentmanager.NewServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager",
		},
		{
			NewFunc: armdeploymentmanager.NewServiceUnitsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdeploymentmanager())
}