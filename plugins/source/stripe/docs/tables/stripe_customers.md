# Table: stripe_customers

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|address|JSON|
|balance|Int|
|cash_balance|JSON|
|created|Int|
|currency|String|
|default_source|JSON|
|deleted|Bool|
|delinquent|Bool|
|description|String|
|discount|JSON|
|email|String|
|invoice_credit_balance|JSON|
|invoice_prefix|String|
|invoice_settings|JSON|
|livemode|Bool|
|metadata|JSON|
|name|String|
|next_invoice_sequence|Int|
|object|String|
|phone|String|
|preferred_locales|StringArray|
|shipping|JSON|
|sources|JSON|
|subscriptions|JSON|
|tax|JSON|
|tax_exempt|String|
|tax_ids|JSON|
|test_clock|JSON|