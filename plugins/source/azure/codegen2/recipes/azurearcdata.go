// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"

func Armazurearcdata() []Table {
	tables := []Table{
		{
			Service:        "armazurearcdata",
			Name:           "postgres_instances",
			Struct:         &armazurearcdata.PostgresInstance{},
			ResponseStruct: &armazurearcdata.PostgresInstancesClientListResponse{},
			Client:         &armazurearcdata.PostgresInstancesClient{},
			ListFunc:       (&armazurearcdata.PostgresInstancesClient{}).NewListPager,
			NewFunc:        armazurearcdata.NewPostgresInstancesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.AzureArcData/postgresInstances",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_AzureArcData)`,
		},
		{
			Service:        "armazurearcdata",
			Name:           "sql_managed_instances",
			Struct:         &armazurearcdata.SQLManagedInstance{},
			ResponseStruct: &armazurearcdata.SQLManagedInstancesClientListResponse{},
			Client:         &armazurearcdata.SQLManagedInstancesClient{},
			ListFunc:       (&armazurearcdata.SQLManagedInstancesClient{}).NewListPager,
			NewFunc:        armazurearcdata.NewSQLManagedInstancesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.AzureArcData/sqlManagedInstances",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_AzureArcData)`,
		},
		{
			Service:        "armazurearcdata",
			Name:           "sql_server_instances",
			Struct:         &armazurearcdata.SQLServerInstance{},
			ResponseStruct: &armazurearcdata.SQLServerInstancesClientListResponse{},
			Client:         &armazurearcdata.SQLServerInstancesClient{},
			ListFunc:       (&armazurearcdata.SQLServerInstancesClient{}).NewListPager,
			NewFunc:        armazurearcdata.NewSQLServerInstancesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.AzureArcData/sqlServerInstances",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_AzureArcData)`,
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, Armazurearcdata()...)
}
