# Table: alicloud_oss_bucket_stats

This table shows data for Alibaba Cloud Object Storage Service (OSS) Bucket Stats.

The composite primary key for this table is (**bucket_name**, **account_id**).

## Relations

This table depends on [alicloud_oss_buckets](alicloud_oss_buckets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|bucket_name (PK)|`utf8`|
|account_id (PK)|`utf8`|
|xml_name|`json`|
|storage|`int64`|
|object_count|`int64`|
|multipart_upload_count|`int64`|
|live_channel_count|`int64`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|standard_storage|`int64`|
|standard_object_count|`int64`|
|infrequent_access_storage|`int64`|
|infrequent_access_real_storage|`int64`|
|infrequent_access_object_count|`int64`|
|archive_storage|`int64`|
|archive_real_storage|`int64`|
|archive_object_count|`int64`|
|cold_archive_storage|`int64`|
|cold_archive_real_storage|`int64`|
|cold_archive_object_count|`int64`|