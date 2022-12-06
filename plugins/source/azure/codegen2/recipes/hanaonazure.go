// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"

func Armhanaonazure() []Table {
	tables := []Table{
		{
      Name: "sap_monitor",
      Struct: &armhanaonazure.SapMonitor{},
      ResponseStruct: &armhanaonazure.SapMonitorsClientListResponse{},
      Client: &armhanaonazure.SapMonitorsClient{},
      ListFunc: (&armhanaonazure.SapMonitorsClient{}).NewListPager,
			NewFunc: armhanaonazure.NewSapMonitorsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.HanaOnAzure/sapMonitors",
		},
	}

	for i := range tables {
		tables[i].Service = "armhanaonazure"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armhanaonazure()...)
}