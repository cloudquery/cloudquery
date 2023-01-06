# Table: alicloud_oss_bucket_stats



The composite primary key for this table is (**name**, **update_date**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|name (PK)|String|
|update_date (PK)|String|
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