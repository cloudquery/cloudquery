// Auto generated code - DO NOT EDIT.

package cosmosdb

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
)

func mongoDBDatabases() *schema.Table {
	return &schema.Table{
		Name:     "azure_cosmosdb_mongo_db_databases",
		Resolver: fetchCosmosDBMongoDBDatabases,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cosmosdb_account_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentResourceFieldResolver("id"),
			},
			{
				Name:     "resource",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Resource"),
			},
			{
				Name:     "options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Options"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}

func fetchCosmosDBMongoDBDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().CosmosDB.MongoDBDatabases

	account := parent.Item.(documentdb.DatabaseAccountGetResults)
	resource, err := client.ParseResourceID(*account.ID)
	if err != nil {
		return err
	}
	response, err := svc.ListMongoDBDatabases(ctx, resource.ResourceGroup, *account.Name)
	if err != nil {
		return err
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value

	return nil
}
