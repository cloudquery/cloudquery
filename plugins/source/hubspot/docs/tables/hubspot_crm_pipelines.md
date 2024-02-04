# Table: hubspot_crm_pipelines

This table shows data for HubSpot CRM Pipelines.

https://developers.hubspot.com/docs/api/crm/pipelines

The composite primary key for this table is (**object_type**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|object_type (PK)|`utf8`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|archived_at|`timestamp[us, tz=UTC]`|
|archived|`bool`|
|display_order|`int64`|
|stages|`json`|
|label|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|