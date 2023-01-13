# Table: alicloud_oss_bucket_stats

The primary key for this table is **_cq_id**.

## Relations

This table depends on [alicloud_oss_buckets](alicloud_oss_buckets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|xml_name|JSON|
|storage|Int|
|object_count|Int|
|multipart_upload_count|Int|
|live_channel_count|Int|
|last_modified_time|Int|
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