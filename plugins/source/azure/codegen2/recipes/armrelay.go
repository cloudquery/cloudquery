// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"

func Armrelay() []Table {
	tables := []Table{
		{
      Name: "private_endpoint_connection",
      Struct: &armrelay.PrivateEndpointConnection{},
      ResponseStruct: &armrelay.PrivateEndpointConnectionsClientListResponse{},
      Client: &armrelay.PrivateEndpointConnectionsClient{},
      ListFunc: (&armrelay.PrivateEndpointConnectionsClient{}).NewListPager,
			NewFunc: armrelay.NewPrivateEndpointConnectionsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/privateEndpointConnections",
		},
		{
      Name: "namespace",
      Struct: &armrelay.Namespace{},
      ResponseStruct: &armrelay.NamespacesClientListResponse{},
      Client: &armrelay.NamespacesClient{},
      ListFunc: (&armrelay.NamespacesClient{}).NewListPager,
			NewFunc: armrelay.NewNamespacesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Relay/namespaces",
		},
	}

	for i := range tables {
		tables[i].Service = "armrelay"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armrelay()...)
}