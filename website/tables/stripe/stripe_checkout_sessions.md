# Table: stripe_checkout_sessions

https://stripe.com/docs/api/checkout_sessions

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|after_expiration|JSON|
|allow_promotion_codes|Bool|
|amount_subtotal|Int|
|amount_total|Int|
|automatic_tax|JSON|
|billing_address_collection|String|
|cancel_url|String|
|client_reference_id|String|
|consent|JSON|
|consent_collection|JSON|
|created|Timestamp|
|currency|String|
|customer|JSON|
|customer_creation|String|
|customer_details|JSON|
|customer_email|String|
|custom_text|JSON|
|expires_at|Int|
|invoice|JSON|
|invoice_creation|JSON|
|line_items|JSON|
|livemode|Bool|
|locale|String|
|metadata|JSON|
|mode|String|
|object|String|
|payment_intent|JSON|
|payment_link|JSON|
|payment_method_collection|String|
|payment_method_options|JSON|
|payment_method_types|StringArray|
|payment_status|String|
|phone_number_collection|JSON|
|recovered_from|String|
|setup_intent|JSON|
|shipping_address_collection|JSON|
|shipping_cost|JSON|
|shipping_details|JSON|
|shipping_options|JSON|
|status|String|
|submit_type|String|
|subscription|JSON|
|success_url|String|
|tax_id_collection|JSON|
|total_details|JSON|
|url|String|