# Table: square_locations

This table shows data for Square Locations.

The primary key for this table is **id**.

## Relations

The following tables depend on square_locations:

  - [square_invoices](square_invoices)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`string`|
|name|`string`|
|address|`extension<json<JSONType>>`|
|timezone|`string`|
|capabilities|`extension<json<JSONType>>`|
|status|`string`|
|created_at|`string`|
|merchant_id|`string`|
|country|`string`|
|language_code|`string`|
|currency|`string`|
|phone_number|`string`|
|business_name|`string`|
|type|`string`|
|website_url|`string`|
|business_hours|`extension<json<JSONType>>`|
|business_email|`string`|
|description|`string`|
|twitter_username|`string`|
|instagram_username|`string`|
|facebook_url|`string`|
|coordinates|`extension<json<JSONType>>`|
|logo_url|`string`|
|pos_background_url|`string`|
|mcc|`string`|
|full_format_logo_url|`string`|
|tax_ids|`extension<json<JSONType>>`|