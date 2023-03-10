# Table: azure_security_assessments_metadata

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/assessments-metadata/list?tabs=HTTP#securityassessmentmetadata

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id (PK)|String|
|properties|JSON|
|id (PK)|String|
|name|String|
|type|String|