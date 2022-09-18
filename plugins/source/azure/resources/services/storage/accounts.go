// Auto generated code - DO NOT EDIT.

package storage

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:      "azure_storage_accounts",
		Resolver:  fetchStorageAccounts,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "extended_location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExtendedLocation"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "primary_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrimaryEndpoints"),
			},
			{
				Name:     "primary_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrimaryLocation"),
			},
			{
				Name:     "status_of_primary",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusOfPrimary"),
			},
			{
				Name:     "last_geo_failover_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastGeoFailoverTime"),
			},
			{
				Name:     "secondary_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecondaryLocation"),
			},
			{
				Name:     "status_of_secondary",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatusOfSecondary"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "custom_domain",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CustomDomain"),
			},
			{
				Name:     "secondary_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecondaryEndpoints"),
			},
			{
				Name:     "encryption",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Encryption"),
			},
			{
				Name:     "access_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessTier"),
			},
			{
				Name:     "azure_files_identity_based_authentication",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AzureFilesIdentityBasedAuthentication"),
			},
			{
				Name:     "enable_https_traffic_only",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableHTTPSTrafficOnly"),
			},
			{
				Name:     "network_rule_set",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkRuleSet"),
			},
			{
				Name:     "is_hns_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsHnsEnabled"),
			},
			{
				Name:     "geo_replication_stats",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GeoReplicationStats"),
			},
			{
				Name:     "failover_in_progress",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("FailoverInProgress"),
			},
			{
				Name:     "large_file_shares_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LargeFileSharesState"),
			},
			{
				Name:     "private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateEndpointConnections"),
			},
			{
				Name:     "routing_preference",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RoutingPreference"),
			},
			{
				Name:     "blob_restore_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BlobRestoreStatus"),
			},
			{
				Name:     "allow_blob_public_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AllowBlobPublicAccess"),
			},
			{
				Name:     "minimum_tls_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MinimumTLSVersion"),
			},
			{
				Name:     "allow_shared_key_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AllowSharedKeyAccess"),
			},
			{
				Name:     "enable_nfs_v_3",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableNfsV3"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
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
				Name:     "blob_logging_settings",
				Type:     schema.TypeJSON,
				Resolver: fetchStorageAccountBlobLoggingSettings,
			},
			{
				Name:     "queue_logging_settings",
				Type:     schema.TypeJSON,
				Resolver: fetchStorageAccountQueueLoggingSettings,
			},
		},

		Relations: []*schema.Table{
			containers(),
			blobServices(),
		},
	}
}

func fetchStorageAccountBlobLoggingSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	acc := resource.Item.(storage.Account)

	if !isBlobSupported(&acc) {
		return nil
	}

	storageClient := meta.(*client.Client).Services().Storage.Accounts
	details, err := client.ParseResourceID(*acc.ID)
	if err != nil {
		return err
	}
	result, err := storageClient.GetBlobServiceProperties(ctx, details.ResourceGroup, *acc.Name)
	if err != nil {
		return err
	}
	if result.StorageServiceProperties != nil {
		return resource.Set(c.Name, result.StorageServiceProperties.Logging)
	}
	return nil
}

func fetchStorageAccountQueueLoggingSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	acc := resource.Item.(storage.Account)
	if !isQueueSupported(&acc) {
		return nil
	}

	storageClient := meta.(*client.Client).Services().Storage.Accounts
	details, err := client.ParseResourceID(*acc.ID)
	if err != nil {
		return err
	}
	result, err := storageClient.GetQueueServiceProperties(ctx, details.ResourceGroup, *acc.Name)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, result.Logging)
}

// isQueueSupported checks whether queues are supported for a storage account.
// Premium storage accounts don't support queues.
func isQueueSupported(account *storage.Account) bool {
	return account.Sku.Tier == storage.Standard && account.Kind == storage.StorageV2
}

// isBlobSupported checks whether blobs are supported for a storage account.
func isBlobSupported(account *storage.Account) bool {
	return (account.Kind == storage.Storage) || (account.Kind == storage.StorageV2) ||
		(account.Kind == storage.BlockBlobStorage) || (account.Kind == storage.BlobStorage)
}

func fetchStorageAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Storage.Accounts

	response, err := svc.List(ctx)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
