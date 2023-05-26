# Table: googleads_campaigns

This table shows data for Google Ads Campaigns.

https://developers.google.com/google-ads/api/reference/rpc/v13/Campaign

The composite primary key for this table is (**customer_id**, **resource_name**, **id**).

## Relations

The following tables depend on googleads_campaigns:
  - [googleads_campaign_criteria](googleads_campaign_criteria)
  - [googleads_campaign_labels](googleads_campaign_labels)

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
|primary_status|utf8|
|primary_status_reasons|list<item: int64, nullable>|
|status|utf8|
|serving_status|utf8|
|bidding_strategy_system_status|utf8|
|ad_serving_optimization_status|utf8|
|advertising_channel_type|utf8|
|advertising_channel_sub_type|utf8|
|tracking_url_template|utf8|
|url_custom_parameters|json|
|local_services_campaign_settings|json|
|travel_campaign_settings|json|
|real_time_bidding_setting|json|
|network_settings|json|
|hotel_setting|json|
|dynamic_search_ads_setting|json|
|shopping_setting|json|
|audience_setting|json|
|geo_target_type_setting|json|
|local_campaign_setting|json|
|app_campaign_setting|json|
|labels|list<item: utf8, nullable>|
|experiment_type|utf8|
|base_campaign|utf8|
|campaign_budget|utf8|
|bidding_strategy_type|utf8|
|accessible_bidding_strategy|utf8|
|start_date|utf8|
|campaign_group|utf8|
|end_date|utf8|
|final_url_suffix|utf8|
|frequency_caps|json|
|video_brand_safety_suitability|utf8|
|vanity_pharma|json|
|selective_optimization|json|
|optimization_goal_setting|json|
|tracking_setting|json|
|payment_mode|utf8|
|optimization_score|float64|
|excluded_parent_asset_field_types|list<item: int64, nullable>|
|excluded_parent_asset_set_types|list<item: int64, nullable>|
|url_expansion_opt_out|bool|
|performance_max_upgrade|json|
|hotel_property_asset_set|utf8|