// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"

func Armdigitaltwins() []Table {
	tables := []Table{
		{
      Name: "endpoint_resource",
      Struct: &armdigitaltwins.EndpointResource{},
      ResponseStruct: &armdigitaltwins.EndpointClientListResponse{},
      Client: &armdigitaltwins.EndpointClient{},
      ListFunc: (&armdigitaltwins.EndpointClient{}).NewListPager,
			NewFunc: armdigitaltwins.NewEndpointClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints",
		},
		{
      Name: "time_series_database_connection",
      Struct: &armdigitaltwins.TimeSeriesDatabaseConnection{},
      ResponseStruct: &armdigitaltwins.TimeSeriesDatabaseConnectionsClientListResponse{},
      Client: &armdigitaltwins.TimeSeriesDatabaseConnectionsClient{},
      ListFunc: (&armdigitaltwins.TimeSeriesDatabaseConnectionsClient{}).NewListPager,
			NewFunc: armdigitaltwins.NewTimeSeriesDatabaseConnectionsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/timeSeriesDatabaseConnections",
		},
	}

	for i := range tables {
		tables[i].Service = "armdigitaltwins"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armdigitaltwins()...)
}