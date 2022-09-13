
# Table: gcp_compute_instance_disks
An instance-attached disk resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique ID of gcp_compute_instances table (FK)|
|instance_id|text||
|auto_delete|boolean|Specifies whether the disk will be auto-deleted when the instance is deleted (but not when the disk is detached from the instance)|
|boot|boolean|Indicates that this is a boot disk The virtual machine will use the first partition of the disk for its root filesystem|
|device_name|text|Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux operating system running within the instance This name can be used to reference the device for mounting, resizing, and so on, from within the instance  If not specified, the server chooses a default device name to apply to this disk, in the form persistent-disk-x, where x is a number assigned by Google Compute Engine This field is only applicable for persistent disks|
|disk_encryption_key_kms_key_name|text|The name of the encryption key that is stored in Google Cloud KMS|
|disk_encryption_key_kms_key_service_account|text|The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used|
|disk_encryption_key_raw_key|text|Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource|
|disk_encryption_key_sha256|text|[Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource|
|disk_size_gb|bigint|The size of the disk in GB|
|guest_os_features|text[]|A list of features to enable on the guest operating system Applicable only for bootable images Read  Enabling guest operating system features to see a list of available options|
|index|bigint|A zero-based index to this disk, where 0 is reserved for the boot disk If you have many disks attached to an instance, each disk would have a unique index number|
|description|text|An optional description Provide this property when creating the disk|
|disk_name|text|Specifies the disk name If not specified, the default is to use the name of the instance If a disk with the same name already exists in the given region, the existing disk is attached to the new instance and the new disk is not created|
|initialized_disk_size_gb|bigint|Specifies the size of the disk in base-2 GB The size must be at least 10 GB If you specify a sourceImage, which is required for boot disks, the default size is the size of the sourceImage If you do not specify a sourceImage, the default disk size is 500 GB|
|disk_type|text|Specifies the disk type to use to create the instance|
|labels|jsonb|Labels to apply to this disk These can be later modified by the diskssetLabels method This field is only applicable for persistent disks|
|on_update_action|text|Specifies which action to take on instance update with this disk Default is to use the existing disk|
|provisioned_iops|bigint|Indicates how many IOPS must be provisioned for the disk|
|resource_policies|text[]|Resource policies applied to this disk for automatic snapshot creations Specified using the full or partial URL For instance template, specify only the resource policy name|
|source_image|text|The source image to create this disk|
|source_image_encryption_key_kms_key_name|text|The name of the encryption key that is stored in Google Cloud KMS|
|source_image_encryption_key_kms_key_service_account|text|The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used|
|source_image_encryption_key_raw_key|text|Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource|
|source_image_encryption_key_sha256|text|The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource|
|source_snapshot|text|The source snapshot to create this disk|
|source_snapshot_encryption_key_kms_key_name|text|The name of the encryption key that is stored in Google Cloud KMS|
|source_snapshot_encryption_key_kms_key_service_account|text|The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used|
|source_snapshot_encryption_key_raw_key|text|Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource|
|source_snapshot_encryption_key_sha256|text|[Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource|
|interface|text|Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME The default is SCSI Persistent disks must always use SCSI and the request will fail if you attempt to attach a persistent disk in any other format than SCSI Local SSDs can use either NVME or SCSI For performance characteristics of SCSI over NVMe, see Local SSD performance|
|kind|text|Type of the resource Always compute#attachedDisk for attached disks|
|licenses|text[]|Any valid publicly visible licenses|
|mode|text|The mode in which to attach this disk, either READ_WRITE or READ_ONLY If not specified, the default is to attach the disk in READ_WRITE mode|
|shielded_instance_initial_state_pk_content|text|The raw content in the secure keys file|
|shielded_instance_initial_state_pk_file_type|text|The file type of source file|
|source|text|The source snapshot to create this disk|
|type|text|Specifies the type of the disk, either SCRATCH or PERSISTENT If not specified, the default is PERSISTENT|
