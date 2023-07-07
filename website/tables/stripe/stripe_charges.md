# Table: stripe_charges

This table shows data for Stripe Charges.

https://stripe.com/docs/api/charges

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
|amount_captured|`int64`|
|amount_refunded|`int64`|
|application|`json`|
|application_fee|`json`|
|application_fee_amount|`int64`|
|authorization_code|`utf8`|
|balance_transaction|`json`|
|billing_details|`json`|
|calculated_statement_descriptor|`utf8`|
|captured|`bool`|
|currency|`utf8`|
|customer|`json`|
|description|`utf8`|
|destination|`json`|
|dispute|`json`|
|disputed|`bool`|
|failure_balance_transaction|`json`|
|failure_code|`utf8`|
|failure_message|`utf8`|
|fraud_details|`json`|
|invoice|`json`|
|level3|`json`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|on_behalf_of|`json`|
|outcome|`json`|
|paid|`bool`|
|payment_intent|`json`|
|payment_method|`utf8`|
|payment_method_details|`json`|
|radar_options|`json`|
|receipt_email|`utf8`|
|receipt_number|`utf8`|
|receipt_url|`utf8`|
|refunded|`bool`|
|refunds|`json`|
|review|`json`|
|shipping|`json`|
|source|`json`|
|source_transfer|`json`|
|statement_descriptor|`utf8`|
|statement_descriptor_suffix|`utf8`|
|status|`utf8`|
|transfer|`json`|
|transfer_data|`json`|
|transfer_group|`utf8`|