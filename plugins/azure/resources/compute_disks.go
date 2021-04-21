package resources

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ComputeDisks() *schema.Table {
	return &schema.Table{
		Name:         "azure_compute_disks",
		Resolver:     fetchComputeDisks,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name: "managed_by",
				Type: schema.TypeString,
			},
			{
				Name: "managed_by_extended",
				Type: schema.TypeStringArray,
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
				Name: "zones",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "time_created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DiskProperties.TimeCreated.Time"),
			},
			{
				Name:     "os_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.OsType"),
			},
			{
				Name:     "hyperv_generation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.HyperVGeneration"),
			},
			{
				Name:     "creation_data_create_option",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.CreationData.CreateOption"),
			},
			{
				Name:     "creation_data_storage_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.CreationData.StorageAccountID"),
			},
			{
				Name:     "creation_data_image_reference_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.CreationData.ImageReference.ID"),
			},
			{
				Name:     "creation_data_image_reference_lun",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DiskProperties.CreationData.ImageReference.Lun"),
			},
			{
				Name:     "creation_data_gallery_image_reference_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.CreationData.GalleryImageReference.ID"),
			},
			{
				Name:     "creation_data_gallery_image_reference_lun",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DiskProperties.CreationData.GalleryImageReference.Lun"),
			},
			{
				Name:     "creation_data_source_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.CreationData.SourceURI"),
			},
			{
				Name:     "creation_data_source_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.CreationData.SourceResourceID"),
			},
			{
				Name:     "creation_data_source_unique_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.CreationData.SourceUniqueID"),
			},
			{
				Name:     "creation_data_upload_size_bytes",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DiskProperties.CreationData.UploadSizeBytes"),
			},
			{
				Name:     "disk_size_gb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DiskProperties.DiskSizeGB"),
			},
			{
				Name:     "disk_size_bytes",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DiskProperties.DiskSizeBytes"),
			},
			{
				Name:     "unique_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.UniqueID"),
			},
			{
				Name:     "encryption_settings_collection_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DiskProperties.EncryptionSettingsCollection.Enabled"),
			},
			{
				Name:     "encryption_settings_collection_encryption_settings_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.EncryptionSettingsCollection.EncryptionSettingsVersion"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.ProvisioningState"),
			},
			{
				Name:     "disk_iops_read_write",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DiskProperties.DiskIOPSReadWrite"),
			},
			{
				Name:     "disk_mbps_read_write",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DiskProperties.DiskMBpsReadWrite"),
			},
			{
				Name:     "disk_iops_read_only",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DiskProperties.DiskIOPSReadOnly"),
			},
			{
				Name:     "disk_mbps_read_only",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DiskProperties.DiskMBpsReadOnly"),
			},
			{
				Name:     "disk_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.DiskState"),
			},
			{
				Name:     "encryption_disk_encryption_set_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.Encryption.DiskEncryptionSetID"),
			},
			{
				Name:     "encryption_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.Encryption.Type"),
			},
			{
				Name:     "max_shares",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DiskProperties.MaxShares"),
			},
			{
				Name:     "share_info",
				Type:     schema.TypeStringArray,
				Resolver: resolveComputeDiskShareInfo,
			},
			{
				Name:     "network_access_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.NetworkAccessPolicy"),
			},
			{
				Name:     "disk_access_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DiskProperties.DiskAccessID"),
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
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "azure_compute_disk_encryption_settings",
				Resolver: fetchComputeDiskEncryptionSettings,
				Columns: []schema.Column{
					{
						Name:     "disk_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "disk_encryption_key_source_vault_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DiskEncryptionKey.SourceVault.ID"),
					},
					{
						Name:     "disk_encryption_key_secret_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DiskEncryptionKey.SecretURL"),
					},
					{
						Name:     "key_encryption_key_source_vault_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("KeyEncryptionKey.SourceVault.ID"),
					},
					{
						Name:     "key_encryption_key_key_url",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("KeyEncryptionKey.KeyURL"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeDisks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Compute.Disks
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

func resolveComputeDiskShareInfo(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	disk := resource.Item.(compute.Disk)
	if disk.ShareInfo == nil {
		return nil
	}
	shareInfo := make([]*string, len(*disk.ShareInfo))
	for i, info := range *disk.ShareInfo {
		shareInfo[i] = info.VMURI
	}
	resource.Set("share_info", shareInfo)
	return nil
}

func fetchComputeDiskEncryptionSettings(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	disk := parent.Item.(compute.Disk)
	if disk.EncryptionSettingsCollection == nil || disk.EncryptionSettingsCollection.EncryptionSettings == nil {
		return nil
	}
	for _, e := range *disk.EncryptionSettingsCollection.EncryptionSettings {
		res <- e
	}
	return nil
}
