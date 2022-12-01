# Table: aws_fsx_file_systems

https://docs.aws.amazon.com/fsx/latest/APIReference/API_FileSystem.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|administrative_actions|JSON|
|creation_time|Timestamp|
|dns_name|String|
|failure_details|JSON|
|file_system_id|String|
|file_system_type|String|
|file_system_type_version|String|
|kms_key_id|String|
|lifecycle|String|
|lustre_configuration|JSON|
|network_interface_ids|StringArray|
|ontap_configuration|JSON|
|open_zfs_configuration|JSON|
|owner_id|String|
|storage_capacity|Int|
|storage_type|String|
|subnet_ids|StringArray|
|vpc_id|String|
|windows_configuration|JSON|