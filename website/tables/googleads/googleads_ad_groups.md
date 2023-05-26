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
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|customer_id (PK)|int64|
|resource_name (PK)|utf8|
|id (PK)|int64|
|name|utf8|
|status|utf8|
|type|utf8|
|ad_rotation_mode|utf8|
|base_ad_group|utf8|
|tracking_url_template|utf8|
|url_custom_parameters|json|
|campaign|utf8|
|cpc_bid_micros|int64|
|effective_cpc_bid_micros|int64|
|cpm_bid_micros|int64|
|target_cpa_micros|int64|
|cpv_bid_micros|int64|
|target_cpm_micros|int64|
|target_roas|float64|
|percent_cpc_bid_micros|int64|
|optimized_targeting_enabled|bool|
|display_custom_bid_dimension|utf8|
|final_url_suffix|utf8|
|audience_setting|json|
|effective_target_cpa_micros|int64|
|effective_target_cpa_source|utf8|
|effective_target_roas|float64|
|effective_target_roas_source|utf8|
|labels|list<item: utf8, nullable>|
|excluded_parent_asset_field_types|list<item: int64, nullable>|
|excluded_parent_asset_set_types|list<item: int64, nullable>|