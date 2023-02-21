# Table: oracle_identity_network_sources

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|compartment_id|String|
|name|String|
|description|String|
|public_source_list|StringArray|
|virtual_source_list|JSON|
|services|StringArray|
|time_created|Timestamp|
|freeform_tags|JSON|
|defined_tags|JSON|