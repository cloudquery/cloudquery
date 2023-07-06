# Table: facebookmarketing_adaccounts

This table shows data for Facebook Marketing Ad Accounts.

https://developers.facebook.com/docs/marketing-api/reference/ad-account#Reading

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|account_status|`int64`|
|age|`float64`|
|agency_client_declaration|`json`|
|amount_spent|`utf8`|
|attribution_spec|`json`|
|balance|`utf8`|
|business|`json`|
|business_city|`utf8`|
|business_country_code|`utf8`|
|business_name|`utf8`|
|business_state|`utf8`|
|business_street|`utf8`|
|business_street2|`utf8`|
|business_zip|`utf8`|
|can_create_brand_lift_study|`bool`|
|capabilities|`list<item: utf8, nullable>`|
|created_time|`timestamp[us, tz=UTC]`|
|currency|`utf8`|
|disable_reason|`int64`|
|end_advertiser|`utf8`|
|end_advertiser_name|`utf8`|
|extended_credit_invoice_group|`json`|
|failed_delivery_checks|`json`|
|fb_entity|`int64`|
|funding_source|`utf8`|
|funding_source_details|`json`|
|has_advertiser_opted_in_odax|`bool`|
|has_migrated_permissions|`bool`|
|id (PK)|`utf8`|
|io_number|`utf8`|
|is_attribution_spec_system_default|`bool`|
|is_direct_deals_enabled|`bool`|
|is_in_3ds_authorization_enabled_market|`bool`|
|is_notifications_enabled|`bool`|
|is_personal|`int64`|
|is_prepay_account|`bool`|
|is_tax_id_required|`bool`|
|line_numbers|`list<item: int64, nullable>`|
|media_agency|`utf8`|
|min_campaign_group_spend_cap|`utf8`|
|min_daily_budget|`int64`|
|name|`utf8`|
|offsite_pixels_tos_accepted|`bool`|
|owner|`utf8`|
|partner|`utf8`|
|rf_spec|`json`|
|spend_cap|`utf8`|
|tax_id|`utf8`|
|tax_id_status|`int64`|
|tax_id_type|`utf8`|
|timezone_id|`int64`|
|timezone_name|`utf8`|
|timezone_offset_hours_utc|`float64`|
|tos_accepted|`json`|
|user_tasks|`list<item: utf8, nullable>`|
|user_tos_accepted|`json`|