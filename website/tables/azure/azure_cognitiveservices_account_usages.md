# Table: azure_cognitiveservices_account_usages

This table shows data for Azure Cognitive Services Account Usages.

https://learn.microsoft.com/en-us/rest/api/cognitiveservices/accountmanagement/accounts/list-usages?tabs=HTTP#usage

The primary key for this table is **_cq_id**.

## Relations

This table depends on [azure_cognitiveservices_accounts](azure_cognitiveservices_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|account_id|`utf8`|
|current_value|`float64`|
|limit|`float64`|
|name|`json`|
|next_reset_time|`utf8`|
|quota_period|`utf8`|
|status|`utf8`|
|unit|`utf8`|