# Table: azure_datalakestore_accounts

This table shows data for Azure Data Lake Store Accounts.

https://learn.microsoft.com/en-us/rest/api/datalakestore/accounts/list?tabs=HTTP#datalakestoreaccountbasic

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|location|String|
|name|String|
|properties|JSON|
|tags|JSON|
|type|String|