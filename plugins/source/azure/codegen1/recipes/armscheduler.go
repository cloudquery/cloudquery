// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scheduler/armscheduler"

func Armscheduler() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armscheduler.NewJobCollectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scheduler/armscheduler",
			URL: "",
		},
		{
			NewFunc: armscheduler.NewJobsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scheduler/armscheduler",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armscheduler())
}