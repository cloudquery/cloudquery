# Table: stripe_invoices

This table shows data for Stripe Invoices.

https://stripe.com/docs/api/invoices

The primary key for this table is **id**.
It supports incremental syncs based on the **created** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created (Incremental Key)|`timestamp[us, tz=UTC]`|
|account_country|`utf8`|
|account_name|`utf8`|
|account_tax_ids|`json`|
|amount_due|`int64`|
|amount_paid|`int64`|
|amount_remaining|`int64`|
|amount_shipping|`int64`|
|application|`json`|
|application_fee_amount|`int64`|
|attempt_count|`int64`|
|attempted|`bool`|
|auto_advance|`bool`|
|automatic_tax|`json`|
|billing_reason|`utf8`|
|charge|`json`|
|collection_method|`utf8`|
|currency|`utf8`|
|customer|`json`|
|customer_address|`json`|
|customer_email|`utf8`|
|customer_name|`utf8`|
|customer_phone|`utf8`|
|customer_shipping|`json`|
|customer_tax_exempt|`utf8`|
|customer_tax_ids|`json`|
|custom_fields|`json`|
|default_payment_method|`json`|
|default_source|`json`|
|default_tax_rates|`json`|
|deleted|`bool`|
|description|`utf8`|
|discount|`json`|
|discounts|`json`|
|due_date|`int64`|
|ending_balance|`int64`|
|footer|`utf8`|
|from_invoice|`json`|
|hosted_invoice_url|`utf8`|
|invoice_pdf|`utf8`|
|last_finalization_error|`json`|
|latest_revision|`json`|
|lines|`json`|
|livemode|`bool`|
|metadata|`json`|
|next_payment_attempt|`int64`|
|number|`utf8`|
|object|`utf8`|
|on_behalf_of|`json`|
|paid|`bool`|
|paid_out_of_band|`bool`|
|payment_intent|`json`|
|payment_settings|`json`|
|period_end|`int64`|
|period_start|`int64`|
|post_payment_credit_notes_amount|`int64`|
|pre_payment_credit_notes_amount|`int64`|
|quote|`json`|
|receipt_number|`utf8`|
|rendering_options|`json`|
|shipping_cost|`json`|
|shipping_details|`json`|
|starting_balance|`int64`|
|statement_descriptor|`utf8`|
|status|`utf8`|
|status_transitions|`json`|
|subscription|`json`|
|subscription_proration_date|`int64`|
|subtotal|`int64`|
|subtotal_excluding_tax|`int64`|
|tax|`int64`|
|test_clock|`json`|
|threshold_reason|`json`|
|total|`int64`|
|total_discount_amounts|`json`|
|total_excluding_tax|`int64`|
|total_tax_amounts|`json`|
|transfer_data|`json`|
|webhooks_delivered_at|`int64`|