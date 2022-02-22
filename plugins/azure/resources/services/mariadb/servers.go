package mariadb

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func MariadbServers() *schema.Table {
	return &schema.Table{
		Name:         "azure_mariadb_servers",
		Description:  "Server represents a server.",
		Resolver:     fetchMariadbServers,
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
				Name:        "sku_name",
				Description: "Name - The name of the sku.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The tier of the particular SKU.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "sku_capacity",
				Description: "The scale up/out capacity, representing server's compute units.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Sku.Capacity"),
			},
			{
				Name:        "sku_size",
				Description: "The size code, to be interpreted by resource as appropriate.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Size"),
			},
			{
				Name:        "sku_family",
				Description: "The family of hardware.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Family"),
			},
			{
				Name:        "administrator_login",
				Description: "The administrator's login name of a server.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.AdministratorLogin"),
			},
			{
				Name:        "version",
				Description: "Server version.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.Version"),
			},
			{
				Name:        "ssl_enforcement",
				Description: "Enable ssl enforcement or not when connect to server.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.SslEnforcement"),
			},
			{
				Name:        "user_visible_state",
				Description: "A state of a server that is visible to user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.UserVisibleState"),
			},
			{
				Name:        "fully_qualified_domain_name",
				Description: "The fully qualified domain name of a server.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.FullyQualifiedDomainName"),
			},
			{
				Name:     "earliest_restore_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ServerProperties.EarliestRestoreDate.Time"),
			},
			{
				Name:        "backup_retention_days",
				Description: "Backup retention days for the server.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ServerProperties.StorageProfile.BackupRetentionDays"),
			},
			{
				Name:        "geo_redundant_backup",
				Description: "Enable Geo-redundant or not for server backup.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.StorageProfile.GeoRedundantBackup"),
			},
			{
				Name:        "storage_mb",
				Description: "Max storage allowed for a server.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ServerProperties.StorageProfile.StorageMB"),
			},
			{
				Name:        "storage_autogrow",
				Description: "Enable Storage Auto Grow",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.StorageProfile.StorageAutogrow"),
			},
			{
				Name:        "replication_role",
				Description: "The replication role of the server.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.ReplicationRole"),
			},
			{
				Name:        "master_server_id",
				Description: "The master server id of a replica server.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.MasterServerID"),
			},
			{
				Name:        "replica_capacity",
				Description: "The maximum number of replicas that a master server can have.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ServerProperties.ReplicaCapacity"),
			},
			{
				Name:        "public_network_access",
				Description: "Whether or not public network access is allowed for this server.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.PublicNetworkAccess"),
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "location",
				Description: "The geo-location where the resource lives",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "Fully qualified resource ID for the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_mariadb_server_private_endpoint_connections",
				Description: "List of private endpoint connections on a server",
				Resolver:    resolveServerPrivateEndpointConnections,
				Columns: []schema.Column{
					{
						Name:        "server_cq_id",
						Description: "Unique ID of azure_mariadb_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Resource Id of the private endpoint connection.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "status",
						Description: "The private link service connection status.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:        "status_description",
						Description: "The private link service connection description.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:        "provisioning_state",
						Description: "State of the private endpoint connection.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.ProvisioningState"),
					},
				},
			},
			{
				Name:        "azure_mariadb_server_configurations",
				Description: "MariaDB server configuration",
				Resolver:    resolveMariadbServerConfigurations,
				Columns: []schema.Column{
					{
						Name:        "server_cq_id",
						Description: "Unique ID of azure_mariadb_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Fully qualified resource ID for the resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The name of the resource.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of the resource.",
						Type:        schema.TypeString,
					},
					{
						Name:        "value",
						Description: "Value of the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "Description of the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "default_value",
						Description: "Default value of the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "data_type",
						Description: "Data type of the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "allowed_values",
						Description: "Allowed values of the configuration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "source",
						Description: "Source of the configuration.",
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

func fetchMariadbServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().MariaDB.Servers
	r, err := svc.List(ctx)
	if err != nil {
		return err
	}
	if r.Value == nil {
		return nil
	}
	res <- *r.Value
	return nil
}

func resolveServerPrivateEndpointConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	s := parent.Item.(mariadb.Server)
	if s.ServerProperties == nil || s.ServerProperties.PrivateEndpointConnections == nil {
		return nil
	}
	res <- *s.ServerProperties.PrivateEndpointConnections
	return nil
}

func resolveMariadbServerConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	s := parent.Item.(mariadb.Server)
	svc := meta.(*client.Client).Services().MariaDB.Configurations
	resourceDetails, err := client.ParseResourceID(*s.ID)
	if err != nil {
		return err
	}
	r, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *s.Name)
	if err != nil {
		return err
	}
	if r.Value == nil {
		return nil
	}
	res <- *r.Value
	return nil
}
