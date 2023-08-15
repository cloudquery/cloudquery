# Table: azure_security_assessments_metadata

This table shows data for Azure Security Assessments Metadata.

https://learn.microsoft.com/en-us/rest/api/defenderforcloud/assessments-metadata/list?tabs=HTTP#securityassessmentmetadata

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id (PK)|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|