// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"

func Armwebpubsub() []Table {
	tables := []Table{
		{
      Name: "hub",
      Struct: &armwebpubsub.Hub{},
      ResponseStruct: &armwebpubsub.HubsClientListResponse{},
      Client: &armwebpubsub.HubsClient{},
      ListFunc: (&armwebpubsub.HubsClient{}).NewListPager,
			NewFunc: armwebpubsub.NewHubsClient,
		},
		{
      Name: "operation",
      Struct: &armwebpubsub.Operation{},
      ResponseStruct: &armwebpubsub.OperationsClientListResponse{},
      Client: &armwebpubsub.OperationsClient{},
      ListFunc: (&armwebpubsub.OperationsClient{}).NewListPager,
			NewFunc: armwebpubsub.NewOperationsClient,
		},
		{
      Name: "private_endpoint_connection",
      Struct: &armwebpubsub.PrivateEndpointConnection{},
      ResponseStruct: &armwebpubsub.PrivateEndpointConnectionsClientListResponse{},
      Client: &armwebpubsub.PrivateEndpointConnectionsClient{},
      ListFunc: (&armwebpubsub.PrivateEndpointConnectionsClient{}).NewListPager,
			NewFunc: armwebpubsub.NewPrivateEndpointConnectionsClient,
		},
		{
      Name: "private_link_resource",
      Struct: &armwebpubsub.PrivateLinkResource{},
      ResponseStruct: &armwebpubsub.PrivateLinkResourcesClientListResponse{},
      Client: &armwebpubsub.PrivateLinkResourcesClient{},
      ListFunc: (&armwebpubsub.PrivateLinkResourcesClient{}).NewListPager,
			NewFunc: armwebpubsub.NewPrivateLinkResourcesClient,
		},
		{
      Name: "shared_private_link_resource",
      Struct: &armwebpubsub.SharedPrivateLinkResource{},
      ResponseStruct: &armwebpubsub.SharedPrivateLinkResourcesClientListResponse{},
      Client: &armwebpubsub.SharedPrivateLinkResourcesClient{},
      ListFunc: (&armwebpubsub.SharedPrivateLinkResourcesClient{}).NewListPager,
			NewFunc: armwebpubsub.NewSharedPrivateLinkResourcesClient,
		},
		{
      Name: "signal_r_service_usage",
      Struct: &armwebpubsub.SignalRServiceUsage{},
      ResponseStruct: &armwebpubsub.UsagesClientListResponse{},
      Client: &armwebpubsub.UsagesClient{},
      ListFunc: (&armwebpubsub.UsagesClient{}).NewListPager,
			NewFunc: armwebpubsub.NewUsagesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armwebpubsub"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armwebpubsub()...)
}