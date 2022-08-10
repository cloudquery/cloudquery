
# Table: gcp_compute_images
Represents an Image resource  You can use images to create boot disks for your VM instances
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|archive_size_bytes|bigint|Size of the image targz archive stored in Google Cloud Storage (in bytes)|
|creation_timestamp|text|Creation timestamp in RFC3339 text format|
|deprecated_deleted|text|An optional RFC3339 timestamp on or after which the state of this resource is intended to change to DELETED This is only informational and the status will not change unless the client explicitly changes it|
|deprecated|text||
|deprecated_obsolete|text|An optional RFC3339 timestamp on or after which the state of this resource is intended to change to OBSOLETE This is only informational and the status will not change unless the client explicitly changes it|
|deprecated_replacement|text|The URL of the suggested replacement for a deprecated resource The suggested replacement resource must be the same kind of resource as the deprecated resource|
|deprecated_state|text|The deprecation state of this resource This can be ACTIVE, DEPRECATED, OBSOLETE, or DELETED Operations which communicate the end of life date for an image, can use ACTIVE Operations which create a new resource using a DEPRECATED resource will return successfully, but with a warning indicating the deprecated resource and recommending its replacement Operations which use OBSOLETE or DELETED resources will be rejected and result in an error|
|description|text|An optional description of this resource Provide this property when you create the resource|
|disk_size_gb|bigint|Size of the image when restored onto a persistent disk (in GB)|
|family|text|The name of the image family to which this image belongs You can create disks by specifying an image family instead of a specific image name The image family always returns its latest image that is not deprecated The name of the image family must comply with RFC1035|
|guest_os_features|text[]|A list of features to enable on the guest operating system Applicable only for bootable images Read  Enabling guest operating system features to see a list of available options|
|id|text|The unique identifier for the resource This identifier is defined by the server|
|image_encryption_key_kms_key_name|text|The name of the encryption key that is stored in Google Cloud KMS|
|image_encryption_key_kms_key_service_account|text|The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used|
|image_encryption_key_raw_key|text|Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource|
|image_encryption_key_sha256|text|[Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource|
|kind|text|Type of the resource Always compute#image for images|
|label_fingerprint|text|A fingerprint for the labels being applied to this image|
|labels|jsonb|Labels for this resource|
|licenses|text[]|Any applicable license URI|
|name|text|Name of the resource|
|raw_disk_container_type|text|The format used to encode and transmit the block device, which should be TAR This is just a container and transmission format and not a runtime format Provided by the client when the disk image is created|
|raw_disk_source|text|The full Google Cloud Storage URL where the disk image is stored You must provide either this property or the sourceDisk property but not both|
|satisfies_pzs|boolean|Reserved for future use|
|self_link|text|Server-defined URL for the resource|
|shielded_instance_initial_state_pk_content|text|The raw content in the secure keys file|
|shielded_instance_initial_state_pk_file_type|text|The file type of source file|
|source_disk|text|URL of the source disk used to create this image This can be a full or valid partial URL You must provide either this property or the rawDisksource property but not both to create an image For example, the following are valid values: - https://wwwgoogleapis|
|source_disk_encryption_key_kms_key_name|text|The name of the encryption key that is stored in Google Cloud KMS|
|source_disk_encryption_key_kms_key_service_account|text|The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used|
|source_disk_encryption_key_raw_key|text|Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource|
|source_disk_encryption_key_sha256|text|[Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource|
|source_disk_id|text|The ID value of the disk used to create this image This value may be used to determine whether the image was taken from the current or a previous instance of a given disk name|
|source_image|text|URL of the source image used to create this image  In order to create an image, you must provide the full or partial URL of one of the following: - The selfLink URL - This property - The rawDisk|
|source_image_encryption_key_kms_key_name|text|The name of the encryption key that is stored in Google Cloud KMS|
|source_image_encryption_key_kms_key_service_account|text|The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used|
|source_image_encryption_key_raw_key|text|Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource|
|source_image_encryption_key_sha256|text|[Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource|
|source_image_id|text|The ID value of the image used to create this image This value may be used to determine whether the image was taken from the current or a previous instance of a given image name|
|source_snapshot|text|URL of the source snapshot used to create this image  In order to create an image, you must provide the full or partial URL of one of the following: - The selfLink URL - This property - The sourceImage URL - The rawDisk|
|source_snapshot_encryption_key_kms_key_name|text|The name of the encryption key that is stored in Google Cloud KMS|
|source_snapshot_encryption_key_kms_key_service_account|text|The service account being used for the encryption request for the given KMS key If absent, the Compute Engine default service account is used|
|source_snapshot_encryption_key_raw_key|text|Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource|
|source_snapshot_encryption_key_sha256|text|[Output only] The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource|
|source_snapshot_id|text|The ID value of the snapshot used to create this image This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given snapshot name|
|source_type|text|The type of the image used to create this disk|
|status|text|The status of the image An image can be used to create other resources, such as instances, only after the image has been successfully created and the status is set to READY Possible values are FAILED, PENDING, or READY|
|storage_locations|text[]|Cloud Storage bucket storage location of the image (regional or multi-regional)|
