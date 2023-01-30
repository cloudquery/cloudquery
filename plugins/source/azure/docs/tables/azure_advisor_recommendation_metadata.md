# Table: azure_advisor_recommendation_metadata

https://learn.microsoft.com/en-us/rest/api/advisor/recommendation-metadata/list?tabs=HTTP#metadataentity

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id (PK)|String|
|id (PK)|String|
|name|String|
|properties|JSON|
|type|String|