// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs"

func Armnotificationhubs() []*Table {
	tables := []*Table{
		{
			NewFunc:        armnotificationhubs.NewNamespacesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.NotificationHubs/namespaces",
			Namespace:      "Microsoft.NotificationHubs",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_NotificationHubs)`,
			Pager:          `NewListAllPager`,
			ResponseStruct: "NamespacesClientListAllResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armnotificationhubs())
}
