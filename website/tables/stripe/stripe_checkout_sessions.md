# Table: stripe_checkout_sessions

This table shows data for Stripe Checkout Sessions.

https://stripe.com/docs/api/checkout/sessions

The primary key for this table is **id**.

## Relations

The following tables depend on stripe_checkout_sessions:
  - [stripe_checkout_session_line_items](stripe_checkout_session_line_items)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|after_expiration|`json`|
|allow_promotion_codes|`bool`|
|amount_subtotal|`int64`|
|amount_total|`int64`|
|automatic_tax|`json`|
|billing_address_collection|`utf8`|
|cancel_url|`utf8`|
|client_reference_id|`utf8`|
|consent|`json`|
|consent_collection|`json`|
|created|`timestamp[us, tz=UTC]`|
|currency|`utf8`|
|currency_conversion|`json`|
|customer|`json`|
|customer_creation|`utf8`|
|customer_details|`json`|
|customer_email|`utf8`|
|custom_fields|`json`|
|custom_text|`json`|
|expires_at|`int64`|
|invoice|`json`|
|invoice_creation|`json`|
|line_items|`json`|
|livemode|`bool`|
|locale|`utf8`|
|metadata|`json`|
|mode|`utf8`|
|object|`utf8`|
|payment_intent|`json`|
|payment_link|`json`|
|payment_method_collection|`utf8`|
|payment_method_options|`json`|
|payment_method_types|`list<item: utf8, nullable>`|
|payment_status|`utf8`|
|phone_number_collection|`json`|
|recovered_from|`utf8`|
|setup_intent|`json`|
|shipping_address_collection|`json`|
|shipping_cost|`json`|
|shipping_details|`json`|
|shipping_options|`json`|
|status|`utf8`|
|submit_type|`utf8`|
|subscription|`json`|
|success_url|`utf8`|
|tax_id_collection|`json`|
|total_details|`json`|
|url|`utf8`|