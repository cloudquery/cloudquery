# Table: gcp_billing_budgets

https://cloud.google.com/billing/docs/reference/budget/rest/v1/billingAccounts.budgets#Budget

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_billing_billing_accounts](gcp_billing_billing_accounts.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|display_name|String|
|budget_filter|JSON|
|amount|JSON|
|threshold_rules|JSON|
|notifications_rule|JSON|
|etag|String|