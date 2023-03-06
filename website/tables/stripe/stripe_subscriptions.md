# Table: stripe_subscriptions

https://stripe.com/docs/api/subscriptions

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created (Incremental Key)|Timestamp|
|application|JSON|
|application_fee_percent|Float|
|automatic_tax|JSON|
|billing_cycle_anchor|Int|
|billing_thresholds|JSON|
|cancel_at|Int|
|cancel_at_period_end|Bool|
|canceled_at|Int|
|collection_method|String|
|currency|String|
|current_period_end|Int|
|current_period_start|Int|
|customer|JSON|
|days_until_due|Int|
|default_payment_method|JSON|
|default_source|JSON|
|default_tax_rates|JSON|
|description|String|
|discount|JSON|
|ended_at|Int|
|items|JSON|
|latest_invoice|JSON|
|livemode|Bool|
|metadata|JSON|
|next_pending_invoice_item_invoice|Int|
|object|String|
|on_behalf_of|JSON|
|pause_collection|JSON|
|payment_settings|JSON|
|pending_invoice_item_interval|JSON|
|pending_setup_intent|JSON|
|pending_update|JSON|
|schedule|JSON|
|start_date|Int|
|status|String|
|test_clock|JSON|
|transfer_data|JSON|
|trial_end|Int|
|trial_start|Int|