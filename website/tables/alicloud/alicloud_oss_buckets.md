# Table: alicloud_oss_buckets

This table shows data for Alibaba Cloud Object Storage Service (OSS) Buckets.

The composite primary key for this table is (**account_id**, **name**).

## Relations

The following tables depend on alicloud_oss_buckets:
  - [alicloud_oss_bucket_stats](alicloud_oss_bucket_stats)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|xml_name|`json`|
|name (PK)|`utf8`|
|location|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|storage_class|`utf8`|