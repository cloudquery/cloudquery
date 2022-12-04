// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"

func Armappcomplianceautomation() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armappcomplianceautomation.NewReportClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation",
		},
		{
			NewFunc: armappcomplianceautomation.NewReportsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation",
		},
		{
			NewFunc: armappcomplianceautomation.NewSnapshotsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation",
		},
		{
			NewFunc: armappcomplianceautomation.NewSnapshotClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation",
		},
		{
			NewFunc: armappcomplianceautomation.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armappcomplianceautomation())
}