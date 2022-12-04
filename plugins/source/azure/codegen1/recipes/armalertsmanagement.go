// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"

func Armalertsmanagement() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armalertsmanagement.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement",
		},
		{
			NewFunc: armalertsmanagement.NewAlertsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement",
		},
		{
			NewFunc: armalertsmanagement.NewSmartGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement",
		},
		{
			NewFunc: armalertsmanagement.NewAlertProcessingRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armalertsmanagement())
}