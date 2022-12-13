// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"

func Armautomation() []*Table {
	tables := []*Table{
		{
			NewFunc:        armautomation.NewAccountClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Automation/automationAccounts",
			Namespace:      "Microsoft.Automation",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Automation)`,
			Pager:          `NewListPager`,
			ResponseStruct: "AccountClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armautomation())
}
