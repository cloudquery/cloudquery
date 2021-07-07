package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func MySQLServers() *schema.Table {
	return &schema.Table{
		Name:        "azure_mysql_servers",
		Description: "Azure mysql server",
		Resolver:    fetchMySQLServers,
		Multiplex:   client.SubscriptionMultiplex,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "identity_principal_id",
				Description: "The Azure Active Directory principal id",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_type",
				Description: "The identity type Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active Directory principal for the resource Possible values include: 'SystemAssigned'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "The Azure Active Directory tenant id",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "sku_name",
				Description: "The name of the sku, typically, tier + family + cores, eg B_Gen4_1, GP_Gen5_8",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The tier of the particular SKU, eg Basic Possible values include: 'Basic', 'GeneralPurpose', 'MemoryOptimized'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "sku_capacity",
				Description: "The scale up/out capacity, representing server's compute units",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Sku.Capacity"),
			},
			{
				Name:        "sku_size",
				Description: "The size code, to be interpreted by resource as appropriate",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Size"),
			},
			{
				Name:        "sku_family",
				Description: "The family of hardware",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Family"),
			},
			{
				Name:        "administrator_login",
				Description: "The administrator's login name of a server Can only be specified when the server is being created (and is required for creation)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.AdministratorLogin"),
			},
			{
				Name:        "version",
				Description: "Server version Possible values include: 'FiveFullStopSix', 'FiveFullStopSeven', 'EightFullStopZero'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.Version"),
			},
			{
				Name:        "ssl_enforcement",
				Description: "Enable ssl enforcement or not when connect to server Possible values include: 'SslEnforcementEnumEnabled', 'SslEnforcementEnumDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.SslEnforcement"),
			},
			{
				Name:        "minimal_tls_version",
				Description: "Enforce a minimal Tls version for the server Possible values include: 'TLS10', 'TLS11', 'TLS12', 'TLSEnforcementDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.MinimalTLSVersion"),
			},
			{
				Name:        "byok_enforcement",
				Description: "Status showing whether the server data encryption is enabled with customer-managed keys",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.ByokEnforcement"),
			},
			{
				Name:        "infrastructure_encryption",
				Description: "Status showing whether the server enabled infrastructure encryption Possible values include: 'InfrastructureEncryptionEnabled', 'InfrastructureEncryptionDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.InfrastructureEncryption"),
			},
			{
				Name:        "user_visible_state",
				Description: "A state of a server that is visible to user Possible values include: 'ServerStateReady', 'ServerStateDropping', 'ServerStateDisabled', 'ServerStateInaccessible'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.UserVisibleState"),
			},
			{
				Name:        "fully_qualified_domain_name",
				Description: "The fully qualified domain name of a server",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.FullyQualifiedDomainName"),
			},
			{
				Name:        "earliest_restore_date_time",
				Description: "Earliest restore point creation time (ISO8601 format)",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("ServerProperties.EarliestRestoreDate.Time"),
			},
			{
				Name:        "storage_profile_backup_retention_days",
				Description: "Backup retention days for the server",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ServerProperties.StorageProfile.BackupRetentionDays"),
			},
			{
				Name:        "storage_profile_geo_redundant_backup",
				Description: "Enable Geo-redundant or not for server backup Possible values include: 'Enabled', 'Disabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.StorageProfile.GeoRedundantBackup"),
			},
			{
				Name:        "storage_profile_storage_mb",
				Description: "Max storage allowed for a server",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ServerProperties.StorageProfile.StorageMB"),
			},
			{
				Name:        "storage_profile_storage_autogrow",
				Description: "Enable Storage Auto Grow Possible values include: 'StorageAutogrowEnabled', 'StorageAutogrowDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.StorageProfile.StorageAutogrow"),
			},
			{
				Name:        "replication_role",
				Description: "The replication role of the server",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.ReplicationRole"),
			},
			{
				Name:        "master_server_id",
				Description: "The master server id of a replica server",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.MasterServerID"),
			},
			{
				Name:        "replica_capacity",
				Description: "The maximum number of replicas that a master server can have",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ServerProperties.ReplicaCapacity"),
			},
			{
				Name:        "public_network_access",
				Description: "Whether or not public network access is allowed for this server Value is optional but if passed in, must be 'Enabled' or 'Disabled' Possible values include: 'PublicNetworkAccessEnumEnabled', 'PublicNetworkAccessEnumDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.PublicNetworkAccess"),
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "location",
				Description: "The geo-location where the resource lives",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "Fully qualified resource ID for the resource Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}",
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
				Description: "The type of the resource Eg \"MicrosoftCompute/virtualMachines\" or \"MicrosoftStorage/storageAccounts\"",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_mysql_server_private_endpoint_connections",
				Description: "Azure mysql server private endpoint connection",
				Resolver:    fetchMySQLServerPrivateEndpointConnections,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"server_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "server_cq_id",
						Description: "Unique ID of azure_mysql_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Resource Id of the private endpoint connection",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "private_endpoint_id",
						Description: "Resource id of the private endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateEndpoint.ID"),
					},
					{
						Name:        "private_link_service_connection_state_status",
						Description: "The private link service connection status Possible values include: 'Approved', 'Pending', 'Rejected', 'Disconnected'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:        "private_link_service_connection_state_description",
						Description: "The private link service connection description",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:        "private_link_service_connection_state_actions_required",
						Description: "The actions required for private link service connection Possible values include: 'None'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.ActionsRequired"),
					},
					{
						Name:        "provisioning_state",
						Description: "State of the private endpoint connection Possible values include: 'Approving', 'Ready', 'Dropping', 'Failed', 'Rejecting'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.ProvisioningState"),
					},
				},
			},
			{
				Name:        "azure_mysql_server_configurations",
				Description: "Azure mysql server configuration",
				Resolver:    fetchMySQLServerConfigurations,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"server_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "server_cq_id",
						Description: "Unique ID of azure_mysql_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "value",
						Description: "Value of the configuration",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConfigurationProperties.Value"),
					},
					{
						Name:        "description",
						Description: "Description of the configuration",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConfigurationProperties.Description"),
					},
					{
						Name:        "default_value",
						Description: "Default value of the configuration",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConfigurationProperties.DefaultValue"),
					},
					{
						Name:        "data_type",
						Description: "Data type of the configuration",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConfigurationProperties.DataType"),
					},
					{
						Name:        "allowed_values",
						Description: "Allowed values of the configuration",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConfigurationProperties.AllowedValues"),
					},
					{
						Name:        "source",
						Description: "Source of the configuration",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ConfigurationProperties.Source"),
					},
					{
						Name:        "id",
						Description: "Fully qualified resource ID for the resource Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}",
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
						Description: "The type of the resource Eg \"MicrosoftCompute/virtualMachines\" or \"MicrosoftStorage/storageAccounts\"",
						Type:        schema.TypeString,
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
