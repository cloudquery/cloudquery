# Table: azure_cognitiveservices_resource_skus

This table shows data for Azure Cognitive Services Resource Skus.

https://learn.microsoft.com/en-us/rest/api/cognitiveservices/accountmanagement/resource-skus/list?tabs=HTTP#resourcesku

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|kind|`utf8`|
|locations|`list<item: utf8, nullable>`|
|name|`utf8`|
|resource_type|`utf8`|
|restrictions|`json`|
|tier|`utf8`|