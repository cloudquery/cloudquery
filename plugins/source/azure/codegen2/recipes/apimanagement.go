// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"

func Armapimanagement() []Table {
	tables := []Table{
		{
			Name:           "service",
			Struct:         &armapimanagement.ServiceResource{},
			ResponseStruct: &armapimanagement.ServiceClientListResponse{},
			Client:         &armapimanagement.ServiceClient{},
			ListFunc:       (&armapimanagement.ServiceClient{}).NewListPager,
			NewFunc:        armapimanagement.NewServiceClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ApiManagement/service",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ApiManagement)`,
		},
	}

	for i := range tables {
		tables[i].Service = "armapimanagement"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armapimanagement()...)
}
