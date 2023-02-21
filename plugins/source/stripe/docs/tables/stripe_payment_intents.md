# Table: stripe_payment_intents

https://stripe.com/docs/api/payment_intents

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
|amount|Int|
|amount_capturable|Int|
|amount_details|JSON|
|amount_received|Int|
|application|JSON|
|application_fee_amount|Int|
|automatic_payment_methods|JSON|
|canceled_at|Int|
|cancellation_reason|String|
|capture_method|String|
|client_secret|String|
|confirmation_method|String|
|currency|String|
|customer|JSON|
|description|String|
|invoice|JSON|
|last_payment_error|JSON|
|latest_charge|JSON|
|livemode|Bool|
|metadata|JSON|
|next_action|JSON|
|object|String|
|on_behalf_of|JSON|
|payment_method|JSON|
|payment_method_options|JSON|
|payment_method_types|StringArray|
|processing|JSON|
|receipt_email|String|
|review|JSON|
|setup_future_usage|String|
|shipping|JSON|
|source|JSON|
|statement_descriptor|String|
|statement_descriptor_suffix|String|
|status|String|
|transfer_data|JSON|
|transfer_group|String|