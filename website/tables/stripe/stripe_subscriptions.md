# Table: stripe_subscriptions

This table shows data for Stripe Subscriptions.

https://stripe.com/docs/api/subscriptions

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|application|`json`|
|application_fee_percent|`float64`|
|automatic_tax|`json`|
|billing_cycle_anchor|`int64`|
|billing_thresholds|`json`|
|cancel_at|`int64`|
|cancel_at_period_end|`bool`|
|canceled_at|`int64`|
|cancellation_details|`json`|
|collection_method|`utf8`|
|currency|`utf8`|
|current_period_end|`int64`|
|current_period_start|`int64`|
|customer|`json`|
|days_until_due|`int64`|
|default_payment_method|`json`|
|default_source|`json`|
|default_tax_rates|`json`|
|description|`utf8`|
|discount|`json`|
|ended_at|`int64`|
|items|`json`|
|latest_invoice|`json`|
|livemode|`bool`|
|metadata|`json`|
|next_pending_invoice_item_invoice|`int64`|
|object|`utf8`|
|on_behalf_of|`json`|
|pause_collection|`json`|
|payment_settings|`json`|
|pending_invoice_item_interval|`json`|
|pending_setup_intent|`json`|
|pending_update|`json`|
|schedule|`json`|
|start_date|`int64`|
|status|`utf8`|
|test_clock|`json`|
|transfer_data|`json`|
|trial_end|`int64`|
|trial_settings|`json`|
|trial_start|`int64`|