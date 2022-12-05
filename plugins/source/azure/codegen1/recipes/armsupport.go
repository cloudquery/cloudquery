// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"

func Armsupport() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armsupport.NewCommunicationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}/communications",
		},
		{
			NewFunc: armsupport.NewProblemClassificationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport",
			URL: "/providers/Microsoft.Support/services/{serviceName}/problemClassifications",
		},
		{
			NewFunc: armsupport.NewServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport",
			URL: "/providers/Microsoft.Support/services",
		},
		{
			NewFunc: armsupport.NewTicketsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armsupport())
}