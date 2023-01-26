# Table: hubspot_crm_pipelines

https://developers.hubspot.com/docs/api/crm/pipelines

The composite primary key for this table is (**object_type**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|object_type (PK)|String|
|id (PK)|String|
|label|String|
|display_order|Int|
|stages|JSON|
|created_at|Timestamp|
|archived_at|Timestamp|
|updated_at|Timestamp|
|archived|Bool|