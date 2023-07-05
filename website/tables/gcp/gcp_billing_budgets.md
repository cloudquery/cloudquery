# Table: gcp_billing_budgets

This table shows data for GCP Billing Budgets.

https://cloud.google.com/billing/docs/reference/budget/rest/v1/billingAccounts.budgets#Budget

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_billing_billing_accounts](gcp_billing_billing_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|budget_filter|`json`|
|amount|`json`|
|threshold_rules|`json`|
|notifications_rule|`json`|
|etag|`utf8`|