// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"

func Armcosmos() []Table {
	tables := []Table{
		{
      Name: "restorable_sql_container_get_result",
      Struct: &armcosmos.RestorableSQLContainerGetResult{},
      ResponseStruct: &armcosmos.RestorableSQLContainersClientListResponse{},
      Client: &armcosmos.RestorableSQLContainersClient{},
      ListFunc: (&armcosmos.RestorableSQLContainersClient{}).NewListPager,
			NewFunc: armcosmos.NewRestorableSQLContainersClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableSqlContainers",
		},
		{
      Name: "database_account_get_results",
      Struct: &armcosmos.DatabaseAccountGetResults{},
      ResponseStruct: &armcosmos.DatabaseAccountsClientListResponse{},
      Client: &armcosmos.DatabaseAccountsClient{},
      ListFunc: (&armcosmos.DatabaseAccountsClient{}).NewListPager,
			NewFunc: armcosmos.NewDatabaseAccountsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/databaseAccounts",
		},
		{
      Name: "location_get_result",
      Struct: &armcosmos.LocationGetResult{},
      ResponseStruct: &armcosmos.LocationsClientListResponse{},
      Client: &armcosmos.LocationsClient{},
      ListFunc: (&armcosmos.LocationsClient{}).NewListPager,
			NewFunc: armcosmos.NewLocationsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations",
		},
		{
      Name: "database_restore_resource",
      Struct: &armcosmos.DatabaseRestoreResource{},
      ResponseStruct: &armcosmos.RestorableSQLResourcesClientListResponse{},
      Client: &armcosmos.RestorableSQLResourcesClient{},
      ListFunc: (&armcosmos.RestorableSQLResourcesClient{}).NewListPager,
			NewFunc: armcosmos.NewRestorableSQLResourcesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableSqlResources",
		},
		{
      Name: "restorable_mongodb_database_get_result",
      Struct: &armcosmos.RestorableMongodbDatabaseGetResult{},
      ResponseStruct: &armcosmos.RestorableMongodbDatabasesClientListResponse{},
      Client: &armcosmos.RestorableMongodbDatabasesClient{},
      ListFunc: (&armcosmos.RestorableMongodbDatabasesClient{}).NewListPager,
			NewFunc: armcosmos.NewRestorableMongodbDatabasesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbDatabases",
		},
		{
      Name: "database_restore_resource",
      Struct: &armcosmos.DatabaseRestoreResource{},
      ResponseStruct: &armcosmos.RestorableMongodbResourcesClientListResponse{},
      Client: &armcosmos.RestorableMongodbResourcesClient{},
      ListFunc: (&armcosmos.RestorableMongodbResourcesClient{}).NewListPager,
			NewFunc: armcosmos.NewRestorableMongodbResourcesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbResources",
		},
		{
      Name: "restorable_database_account_get_result",
      Struct: &armcosmos.RestorableDatabaseAccountGetResult{},
      ResponseStruct: &armcosmos.RestorableDatabaseAccountsClientListResponse{},
      Client: &armcosmos.RestorableDatabaseAccountsClient{},
      ListFunc: (&armcosmos.RestorableDatabaseAccountsClient{}).NewListPager,
			NewFunc: armcosmos.NewRestorableDatabaseAccountsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/restorableDatabaseAccounts",
		},
		{
      Name: "data_center_resource",
      Struct: &armcosmos.DataCenterResource{},
      ResponseStruct: &armcosmos.CassandraDataCentersClientListResponse{},
      Client: &armcosmos.CassandraDataCentersClient{},
      ListFunc: (&armcosmos.CassandraDataCentersClient{}).NewListPager,
			NewFunc: armcosmos.NewCassandraDataCentersClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/dataCenters",
		},
		{
      Name: "restorable_mongodb_collection_get_result",
      Struct: &armcosmos.RestorableMongodbCollectionGetResult{},
      ResponseStruct: &armcosmos.RestorableMongodbCollectionsClientListResponse{},
      Client: &armcosmos.RestorableMongodbCollectionsClient{},
      ListFunc: (&armcosmos.RestorableMongodbCollectionsClient{}).NewListPager,
			NewFunc: armcosmos.NewRestorableMongodbCollectionsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbCollections",
		},
		{
      Name: "restorable_sql_database_get_result",
      Struct: &armcosmos.RestorableSQLDatabaseGetResult{},
      ResponseStruct: &armcosmos.RestorableSQLDatabasesClientListResponse{},
      Client: &armcosmos.RestorableSQLDatabasesClient{},
      ListFunc: (&armcosmos.RestorableSQLDatabasesClient{}).NewListPager,
			NewFunc: armcosmos.NewRestorableSQLDatabasesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableSqlDatabases",
		},
	}

	for i := range tables {
		tables[i].Service = "armcosmos"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armcosmos()...)
}