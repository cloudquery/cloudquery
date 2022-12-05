// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"

func Armdatabox() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdatabox.NewJobsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DataBox/jobs",
		},
		{
			NewFunc: armdatabox.NewManagementClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox",
			URL: "",
		},
		{
			NewFunc: armdatabox.NewServiceClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox",
			URL: "",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdatabox())
}