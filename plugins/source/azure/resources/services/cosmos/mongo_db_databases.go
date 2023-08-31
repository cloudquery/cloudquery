package cosmos

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func mongo_db_databases() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cosmos_mongo_db_databases",
		Resolver:             fetchMongoDbDatabases,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2022-05-15/mongo-db-resources/list-mongo-db-databases?tabs=HTTP#mongodbdatabasegetresults",
		Transform:            transformers.TransformWithStruct(&armcosmos.MongoDBDatabaseGetResults{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{},
	}
}
