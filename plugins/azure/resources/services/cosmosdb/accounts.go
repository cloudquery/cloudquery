package cosmosdb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CosmosDBAccounts() *schema.Table {
	return &schema.Table{
		Name:         "azure_cosmosdb_accounts",
		Description:  "Azure Cosmos DB database account.",
		Resolver:     fetchCosmosdbAccounts,
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
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseAccountGetProperties.ProvisioningState"),
			},
			{
				Name:        "document_endpoint",
				Description: "The connection endpoint for the Cosmos DB database account.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.DocumentEndpoint"),
			},
			{
				Name:        "database_account_offer_type",
				Description: "The offer type for the Cosmos DB database account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.DatabaseAccountOfferType"),
			},
			{
				Name:        "ip_rules",
				Description: "List of IpRules.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveCosmosdbAccountsIpRules,
			},
			{
				Name:        "capabilities",
				Description: "Capability cosmos DB capability object",
				Type:        schema.TypeStringArray,
				Resolver:    resolveCosmosdbAccountsCapabilities,
			},
			{
				Name:        "is_virtual_network_filter_enabled",
				Description: "Flag to indicate whether to enable/disable Virtual Network ACL rules.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.IsVirtualNetworkFilterEnabled"),
			},
			{
				Name:        "enable_automatic_failover",
				Description: "Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.EnableAutomaticFailover"),
			},
			{
				Name:        "consistency_policy_default_consistency_level",
				Description: "The default consistency level and configuration settings of the Cosmos DB account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.ConsistencyPolicy.DefaultConsistencyLevel"),
			},
			{
				Name:        "consistency_policy_max_staleness_prefix",
				Description: "When used with the Bounded Staleness consistency level, this value represents the number of stale requests tolerated",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.ConsistencyPolicy.MaxStalenessPrefix"),
			},
			{
				Name:        "consistency_policy_max_interval_in_seconds",
				Description: "When used with the Bounded Staleness consistency level, this value represents the time amount of staleness (in seconds) tolerated",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.ConsistencyPolicy.MaxIntervalInSeconds"),
			},
			{
				Name:        "virtual_network_rules",
				Description: "List of Virtual Network ACL rules configured for the Cosmos DB account.",
				Type:        schema.TypeJSON,
				Resolver:    resolveCosmosdbAccountsVirtualNetworkRules,
			},
			{
				Name:        "enable_multiple_write_locations",
				Description: "Enables the account to write in multiple locations",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.EnableMultipleWriteLocations"),
			},
			{
				Name:        "enable_cassandra_connector",
				Description: "Enables the cassandra connector on the Cosmos DB C* account",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.EnableCassandraConnector"),
			},
			{
				Name:        "connector_offer",
				Description: "The cassandra connector offer type for the Cosmos DB database C* account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.ConnectorOffer"),
			},
			{
				Name:        "disable_key_based_metadata_write_access",
				Description: "Disable write operations on metadata resources (databases, containers, throughput) via account keys",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.DisableKeyBasedMetadataWriteAccess"),
			},
			{
				Name:        "key_vault_key_uri",
				Description: "The URI of the key vault",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.KeyVaultKeyURI"),
			},
			{
				Name:        "public_network_access",
				Description: "Whether requests from Public Network are allowed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.PublicNetworkAccess"),
			},
			{
				Name:        "enable_free_tier",
				Description: "Flag to indicate whether Free Tier is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.EnableFreeTier"),
			},
			{
				Name:        "api_properties_server_version",
				Description: "Describes the ServerVersion of an a MongoDB account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.APIProperties.ServerVersion"),
			},
			{
				Name:        "enable_analytical_storage",
				Description: "Flag to indicate whether to enable storage analytics.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DatabaseAccountGetProperties.EnableAnalyticalStorage"),
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
		Relations: []*schema.Table{
			{
				Name:        "azure_cosmosdb_account_write_locations",
				Description: "Location a region in which the Azure Cosmos DB database account is deployed.",
				Resolver:    fetchCosmosdbAccountWriteLocations,
				Columns: []schema.Column{
					{
						Name:        "account_cq_id",
						Description: "Unique CloudQuery ID of azure_cosmosdb_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "The unique identifier of the region within the database account",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "location_name",
						Description: "The name of the region.",
						Type:        schema.TypeString,
					},
					{
						Name:        "document_endpoint",
						Description: "The connection endpoint for the specific region",
						Type:        schema.TypeString,
					},
					{
						Name: "provisioning_state",
						Type: schema.TypeString,
					},
					{
						Name:        "failover_priority",
						Description: "The failover priority of the region",
						Type:        schema.TypeInt,
					},
					{
						Name:        "is_zone_redundant",
						Description: "Flag to indicate whether or not this region is an AvailabilityZone region",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "azure_cosmosdb_account_read_locations",
				Description: "Location a region in which the Azure Cosmos DB database account is deployed.",
				Resolver:    fetchCosmosdbAccountReadLocations,
				Columns: []schema.Column{
					{
						Name:        "account_cq_id",
						Description: "Unique CloudQuery ID of azure_cosmosdb_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "The unique identifier of the region within the database account",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "location_name",
						Description: "The name of the region.",
						Type:        schema.TypeString,
					},
					{
						Name:        "document_endpoint",
						Description: "The connection endpoint for the specific region",
						Type:        schema.TypeString,
					},
					{
						Name: "provisioning_state",
						Type: schema.TypeString,
					},
					{
						Name:        "failover_priority",
						Description: "The failover priority of the region",
						Type:        schema.TypeInt,
					},
					{
						Name:        "is_zone_redundant",
						Description: "Flag to indicate whether or not this region is an AvailabilityZone region",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "azure_cosmosdb_account_locations",
				Description: "Location a region in which the Azure Cosmos DB database account is deployed.",
				Resolver:    fetchCosmosdbAccountLocations,
				Columns: []schema.Column{
					{
						Name:        "account_cq_id",
						Description: "Unique CloudQuery ID of azure_cosmosdb_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "The unique identifier of the region within the database account",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "location_name",
						Description: "The name of the region.",
						Type:        schema.TypeString,
					},
					{
						Name:        "document_endpoint",
						Description: "The connection endpoint for the specific region",
						Type:        schema.TypeString,
					},
					{
						Name: "provisioning_state",
						Type: schema.TypeString,
					},
					{
						Name:        "failover_priority",
						Description: "The failover priority of the region",
						Type:        schema.TypeInt,
					},
					{
						Name:        "is_zone_redundant",
						Description: "Flag to indicate whether or not this region is an AvailabilityZone region",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "azure_cosmosdb_account_failover_policies",
				Description: "FailoverPolicy the failover policy for a given region of a database account.",
				Resolver:    fetchCosmosdbAccountFailoverPolicies,
				Columns: []schema.Column{
					{
						Name:        "account_cq_id",
						Description: "Unique CloudQuery ID of azure_cosmosdb_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "The unique identifier of the region in which the database account replicates to",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "location_name",
						Description: "The name of the region in which the database account exists.",
						Type:        schema.TypeString,
					},
					{
						Name:        "failover_priority",
						Description: "The failover priority of the region",
						Type:        schema.TypeInt,
					},
				},
			},
			{
				Name:        "azure_cosmosdb_account_private_endpoint_connections",
				Description: "PrivateEndpointConnection a private endpoint connection",
				Resolver:    fetchCosmosdbAccountPrivateEndpointConnections,
				Columns: []schema.Column{
					{
						Name:        "account_cq_id",
						Description: "Unique CloudQuery ID of azure_cosmosdb_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "private_endpoint_id",
						Description: "Resource id of the private endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateEndpoint.ID"),
					},
					{
						Name:        "status",
						Description: "The private link service connection status.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:        "actions_required",
						Description: "Any action that is required beyond basic workflow (approve/ reject/ disconnect)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.ActionsRequired"),
					},
					{
						Name:        "description",
						Description: "The private link service connection description.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:        "group_id",
						Description: "Group id of the private endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.GroupID"),
					},
					{
						Name:        "provisioning_state",
						Description: "Provisioning state of the private endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.ProvisioningState"),
					},
					{
						Name:        "id",
						Description: "Fully qualified resource ID for the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The name of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of the resource",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_cosmosdb_account_cors",
				Description: "CorsPolicy the CORS policy for the Cosmos DB database account.",
				Resolver:    fetchCosmosdbAccountCors,
				Columns: []schema.Column{
					{
						Name:        "account_cq_id",
						Description: "Unique CloudQuery ID of azure_cosmosdb_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allowed_origins",
						Description: "The origin domains that are permitted to make a request against the service via CORS.",
						Type:        schema.TypeString,
					},
					{
						Name:        "allowed_methods",
						Description: "The methods (HTTP request verbs) that the origin domain may use for a CORS request.",
						Type:        schema.TypeString,
					},
					{
						Name:        "allowed_headers",
						Description: "The request headers that the origin domain may specify on the CORS request.",
						Type:        schema.TypeString,
					},
					{
						Name:        "exposed_headers",
						Description: "The response headers that may be sent in the response to the CORS request and exposed by the browser to the request issuer.",
						Type:        schema.TypeString,
					},
					{
						Name:        "max_age_in_seconds",
						Description: "The maximum amount time that a browser should cache the preflight OPTIONS request.",
						Type:        schema.TypeBigInt,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCosmosdbAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().CosmosDb.Accounts
	response, err := svc.List(ctx)
	if err != nil {
		return err
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value
	return nil
}
func resolveCosmosdbAccountsIpRules(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	account, ok := resource.Item.(documentdb.DatabaseAccountGetResults)
	if !ok {
		return fmt.Errorf("expected to have documentdb.DatabaseAccountGetResults but got %T", resource.Item)
	}
	if account.IPRules == nil {
		return nil
	}
	ipRules := make([]string, len(*account.IPRules))
	for _, rule := range *account.IPRules {
		ipRules = append(ipRules, *rule.IPAddressOrRange)
	}
	return resource.Set(c.Name, ipRules)
}
func resolveCosmosdbAccountsVirtualNetworkRules(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	account, ok := resource.Item.(documentdb.DatabaseAccountGetResults)
	if !ok {
		return fmt.Errorf("expected to have documentdb.DatabaseAccountGetResults but got %T", resource.Item)
	}
	if account.VirtualNetworkRules == nil {
		return nil
	}
	b, err := json.Marshal(account.VirtualNetworkRules)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}
func resolveCosmosdbAccountsCapabilities(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	account, ok := resource.Item.(documentdb.DatabaseAccountGetResults)
	if !ok {
		return fmt.Errorf("expected to have documentdb.DatabaseAccountGetResults but got %T", resource.Item)
	}
	if account.Capabilities == nil {
		return nil
	}
	capabilities := make([]string, len(*account.Capabilities))
	for _, capability := range *account.Capabilities {
		capabilities = append(capabilities, *capability.Name)
	}
	return resource.Set(c.Name, capabilities)
}
func fetchCosmosdbAccountWriteLocations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	account, ok := parent.Item.(documentdb.DatabaseAccountGetResults)
	if !ok {
		return fmt.Errorf("expected to have documentdb.DatabaseAccountGetResults but got %T", parent.Item)
	}
	if account.WriteLocations == nil {
		return nil
	}
	res <- *account.WriteLocations
	return nil
}
func fetchCosmosdbAccountReadLocations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	account, ok := parent.Item.(documentdb.DatabaseAccountGetResults)
	if !ok {
		return fmt.Errorf("expected to have documentdb.DatabaseAccountGetResults but got %T", parent.Item)
	}
	if account.ReadLocations == nil {
		return nil
	}
	res <- *account.ReadLocations
	return nil
}
func fetchCosmosdbAccountLocations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	account, ok := parent.Item.(documentdb.DatabaseAccountGetResults)
	if !ok {
		return fmt.Errorf("expected to have documentdb.DatabaseAccountGetResults but got %T", parent.Item)
	}
	if account.Locations == nil {
		return nil
	}
	res <- *account.Locations
	return nil
}
func fetchCosmosdbAccountFailoverPolicies(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	account, ok := parent.Item.(documentdb.DatabaseAccountGetResults)
	if !ok {
		return fmt.Errorf("expected to have documentdb.DatabaseAccountGetResults but got %T", parent.Item)
	}
	if account.FailoverPolicies == nil {
		return nil
	}
	res <- *account.FailoverPolicies
	return nil
}
func fetchCosmosdbAccountPrivateEndpointConnections(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	account, ok := parent.Item.(documentdb.DatabaseAccountGetResults)
	if !ok {
		return fmt.Errorf("expected to have documentdb.DatabaseAccountGetResults but got %T", parent.Item)
	}
	if account.PrivateEndpointConnections == nil {
		return nil
	}
	res <- *account.PrivateEndpointConnections
	return nil
}
func fetchCosmosdbAccountCors(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	account, ok := parent.Item.(documentdb.DatabaseAccountGetResults)
	if !ok {
		return fmt.Errorf("expected to have documentdb.DatabaseAccountGetResults but got %T", parent.Item)
	}
	if account.Cors == nil {
		return nil
	}
	res <- *account.Cors
	return nil
}
