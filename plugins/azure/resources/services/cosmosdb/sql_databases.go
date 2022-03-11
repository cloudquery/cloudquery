package cosmosdb

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CosmosDBSqlDatabases() *schema.Table {
	return &schema.Table{
		Name:          "azure_cosmosdb_sql_databases",
		Description:   "Azure Cosmos DB SQL database.",
		Resolver:      fetchCosmosdbSqlDatabases,
		Multiplex:     client.SubscriptionMultiplex,
		DeleteFilter:  client.DeleteSubscriptionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "database_id",
				Description: "Name of the Cosmos DB SQL database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SQLDatabaseGetProperties.Resource.ID"),
			},
			{
				Name:        "database_rid",
				Description: "A system generated property",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SQLDatabaseGetProperties.Resource.Rid"),
			},
			{
				Name:        "database_ts",
				Description: "A system generated property that denotes the last updated timestamp of the resource.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("SQLDatabaseGetProperties.Resource.Ts"),
			},
			{
				Name:        "database_etag",
				Description: "A system generated property representing the resource etag required for optimistic concurrency control.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SQLDatabaseGetProperties.Resource.Etag"),
			},
			{
				Name:        "database_colls",
				Description: "A system generated property that specified the addressable path of the collections resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SQLDatabaseGetProperties.Resource.Colls"),
			},
			{
				Name:        "database_users",
				Description: "A system generated property that specifies the addressable path of the users resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SQLDatabaseGetProperties.Resource.Users"),
			},
			{
				Name:        "sql_database_get_properties_throughput",
				Description: "Value of the Cosmos DB resource throughput or autoscaleSettings",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SQLDatabaseGetProperties.Options.Throughput"),
			},
			{
				Name:        "autoscale_settings_max_throughput",
				Description: "Represents maximum throughput, the resource can scale up to.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SQLDatabaseGetProperties.Options.AutoscaleSettings.MaxThroughput"),
			},
			{
				Name:        "id",
				Description: "The unique resource identifier of the ARM resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the ARM resource.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of Azure resource.",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "The location of the resource group to which the resource belongs.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCosmosdbSqlDatabases(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	m := meta.(*client.Client)
	accSvc := m.Services().CosmosDb.Accounts
	sqlSvc := m.Services().CosmosDb.SQL

	response, err := accSvc.List(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	if response.Value == nil {
		return nil
	}

	for _, account := range *response.Value {
		if account.Name == nil {
			m.Logger().Debug("could not found cosmosdb account name", "accountId", account.ID)
			continue
		}
		details, err := client.ParseResourceID(*account.ID)
		if err != nil {
			m.Logger().Debug("could not parse cosmosdb account id", "accountId", account.ID)
			continue
		}

		response, err := sqlSvc.ListSQLDatabases(ctx, details.ResourceGroup, *account.Name)
		if err != nil {
			return diag.WrapError(err)
		}
		if response.Value == nil {
			continue
		}
		res <- *response.Value
	}

	return nil
}
