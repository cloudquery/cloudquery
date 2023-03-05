# Table: stripe_charges

https://stripe.com/docs/api/charges

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
|amount_captured|Int|
|amount_refunded|Int|
|application|JSON|
|application_fee|JSON|
|application_fee_amount|Int|
|authorization_code|String|
|balance_transaction|JSON|
|billing_details|JSON|
|calculated_statement_descriptor|String|
|captured|Bool|
|currency|String|
|customer|JSON|
|description|String|
|destination|JSON|
|dispute|JSON|
|disputed|Bool|
|failure_balance_transaction|JSON|
|failure_code|String|
|failure_message|String|
|fraud_details|JSON|
|invoice|JSON|
|level3|JSON|
|livemode|Bool|
|metadata|JSON|
|object|String|
|on_behalf_of|JSON|
|outcome|JSON|
|paid|Bool|
|payment_intent|JSON|
|payment_method|String|
|payment_method_details|JSON|
|radar_options|JSON|
|receipt_email|String|
|receipt_number|String|
|receipt_url|String|
|refunded|Bool|
|refunds|JSON|
|review|JSON|
|shipping|JSON|
|source|JSON|
|source_transfer|JSON|
|statement_descriptor|String|
|statement_descriptor_suffix|String|
|status|String|
|transfer|JSON|
|transfer_data|JSON|
|transfer_group|String|