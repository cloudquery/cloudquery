# Table: stripe_payment_intents

This table shows data for Stripe Payment Intents.

https://stripe.com/docs/api/payment_intents

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|amount|`int64`|
|amount_capturable|`int64`|
|amount_details|`json`|
|amount_received|`int64`|
|application|`json`|
|application_fee_amount|`int64`|
|automatic_payment_methods|`json`|
|canceled_at|`int64`|
|cancellation_reason|`utf8`|
|capture_method|`utf8`|
|client_secret|`utf8`|
|confirmation_method|`utf8`|
|currency|`utf8`|
|customer|`json`|
|description|`utf8`|
|invoice|`json`|
|last_payment_error|`json`|
|latest_charge|`json`|
|livemode|`bool`|
|metadata|`json`|
|next_action|`json`|
|object|`utf8`|
|on_behalf_of|`json`|
|payment_method|`json`|
|payment_method_options|`json`|
|payment_method_types|`list<item: utf8, nullable>`|
|processing|`json`|
|receipt_email|`utf8`|
|review|`json`|
|setup_future_usage|`utf8`|
|shipping|`json`|
|source|`json`|
|statement_descriptor|`utf8`|
|statement_descriptor_suffix|`utf8`|
|status|`utf8`|
|transfer_data|`json`|
|transfer_group|`utf8`|