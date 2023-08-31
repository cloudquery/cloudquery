# Table: stripe_payment_links

This table shows data for Stripe Payment Links.

https://stripe.com/docs/api/payment_links

The primary key for this table is **id**.

## Relations

The following tables depend on stripe_payment_links:
  - [stripe_payment_link_line_items](stripe_payment_link_line_items)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|active|`bool`|
|after_completion|`json`|
|allow_promotion_codes|`bool`|
|application_fee_amount|`int64`|
|application_fee_percent|`float64`|
|automatic_tax|`json`|
|billing_address_collection|`utf8`|
|consent_collection|`json`|
|currency|`utf8`|
|customer_creation|`utf8`|
|custom_fields|`json`|
|custom_text|`json`|
|invoice_creation|`json`|
|line_items|`json`|
|livemode|`bool`|
|metadata|`json`|
|object|`utf8`|
|on_behalf_of|`json`|
|payment_intent_data|`json`|
|payment_method_collection|`utf8`|
|payment_method_types|`list<item: utf8, nullable>`|
|phone_number_collection|`json`|
|shipping_address_collection|`json`|
|shipping_options|`json`|
|submit_type|`utf8`|
|subscription_data|`json`|
|tax_id_collection|`json`|
|transfer_data|`json`|
|url|`utf8`|