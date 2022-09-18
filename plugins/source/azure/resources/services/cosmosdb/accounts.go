// Auto generated code - DO NOT EDIT.

package cosmosdb

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:      "azure_cosmosdb_accounts",
		Resolver:  fetchCosmosDBAccounts,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "document_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DocumentEndpoint"),
			},
			{
				Name:     "database_account_offer_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseAccountOfferType"),
			},
			{
				Name:     "ip_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IPRules"),
			},
			{
				Name:     "is_virtual_network_filter_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsVirtualNetworkFilterEnabled"),
			},
			{
				Name:     "enable_automatic_failover",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableAutomaticFailover"),
			},
			{
				Name:     "consistency_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConsistencyPolicy"),
			},
			{
				Name:     "capabilities",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Capabilities"),
			},
			{
				Name:     "write_locations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WriteLocations"),
			},
			{
				Name:     "read_locations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ReadLocations"),
			},
			{
				Name:     "locations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Locations"),
			},
			{
				Name:     "failover_policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FailoverPolicies"),
			},
			{
				Name:     "virtual_network_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VirtualNetworkRules"),
			},
			{
				Name:     "private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateEndpointConnections"),
			},
			{
				Name:     "enable_multiple_write_locations",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableMultipleWriteLocations"),
			},
			{
				Name:     "enable_cassandra_connector",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableCassandraConnector"),
			},
			{
				Name:     "connector_offer",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectorOffer"),
			},
			{
				Name:     "disable_key_based_metadata_write_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisableKeyBasedMetadataWriteAccess"),
			},
			{
				Name:     "key_vault_key_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyVaultKeyURI"),
			},
			{
				Name:     "public_network_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicNetworkAccess"),
			},
			{
				Name:     "enable_free_tier",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableFreeTier"),
			},
			{
				Name:     "api_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("APIProperties"),
			},
			{
				Name:     "enable_analytical_storage",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableAnalyticalStorage"),
			},
			{
				Name:     "cors",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Cors"),
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

		Relations: []*schema.Table{
			mongoDBDatabases(),
			sQLDatabases(),
		},
	}
}

func fetchCosmosDBAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().CosmosDB.Accounts

	response, err := svc.List(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value

	return nil
}
