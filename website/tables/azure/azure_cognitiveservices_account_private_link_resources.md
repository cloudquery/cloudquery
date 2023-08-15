# Table: azure_cognitiveservices_account_private_link_resources

This table shows data for Azure Cognitive Services Account Private Link Resources.

https://learn.microsoft.com/en-us/rest/api/cognitiveservices/accountmanagement/private-link-resources/list?tabs=HTTP#privatelinkresource

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
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|