# Table: gcp_billing_billing_accounts

https://cloud.google.com/billing/docs/reference/rest/v1/billingAccounts#BillingAccount

The primary key for this table is **name**.

## Relations

The following tables depend on gcp_billing_billing_accounts:
  - [gcp_billing_budgets](gcp_billing_budgets.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|name (PK)|String|
|open|Bool|
|display_name|String|
|master_billing_account|String|