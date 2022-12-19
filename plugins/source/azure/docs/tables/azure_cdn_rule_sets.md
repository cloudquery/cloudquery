# Table: azure_cdn_rule_sets

The primary key for this table is **id**.

## Relations

This table depends on [azure_cdn_profiles](azure_cdn_profiles.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|properties|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|