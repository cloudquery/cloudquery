
# Table: azure_compute_disks
Azure compute disk
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|managed_by|text|A relative URI containing the ID of the VM that has the disk attached|
|managed_by_extended|text[]|List of relative URIs containing the IDs of the VMs that have the disk attached maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs|
|sku_name|text|The sku name Possible values include: 'StandardLRS', 'PremiumLRS', 'StandardSSDLRS', 'UltraSSDLRS'|
|sku_tier|text|The sku tier|
|zones|text[]|The Logical zone list for Disk|
|time_created|timestamp without time zone|The time when the disk was created|
|os_type|text|The Operating System type Possible values include: 'Windows', 'Linux'|
|hyperv_generation|text|The hypervisor generation of the Virtual Machine Applicable to OS disks only Possible values include: 'V1', 'V2'|
|creation_data_create_option|text|This enumerates the possible sources of a disk's creation Possible values include: 'Empty', 'Attach', 'FromImage', 'Import', 'Copy', 'Restore', 'Upload'|
|creation_data_storage_account_id|text|Required if createOption is Import The Azure Resource Manager identifier of the storage account containing the blob to import as a disk|
|creation_data_image_reference_id|text|A relative uri containing either a Platform Image Repository or user image reference|
|creation_data_image_reference_lun|integer|If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use For OS disks, this field is null|
|creation_data_gallery_image_reference_id|text|A relative uri containing either a Platform Image Repository or user image reference|
|creation_data_gallery_image_reference_lun|integer|If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the image to use For OS disks, this field is null|
|creation_data_source_uri|text|If createOption is Import, this is the URI of a blob to be imported into a managed disk|
|creation_data_source_resource_id|text|If createOption is Copy, this is the ARM id of the source snapshot or disk|
|creation_data_source_unique_id|text|If this field is set, this is the unique id identifying the source of this resource|
|creation_data_upload_size_bytes|bigint|If createOption is Upload, this is the size of the contents of the upload including the VHD footer This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512 bytes for the VHD footer)|
|disk_size_gb|integer|If creationDatacreateOption is Empty, this field is mandatory and it indicates the size of the disk to create If this field is present for updates or creation with other options, it indicates a resize Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size|
|disk_size_bytes|bigint|The size of the disk in bytes This field is read only|
|unique_id|text|Unique Guid identifying the resource|
|encryption_settings_collection_enabled|boolean|Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption Set this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption If EncryptionSettings is null in the request object, the existing settings remain unchanged|
|encryption_settings_collection_encryption_settings_version|text|Describes what type of encryption is used for the disks Once this field is set, it cannot be overwritten '10' corresponds to Azure Disk Encryption with AAD app'11' corresponds to Azure Disk Encryption|
|provisioning_state|text|The disk provisioning state|
|disk_iops_read_write|bigint|only settable for UltraSSD disks One operation can transfer between 4k and 256k bytes|
|disk_mbps_read_write|bigint|only settable for UltraSSD disks MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10|
|disk_iops_read_only|bigint|The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly One operation can transfer between 4k and 256k bytes|
|disk_mbps_read_only|bigint|The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10|
|disk_state|text|The state of the disk Possible values include: 'Unattached', 'Attached', 'Reserved', 'ActiveSAS', 'ReadyToUpload', 'ActiveUpload'|
|encryption_disk_encryption_set_id|text|ResourceId of the disk encryption set to use for enabling encryption at rest|
|encryption_type|text|Possible values include: 'EncryptionAtRestWithPlatformKey', 'EncryptionAtRestWithCustomerKey', 'EncryptionAtRestWithPlatformAndCustomerKeys'|
|max_shares|integer|The maximum number of VMs that can attach to the disk at the same time Value greater than one indicates a disk that can be mounted on multiple VMs at the same time|
|share_info|text[]|Details of the list of all VMs that have the disk attached maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs|
|network_access_policy|text|Possible values include: 'AllowAll', 'AllowPrivate', 'DenyAll'|
|disk_access_id|text|ARM id of the DiskAccess resource for using private endpoints on disks|
|id|text|Resource Id|
|name|text|Resource name|
|type|text|Resource type|
|location|text|Resource location|
|tags|jsonb|Resource tags|
