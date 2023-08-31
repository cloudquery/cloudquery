# Table: azure_datalakeanalytics_accounts

This table shows data for Azure Data Lake Analytics Accounts.

https://learn.microsoft.com/en-us/rest/api/datalakeanalytics/accounts/list?tabs=HTTP#datalakeanalyticsaccountbasic

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|name|`utf8`|
|properties|`json`|
|tags|`json`|
|type|`utf8`|