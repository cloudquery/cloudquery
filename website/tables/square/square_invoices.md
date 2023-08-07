# Table: square_invoices

This table shows data for Square Invoices.

The primary key for this table is **id**.

## Relations

This table depends on [square_locations](square_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`string`|
|version|`int64`|
|location_id|`string`|
|order_id|`string`|
|primary_recipient|`extension<json<JSONType>>`|
|payment_requests|`extension<json<JSONType>>`|
|delivery_method|`string`|
|invoice_number|`string`|
|title|`string`|
|description|`string`|
|scheduled_at|`string`|
|public_url|`string`|
|next_payment_amount_money|`extension<json<JSONType>>`|
|status|`string`|
|timezone|`string`|
|created_at|`string`|
|updated_at|`string`|
|accepted_payment_methods|`extension<json<JSONType>>`|
|custom_fields|`extension<json<JSONType>>`|
|subscription_id|`string`|
|sale_or_service_date|`string`|
|payment_conditions|`string`|
|store_payment_method_enabled|`bool`|