package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func MySQLServers() *schema.Table {
	return &schema.Table{
		Name:         "azure_mysql_servers",
		Resolver:     fetchMySQLServers,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "identity_principal_id",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:     "identity_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Identity.Type"),
			},
			{
				Name:     "identity_tenant_id",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:     "sku_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Sku.Name"),
			},
			{
				Name:     "sku_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Sku.Tier"),
			},
			{
				Name:     "sku_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Sku.Capacity"),
			},
			{
				Name:     "sku_size",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Sku.Size"),
			},
			{
				Name:     "sku_family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Sku.Family"),
			},
			{
				Name:     "administrator_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.AdministratorLogin"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.Version"),
			},
			{
				Name:     "ssl_enforcement",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.SslEnforcement"),
			},
			{
				Name:     "minimal_tls_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.MinimalTLSVersion"),
			},
			{
				Name:     "byok_enforcement",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.ByokEnforcement"),
			},
			{
				Name:     "infrastructure_encryption",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.InfrastructureEncryption"),
			},
			{
				Name:     "user_visible_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.UserVisibleState"),
			},
			{
				Name:     "fully_qualified_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.FullyQualifiedDomainName"),
			},
			{
				Name:     "earliest_restore_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ServerProperties.EarliestRestoreDate.Time"),
			},
			{
				Name:     "storage_profile_backup_retention_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ServerProperties.StorageProfile.BackupRetentionDays"),
			},
			{
				Name:     "storage_profile_geo_redundant_backup",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.StorageProfile.GeoRedundantBackup"),
			},
			{
				Name:     "storage_profile_storage_mb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ServerProperties.StorageProfile.StorageMB"),
			},
			{
				Name:     "storage_profile_storage_autogrow",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.StorageProfile.StorageAutogrow"),
			},
			{
				Name:     "replication_role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.ReplicationRole"),
			},
			{
				Name:     "master_server_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.MasterServerID"),
			},
			{
				Name:     "replica_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ServerProperties.ReplicaCapacity"),
			},
			{
				Name:     "public_network_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.PublicNetworkAccess"),
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "azure_mysql_server_private_endpoint_connections",
				Resolver: fetchMySQLServerPrivateEndpointConnections,
				Columns: []schema.Column{
					{
						Name:     "server_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name:     "private_endpoint_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Properties.PrivateEndpoint.ID"),
					},
					{
						Name:     "private_link_service_connection_state_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:     "private_link_service_connection_state_description",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:     "private_link_service_connection_state_actions_required",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Properties.PrivateLinkServiceConnectionState.ActionsRequired"),
					},
					{
						Name:     "provisioning_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Properties.ProvisioningState"),
					},
				},
			},
			{
				Name:     "azure_mysql_server_configurations",
				Resolver: fetchMySQLServerConfigurations,
				Columns: []schema.Column{
					{
						Name:     "server_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "value",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConfigurationProperties.Value"),
					},
					{
						Name:     "description",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConfigurationProperties.Description"),
					},
					{
						Name:     "default_value",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConfigurationProperties.DefaultValue"),
					},
					{
						Name:     "data_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConfigurationProperties.DataType"),
					},
					{
						Name:     "allowed_values",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConfigurationProperties.AllowedValues"),
					},
					{
						Name:     "source",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ConfigurationProperties.Source"),
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchMySQLServers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().MySQL.Servers
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
func fetchMySQLServerPrivateEndpointConnections(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	server := parent.Item.(mysql.Server)
	if server.PrivateEndpointConnections == nil {
		return nil
	}
	res <- *server.PrivateEndpointConnections
	return nil
}
func fetchMySQLServerConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	server := parent.Item.(mysql.Server)
	svc := meta.(*client.Client).Services().MySQL.Configuration

	resourceDetails, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}
	configurations, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *server.Name)
	if err != nil {
		return err
	}
	if configurations.Value == nil {
		return nil
	}
	res <- *configurations.Value
	return nil
}
