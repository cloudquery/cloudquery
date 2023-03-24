# Table: googleads_ad_groups

This table shows data for Google Ads Ad Groups.

https://developers.google.com/google-ads/api/reference/rpc/v13/AdGroup

The composite primary key for this table is (**customer_id**, **resource_name**, **id**).

## Relations

The following tables depend on googleads_ad_groups:
  - [googleads_ad_group_ads](googleads_ad_group_ads)
  - [googleads_ad_group_criteria](googleads_ad_group_criteria)
  - [googleads_ad_group_labels](googleads_ad_group_labels)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|customer_id (PK)|Int|
|resource_name (PK)|String|
|id (PK)|Int|
|name|String|
|status|String|
|type|String|
|ad_rotation_mode|String|
|base_ad_group|String|
|tracking_url_template|String|
|url_custom_parameters|JSON|
|campaign|String|
|cpc_bid_micros|Int|
|effective_cpc_bid_micros|Int|
|cpm_bid_micros|Int|
|target_cpa_micros|Int|
|cpv_bid_micros|Int|
|target_cpm_micros|Int|
|target_roas|Float|
|percent_cpc_bid_micros|Int|
|optimized_targeting_enabled|Bool|
|display_custom_bid_dimension|String|
|final_url_suffix|String|
|audience_setting|JSON|
|effective_target_cpa_micros|Int|
|effective_target_cpa_source|String|
|effective_target_roas|Float|
|effective_target_roas_source|String|
|labels|StringArray|
|excluded_parent_asset_field_types|IntArray|
|excluded_parent_asset_set_types|IntArray|