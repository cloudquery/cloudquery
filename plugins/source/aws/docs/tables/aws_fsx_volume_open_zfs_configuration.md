
# Table: aws_fsx_volume_open_zfs_configuration
The configuration of an Amazon FSx for OpenZFS volume
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|volume_cq_id|uuid|Unique CloudQuery ID of aws_fsx_volumes table (FK)|
|copy_tags_to_snapshots|boolean|A Boolean value indicating whether tags for the volume should be copied to snapshots|
|data_compression_type|text|Specifies the method used to compress the data on the volume|
|nfs_exports|jsonb|The configuration object for mounting a Network File System (NFS) file system|
|origin_snapshot_copy_strategy|text|The strategy used when copying data from the snapshot to the new volume|
|origin_snapshot_arn|text|The Amazon Resource Name (ARN) for a given resource|
|parent_volume_id|text|The ID of the parent volume|
|read_only|boolean|A Boolean value indicating whether the volume is read-only|
|record_size|bigint|The record size of an OpenZFS volume, in kibibytes (KiB)|
|storage_capacity_quota|bigint|The maximum amount of storage in gibibtyes (GiB) that the volume can use from its parent|
|storage_capacity_reservation|bigint|The amount of storage in gibibytes (GiB) to reserve from the parent volume|
|user_and_group_quotas|jsonb|An object specifying how much storage users or groups can use on the volume|
|volume_path|text|The path to the volume from the root volume|
