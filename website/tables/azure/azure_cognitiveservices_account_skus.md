# Table: azure_cognitiveservices_account_skus

This table shows data for Azure Cognitive Services Account Skus.

https://learn.microsoft.com/en-us/rest/api/cognitiveservices/accountmanagement/accounts/list-skus?tabs=HTTP#accountsku

The primary key for this table is **_cq_id**.

## Relations

This table depends on [azure_cognitiveservices_accounts](azure_cognitiveservices_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|account_id|`utf8`|
|resource_type|`utf8`|
|sku|`json`|