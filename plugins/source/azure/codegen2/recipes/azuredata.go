// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"

func Armazuredata() []Table {
	tables := []Table{
		{
			Name:           "sql_server_registration",
			Struct:         &armazuredata.SQLServerRegistration{},
			ResponseStruct: &armazuredata.SQLServerRegistrationsClientListResponse{},
			Client:         &armazuredata.SQLServerRegistrationsClient{},
			ListFunc:       (&armazuredata.SQLServerRegistrationsClient{}).NewListPager,
			NewFunc:        armazuredata.NewSQLServerRegistrationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.AzureData/sqlServerRegistrations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.AzureData")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armazuredata"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armazuredata()...)
}
