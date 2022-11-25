# Table: aws_fsx_file_caches

https://docs.aws.amazon.com/fsx/latest/APIReference/API_FileCache.html

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
|creation_time|Timestamp|
|dns_name|String|
|data_repository_association_ids|StringArray|
|failure_details|JSON|
|file_cache_id|String|
|file_cache_type|String|
|file_cache_type_version|String|
|kms_key_id|String|
|lifecycle|String|
|lustre_configuration|JSON|
|network_interface_ids|StringArray|
|owner_id|String|
|storage_capacity|Int|
|subnet_ids|StringArray|
|vpc_id|String|