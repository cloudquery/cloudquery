package cosmos

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func mongo_db_databases() *schema.Table {
	return &schema.Table{
		Name:      "azure_cosmos_mongo_db_databases",
		Resolver:  fetchMongoDbDatabases,
		Transform: transformers.TransformWithStruct(&armcosmos.MongoDBDatabaseGetResults{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
