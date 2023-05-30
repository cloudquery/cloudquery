# Table: azure_consumption_billing_account_marketplaces

This table shows data for Azure Consumption Billing Account Marketplaces.

https://learn.microsoft.com/en-us/rest/api/consumption/marketplaces/list?tabs=HTTP#marketplace

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|etag|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|tags|`json`|
|type|`utf8`|