// Auto generated code - DO NOT EDIT.

package cosmosdb

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/cosmos-db/mgmt/documentdb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func SQLDatabases() *schema.Table {
	return &schema.Table{
		Name:     "azure_cosmosdb_sql_databases",
		Resolver: fetchCosmosDBSQLDatabases,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "sql_database_get_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SQLDatabaseGetProperties"),
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

func fetchCosmosDBSQLDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().CosmosDB.SQLDatabases

	account := parent.Item.(documentdb.DatabaseAccountGetResults)
	resource, err := client.ParseResourceID(*account.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	response, err := svc.ListSQLDatabases(ctx, resource.ResourceGroup, *account.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value

	return nil
}
