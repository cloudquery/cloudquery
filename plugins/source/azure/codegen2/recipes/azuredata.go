// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"

func Armazuredata() []Table {
	tables := []Table{
		{
			Service:        "armazuredata",
			Name:           "sql_server_registrations",
			Struct:         &armazuredata.SQLServerRegistration{},
			ResponseStruct: &armazuredata.SQLServerRegistrationsClientListResponse{},
			Client:         &armazuredata.SQLServerRegistrationsClient{},
			ListFunc:       (&armazuredata.SQLServerRegistrationsClient{}).NewListPager,
			NewFunc:        armazuredata.NewSQLServerRegistrationsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.AzureData/sqlServerRegistrations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_AzureData)`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, Armazuredata()...)
}
