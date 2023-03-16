# Table: azure_billing_accounts

This table shows data for Azure Billing Accounts.

https://learn.microsoft.com/en-us/rest/api/billing/2020-05-01/billing-accounts/list?tabs=HTTP#billingaccount

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|properties|JSON|
|id (PK)|String|
|name|String|
|type|String|