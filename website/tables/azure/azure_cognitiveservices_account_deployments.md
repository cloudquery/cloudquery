# Table: azure_cognitiveservices_account_deployments

This table shows data for Azure Cognitive Services Account Deployments.

https://learn.microsoft.com/en-us/rest/api/cognitiveservices/accountmanagement/deployments/list?tabs=HTTP#deployment

The primary key for this table is **id**.

## Relations

This table depends on [azure_cognitiveservices_accounts](azure_cognitiveservices_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|etag|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|