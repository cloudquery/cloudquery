# Table: azure_billing_accounts

This table shows data for Azure Billing Accounts.

https://learn.microsoft.com/en-us/rest/api/billing/2020-05-01/billing-accounts/list?tabs=HTTP#billingaccount

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|