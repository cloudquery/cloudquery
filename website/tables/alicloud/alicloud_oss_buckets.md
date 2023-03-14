# Table: alicloud_oss_buckets

This table shows data for Alicloud Oss Buckets.

The composite primary key for this table is (**account_id**, **name**).

## Relations

The following tables depend on alicloud_oss_buckets:
  - [alicloud_oss_bucket_stats](alicloud_oss_bucket_stats)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|xml_name|JSON|
|name (PK)|String|
|location|String|
|creation_date|Timestamp|
|storage_class|String|