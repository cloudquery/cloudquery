# Table: alicloud_oss_buckets

The primary key for this table is **name**.

## Relations

The following tables depend on alicloud_oss_buckets:
  - [alicloud_oss_bucket_stats](alicloud_oss_bucket_stats.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|xml_name|JSON|
|name (PK)|String|
|location|String|
|creation_date|Timestamp|
|storage_class|String|