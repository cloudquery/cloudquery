# Table: azure_advisor_recommendation_metadata

This table shows data for Azure Advisor Recommendation Metadata.

https://learn.microsoft.com/en-us/rest/api/advisor/recommendation-metadata/list?tabs=HTTP#metadataentity

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|properties|`json`|
|type|`utf8`|