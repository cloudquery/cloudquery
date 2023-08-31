# Table: stripe_quotes

This table shows data for Stripe Quotes.

https://stripe.com/docs/api/quotes

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|amount_subtotal|`int64`|
|amount_total|`int64`|
|application|`json`|
|application_fee_amount|`int64`|
|application_fee_percent|`float64`|
|automatic_tax|`json`|
|collection_method|`utf8`|
|computed|`json`|
|created|`timestamp[us, tz=UTC]`|
|currency|`utf8`|
|customer|`json`|
|default_tax_rates|`json`|
|description|`utf8`|
|discounts|`json`|
|expires_at|`int64`|
|footer|`utf8`|
|from_quote|`json`|
|header|`utf8`|
|invoice|`json`|
|invoice_settings|`json`|
|line_items|`json`|
|livemode|`bool`|
|metadata|`json`|
|number|`utf8`|
|object|`utf8`|
|on_behalf_of|`json`|
|status|`utf8`|
|status_transitions|`json`|
|subscription|`json`|
|subscription_data|`json`|
|subscription_schedule|`json`|
|test_clock|`json`|
|total_details|`json`|
|transfer_data|`json`|