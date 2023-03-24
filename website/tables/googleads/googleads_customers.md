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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|resource_name (PK)|String|
|id (PK)|Int|
|descriptive_name|String|
|currency_code|String|
|time_zone|String|
|tracking_url_template|String|
|final_url_suffix|String|
|auto_tagging_enabled|Bool|
|has_partners_badge|Bool|
|manager|Bool|
|test_account|Bool|
|call_reporting_setting|JSON|
|conversion_tracking_setting|JSON|
|remarketing_setting|JSON|
|pay_per_conversion_eligibility_failure_reasons|IntArray|
|optimization_score|Float|
|optimization_score_weight|Float|
|status|String|
|location_asset_auto_migration_done|Bool|
|image_asset_auto_migration_done|Bool|
|location_asset_auto_migration_done_date_time|String|
|image_asset_auto_migration_done_date_time|String|