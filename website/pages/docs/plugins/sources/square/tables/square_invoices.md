# Table: square_invoices

This table shows data for Square Invoices.

The primary key for this table is **id**.

## Relations

This table depends on [square_locations](square_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`utf8`|
|version|`int64`|
|location_id|`utf8`|
|order_id|`utf8`|
|primary_recipient|`json`|
|payment_requests|`json`|
|delivery_method|`utf8`|
|invoice_number|`utf8`|
|title|`utf8`|
|description|`utf8`|
|scheduled_at|`utf8`|
|public_url|`utf8`|
|next_payment_amount_money|`json`|
|status|`utf8`|
|timezone|`utf8`|
|created_at|`utf8`|
|updated_at|`utf8`|
|accepted_payment_methods|`json`|
|custom_fields|`json`|
|subscription_id|`utf8`|
|sale_or_service_date|`utf8`|
|payment_conditions|`utf8`|
|store_payment_method_enabled|`bool`|