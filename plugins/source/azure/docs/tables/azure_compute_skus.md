# Table: azure_compute_skus

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|api_versions|StringArray|
|capabilities|JSON|
|capacity|JSON|
|costs|JSON|
|family|String|
|kind|String|
|location_info|JSON|
|locations|StringArray|
|name|String|
|resource_type|String|
|restrictions|JSON|
|size|String|
|tier|String|