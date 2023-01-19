# Table: alicloud_oss_bucket_stats

The composite primary key for this table is (**bucket_name**, **account_id**).

## Relations

This table depends on [alicloud_oss_buckets](alicloud_oss_buckets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|bucket_name (PK)|String|
|account_id (PK)|String|
|xml_name|JSON|
|storage|Int|
|object_count|Int|
|multipart_upload_count|Int|
|live_channel_count|Int|
|last_modified_time|Timestamp|
|standard_storage|Int|
|standard_object_count|Int|
|infrequent_access_storage|Int|
|infrequent_access_real_storage|Int|
|infrequent_access_object_count|Int|
|archive_storage|Int|
|archive_real_storage|Int|
|archive_object_count|Int|
|cold_archive_storage|Int|
|cold_archive_real_storage|Int|
|cold_archive_object_count|Int|