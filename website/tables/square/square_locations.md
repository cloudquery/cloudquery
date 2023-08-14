# Table: square_locations

This table shows data for Square Locations.

The primary key for this table is **id**.

## Relations

The following tables depend on square_locations:
  - [square_invoices](square_invoices)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|id (PK)|`utf8`|
|name|`utf8`|
|address|`json`|
|timezone|`utf8`|
|capabilities|`json`|
|status|`utf8`|
|created_at|`utf8`|
|merchant_id|`utf8`|
|country|`utf8`|
|language_code|`utf8`|
|currency|`utf8`|
|phone_number|`utf8`|
|business_name|`utf8`|
|type|`utf8`|
|website_url|`utf8`|
|business_hours|`json`|
|business_email|`utf8`|
|description|`utf8`|
|twitter_username|`utf8`|
|instagram_username|`utf8`|
|facebook_url|`utf8`|
|coordinates|`json`|
|logo_url|`utf8`|
|pos_background_url|`utf8`|
|mcc|`utf8`|
|full_format_logo_url|`utf8`|
|tax_ids|`json`|