package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"

func ArmcosmosE() []Table {
	tables := []Table{
		{
			Service:        "armcosmos",
			Name:           "database_accounts",
			Struct:         &armcosmos.DatabaseAccountGetResults{},
			ResponseStruct: &armcosmos.DatabaseAccountsClientListResponse{},
			Client:         &armcosmos.DatabaseAccountsClient{},
			ListFunc:       (&armcosmos.DatabaseAccountsClient{}).NewListPager,
			NewFunc:        armcosmos.NewDatabaseAccountsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/databaseAccounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DocumentDB)`,
			Relations: []*Table{
				{
					Service:        "armcosmos",
					Name:           "mongo_db_databases",
					Struct:         &armcosmos.MongoDBDatabaseGetResults{},
					ResponseStruct: &armcosmos.MongoDBResourcesClientListMongoDBDatabasesResponse{},
					Client:         &armcosmos.MongoDBResourcesClient{},
					ListFunc:       (&armcosmos.MongoDBResourcesClient{}).NewListMongoDBDatabasesPager,
					NewFunc:        armcosmos.NewMongoDBResourcesClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases",
					SkipFetch:      true,
				},
				{
					Service:        "armcosmos",
					Name:           "sql_databases",
					Struct:         &armcosmos.SQLDatabaseGetResults{},
					ResponseStruct: &armcosmos.SQLResourcesClientListSQLDatabasesResponse{},
					Client:         &armcosmos.SQLResourcesClient{},
					ListFunc:       (&armcosmos.SQLResourcesClient{}).NewListSQLDatabasesPager,
					NewFunc:        armcosmos.NewSQLResourcesClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases",
					SkipFetch:      true,
				},
			},
		},
	}

	return tables
}

func init() {
	Tables = append(Tables, ArmcosmosE()...)
}
