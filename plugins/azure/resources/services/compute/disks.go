package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ComputeDisks() *schema.Table {
	return &schema.Table{
		Name:         "azure_compute_disks",
		Description:  "Azure compute disk",
		Resolver:     fetchComputeDisks,
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
				Name:        "managed_by",
				Description: "A relative URI containing the ID of the VM that has the disk attached",
				Type:        schema.TypeString,
			},
			{
				Name:          "managed_by_extended",
				Description:   "List of relative URIs containing the IDs of the VMs that have the disk attached maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "sku_name",
				Description: "The sku name Possible values include: 'StandardLRS', 'PremiumLRS', 'StandardSSDLRS', 'UltraSSDLRS'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The sku tier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:          "zones",
				Description:   "The Logical zone list for Disk",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "time_created",
				Description: "The time when the disk was created",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DiskProperties.TimeCreated.Time"),
			},
			{
				Name:        "os_type",
				Description: "The Operating System type Possible values include: 'Windows', 'Linux'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskProperties.OsType"),
			},
			{
				Name:        "hyperv_generation",
				Description: "The hypervisor generation of the Virtual Machine Applicable to OS disks only Possible values include: 'V1', 'V2'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskProperties.HyperVGeneration"),
			},
			{
				Name:        "creation_data_create_option",
				Description: "This enumerates the possible sources of a disk's creation Possible values include: 'Empty', 'Attach', 'FromImage', 'Import', 'Copy', 'Restore', 'Upload'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskProperties.CreationData.CreateOption"),
			},
			{
				Name:          "creation_data_storage_account_id",
				Description:   "Required if createOption is Import The Azure Resource Manager identifier of the storage account containing the blob to import as a disk",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DiskProperties.CreationData.StorageAccountID"),
				IgnoreInTests: true,
			},
			{
				Name:        "creation_data_image_reference_id",
				Description: "A relative uri containing either a Platform Image Repository or user image reference",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskProperties.CreationData.ImageReference.ID"),
			},
			{
				Name:          "creation_data_image_reference_lun",
				Description:   "If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use For OS disks, this field is null",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("DiskProperties.CreationData.ImageReference.Lun"),
				IgnoreInTests: true,
			},
			{
				Name:          "creation_data_gallery_image_reference_id",
				Description:   "A relative uri containing either a Platform Image Repository or user image reference",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DiskProperties.CreationData.GalleryImageReference.ID"),
				IgnoreInTests: true,
			},
			{
				Name:          "creation_data_gallery_image_reference_lun",
				Description:   "If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use For OS disks, this field is null",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("DiskProperties.CreationData.GalleryImageReference.Lun"),
				IgnoreInTests: true,
			},
			{
				Name:          "creation_data_source_uri",
				Description:   "If createOption is Import, this is the URI of a blob to be imported into a managed disk",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DiskProperties.CreationData.SourceURI"),
				IgnoreInTests: true,
			},
			{
				Name:          "creation_data_source_resource_id",
				Description:   "If createOption is Copy, this is the ARM id of the source snapshot or disk",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DiskProperties.CreationData.SourceResourceID"),
				IgnoreInTests: true,
			},
			{
				Name:          "creation_data_source_unique_id",
				Description:   "If this field is set, this is the unique id identifying the source of this resource",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DiskProperties.CreationData.SourceUniqueID"),
				IgnoreInTests: true,
			},
			{
				Name:          "creation_data_upload_size_bytes",
				Description:   "If createOption is Upload, this is the size of the contents of the upload including the VHD footer This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512 bytes for the VHD footer)",
				Type:          schema.TypeBigInt,
				Resolver:      schema.PathResolver("DiskProperties.CreationData.UploadSizeBytes"),
				IgnoreInTests: true,
			},
			{
				Name:        "disk_size_gb",
				Description: "If creationDatacreateOption is Empty, this field is mandatory and it indicates the size of the disk to create If this field is present for updates or creation with other options, it indicates a resize Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DiskProperties.DiskSizeGB"),
			},
			{
				Name:        "disk_size_bytes",
				Description: "The size of the disk in bytes This field is read only",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DiskProperties.DiskSizeBytes"),
			},
			{
				Name:        "unique_id",
				Description: "Unique Guid identifying the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskProperties.UniqueID"),
			},
			{
				Name:        "encryption_settings_collection_enabled",
				Description: "Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption Set this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption If EncryptionSettings is null in the request object, the existing settings remain unchanged",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DiskProperties.EncryptionSettingsCollection.Enabled"),
			},
			{
				Name:        "encryption_settings_collection_encryption_settings_version",
				Description: "Describes what type of encryption is used for the disks Once this field is set, it cannot be overwritten '10' corresponds to Azure Disk Encryption with AAD app'11' corresponds to Azure Disk Encryption",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskProperties.EncryptionSettingsCollection.EncryptionSettingsVersion"),
			},
			{
				Name:        "provisioning_state",
				Description: "The disk provisioning state",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskProperties.ProvisioningState"),
			},
			{
				Name:        "disk_iops_read_write",
				Description: "only settable for UltraSSD disks One operation can transfer between 4k and 256k bytes",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DiskProperties.DiskIOPSReadWrite"),
			},
			{
				Name:        "disk_mbps_read_write",
				Description: "only settable for UltraSSD disks MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DiskProperties.DiskMBpsReadWrite"),
			},
			{
				Name:          "disk_iops_read_only",
				Description:   "The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly One operation can transfer between 4k and 256k bytes",
				Type:          schema.TypeBigInt,
				Resolver:      schema.PathResolver("DiskProperties.DiskIOPSReadOnly"),
				IgnoreInTests: true,
			},
			{
				Name:          "disk_mbps_read_only",
				Description:   "The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10",
				Type:          schema.TypeBigInt,
				Resolver:      schema.PathResolver("DiskProperties.DiskMBpsReadOnly"),
				IgnoreInTests: true,
			},
			{
				Name:        "disk_state",
				Description: "The state of the disk Possible values include: 'Unattached', 'Attached', 'Reserved', 'ActiveSAS', 'ReadyToUpload', 'ActiveUpload'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskProperties.DiskState"),
			},
			{
				Name:          "encryption_disk_encryption_set_id",
				Description:   "ResourceId of the disk encryption set to use for enabling encryption at rest",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DiskProperties.Encryption.DiskEncryptionSetID"),
				IgnoreInTests: true,
			},
			{
				Name:        "encryption_type",
				Description: "Possible values include: 'EncryptionAtRestWithPlatformKey', 'EncryptionAtRestWithCustomerKey', 'EncryptionAtRestWithPlatformAndCustomerKeys'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskProperties.Encryption.Type"),
			},
			{
				Name:          "max_shares",
				Description:   "The maximum number of VMs that can attach to the disk at the same time Value greater than one indicates a disk that can be mounted on multiple VMs at the same time",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("DiskProperties.MaxShares"),
				IgnoreInTests: true,
			},
			{
				Name:          "share_info",
				Description:   "Details of the list of all VMs that have the disk attached maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs",
				Type:          schema.TypeStringArray,
				Resolver:      resolveComputeDiskShareInfo,
				IgnoreInTests: true,
			},
			{
				Name:        "network_access_policy",
				Description: "Possible values include: 'AllowAll', 'AllowPrivate', 'DenyAll'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiskProperties.NetworkAccessPolicy"),
			},
			{
				Name:          "disk_access_id",
				Description:   "ARM id of the DiskAccess resource for using private endpoints on disks",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DiskProperties.DiskAccessID"),
				IgnoreInTests: true,
			},
			{
				Name:        "id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_compute_disk_encryption_settings",
				Description: "Azure compute disk encryption setting",
				Resolver:    fetchComputeDiskEncryptionSettings,
				Columns: []schema.Column{
					{
						Name:        "disk_cq_id",
						Description: "Unique ID of azure_compute_disks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "disk_encryption_key_source_vault_id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DiskEncryptionKey.SourceVault.ID"),
					},
					{
						Name:        "disk_encryption_key_secret_url",
						Description: "Url pointing to a key or secret in KeyVault",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DiskEncryptionKey.SecretURL"),
					},
					{
						Name:          "key_encryption_key_source_vault_id",
						Description:   "Resource Id",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("KeyEncryptionKey.SourceVault.ID"),
						IgnoreInTests: true,
					},
					{
						Name:          "key_encryption_key_key_url",
						Description:   "Url pointing to a key or secret in KeyVault",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("KeyEncryptionKey.KeyURL"),
						IgnoreInTests: true,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeDisks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
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
	return resource.Set("share_info", shareInfo)
}

func fetchComputeDiskEncryptionSettings(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	disk := parent.Item.(compute.Disk)
	if disk.EncryptionSettingsCollection == nil || disk.EncryptionSettingsCollection.EncryptionSettings == nil {
		return nil
	}
	for _, e := range *disk.EncryptionSettingsCollection.EncryptionSettings {
		res <- e
	}
	return nil
}
