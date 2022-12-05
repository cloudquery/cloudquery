// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"

func Armappconfiguration() []Table {
	tables := []Table{
		{
      Name: "configuration_store",
      Struct: &armappconfiguration.ConfigurationStore{},
      ResponseStruct: &armappconfiguration.ConfigurationStoresClientListResponse{},
      Client: &armappconfiguration.ConfigurationStoresClient{},
      ListFunc: (&armappconfiguration.ConfigurationStoresClient{}).NewListPager,
			NewFunc: armappconfiguration.NewConfigurationStoresClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/configurationStores",
		},
	}

	for i := range tables {
		tables[i].Service = "armappconfiguration"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armappconfiguration()...)
}