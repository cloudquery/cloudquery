# Table: gcp_billing_billing_accounts

This table shows data for GCP Billing Billing Accounts.

https://cloud.google.com/billing/docs/reference/rest/v1/billingAccounts#BillingAccount

The primary key for this table is **name**.

## Relations

The following tables depend on gcp_billing_billing_accounts:
  - [gcp_billing_budgets](gcp_billing_budgets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|name (PK)|`utf8`|
|open|`bool`|
|display_name|`utf8`|
|master_billing_account|`utf8`|