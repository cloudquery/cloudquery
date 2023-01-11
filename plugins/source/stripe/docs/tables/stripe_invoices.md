# Table: stripe_invoices

https://stripe.com/docs/api/invoices

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
|account_country|String|
|account_name|String|
|account_tax_ids|JSON|
|amount_due|Int|
|amount_paid|Int|
|amount_remaining|Int|
|application|JSON|
|application_fee_amount|Int|
|attempt_count|Int|
|attempted|Bool|
|auto_advance|Bool|
|automatic_tax|JSON|
|billing_reason|String|
|charge|JSON|
|collection_method|String|
|currency|String|
|customer|JSON|
|customer_address|JSON|
|customer_email|String|
|customer_name|String|
|customer_phone|String|
|customer_shipping|JSON|
|customer_tax_exempt|String|
|customer_tax_ids|JSON|
|custom_fields|JSON|
|default_payment_method|JSON|
|default_source|JSON|
|default_tax_rates|JSON|
|deleted|Bool|
|description|String|
|discount|JSON|
|discounts|JSON|
|due_date|Int|
|ending_balance|Int|
|footer|String|
|from_invoice|JSON|
|hosted_invoice_url|String|
|invoice_pdf|String|
|last_finalization_error|JSON|
|latest_revision|JSON|
|lines|JSON|
|livemode|Bool|
|metadata|JSON|
|next_payment_attempt|Int|
|number|String|
|object|String|
|on_behalf_of|JSON|
|paid|Bool|
|paid_out_of_band|Bool|
|payment_intent|JSON|
|payment_settings|JSON|
|period_end|Int|
|period_start|Int|
|post_payment_credit_notes_amount|Int|
|pre_payment_credit_notes_amount|Int|
|quote|JSON|
|receipt_number|String|
|rendering_options|JSON|
|starting_balance|Int|
|statement_descriptor|String|
|status|String|
|status_transitions|JSON|
|subscription|JSON|
|subscription_proration_date|Int|
|subtotal|Int|
|subtotal_excluding_tax|Int|
|tax|Int|
|test_clock|JSON|
|threshold_reason|JSON|
|total|Int|
|total_discount_amounts|JSON|
|total_excluding_tax|Int|
|total_tax_amounts|JSON|
|transfer_data|JSON|
|webhooks_delivered_at|Int|