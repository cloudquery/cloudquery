package cosmosdb

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CosmosDBMongoDBDatabases() *schema.Table {
	return &schema.Table{
		Name:         "azure_cosmosdb_mongodb_databases",
		Description:  "Azure Cosmos DB MongoDB database.",
		Resolver:     fetchCosmosdbMongodbDatabases,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "database_id",
				Description: "Name of the Cosmos DB MongoDB database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Resource.ID"),
			},
			{
				Name:          "database_rid",
				Description:   "A system generated property",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Resource.Rid"),
				IgnoreInTests: true,
			},
			{
				Name:        "database_ts",
				Description: "A system generated property that denotes the last updated timestamp of the resource.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Resource.Ts"),
				// This seems to be null or at least when provisioned with terraform
				IgnoreInTests: true,
			},
			{
				Name:          "database_etag",
				Description:   "A system generated property representing the resource etag required for optimistic concurrency control.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Resource.Etag"),
				IgnoreInTests: true,
			},
			{
				Name:        "throughput",
				Description: "Value of the Cosmos DB resource throughput or autoscaleSettings",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Options.Throughput"),
				// This seems to be null or at least when provisioned with terraform
				// even if we specify options in terraform - options looks always nil
				IgnoreInTests: true,
			},
			{
				Name:        "autoscale_settings_max_throughput",
				Description: "Represents maximum throughput, the resource can scale up to.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Options.AutoscaleSettings.MaxThroughput"),
				// This seems to be null or at least when provisioned with terraform
				// even if we specify options in terraform - options looks always nil
				IgnoreInTests: true,
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
				Name:          "location",
				Description:   "The location of the resource group to which the resource belongs.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
				// Not possible to set tags via terraform here
				// https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/cosmosdb_mongo_database
				IgnoreInTests: true,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCosmosdbMongodbDatabases(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	m := meta.(*client.Client)
	accSvc := m.Services().CosmosDb.Accounts
	mongoDBSvc := m.Services().CosmosDb.MongoDB

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

		response, err := mongoDBSvc.ListMongoDBDatabases(ctx, details.ResourceGroup, *account.Name)
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
