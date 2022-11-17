package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/cosmos"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func CosmosDB() []*resource.Resource {
	return []*resource.Resource{
		{
			Service:    "cosmosdb",
			SubService: "accounts",
			Struct:     new(armcosmos.DatabaseAccountGetResults),
			Resolver: &resource.FuncParams{
				Func: cosmos.DatabaseAccountsClient.NewListPager,
			},
			Children: []*resource.Resource{
				{
					SubService: "mongo_db_databases",
					Struct:     new(armcosmos.MongoDBDatabaseGetResults),
					Resolver: &resource.FuncParams{
						Func:   cosmos.MongoDBResourcesClient.NewListMongoDBDatabasesPager,
						Params: []string{"id.ResourceGroupName", "*account.Name"},
					},
				},
				{
					SubService: "sql_databases",
					Struct:     new(armcosmos.MongoDBDatabaseGetResults),
					Resolver: &resource.FuncParams{
						Func:   cosmos.SQLResourcesClient.NewListSQLDatabasesPager,
						Params: []string{"id.ResourceGroupName", "*account.Name"},
					},
				},
			},
		},
	}
}
