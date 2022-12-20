// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"

func Armappcomplianceautomation() []*Table {
	tables := []*Table{
		{
			NewFunc:        armappcomplianceautomation.NewReportsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation",
			URL:            "/providers/Microsoft.AppComplianceAutomation/reports",
			Namespace:      "microsoft.appcomplianceautomation",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_appcomplianceautomation)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ReportsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armappcomplianceautomation())
}
