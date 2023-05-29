# Table: googleads_customers

This table shows data for Google Ads Customers.

https://developers.google.com/google-ads/api/reference/rpc/v13/Customer

The composite primary key for this table is (**resource_name**, **id**).

## Relations

The following tables depend on googleads_customers:
  - [googleads_customer_labels](googleads_customer_labels)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|resource_name (PK)|`utf8`|
|id (PK)|`int64`|
|descriptive_name|`utf8`|
|currency_code|`utf8`|
|time_zone|`utf8`|
|tracking_url_template|`utf8`|
|final_url_suffix|`utf8`|
|auto_tagging_enabled|`bool`|
|has_partners_badge|`bool`|
|manager|`bool`|
|test_account|`bool`|
|call_reporting_setting|`json`|
|conversion_tracking_setting|`json`|
|remarketing_setting|`json`|
|pay_per_conversion_eligibility_failure_reasons|`list<item: int64, nullable>`|
|optimization_score|`float64`|
|optimization_score_weight|`float64`|
|status|`utf8`|
|location_asset_auto_migration_done|`bool`|
|image_asset_auto_migration_done|`bool`|
|location_asset_auto_migration_done_date_time|`utf8`|
|image_asset_auto_migration_done_date_time|`utf8`|