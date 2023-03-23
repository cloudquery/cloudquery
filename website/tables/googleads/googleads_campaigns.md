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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|customer_id (PK)|Int|
|resource_name (PK)|String|
|id (PK)|Int|
|name|String|
|primary_status|String|
|primary_status_reasons|IntArray|
|status|String|
|serving_status|String|
|bidding_strategy_system_status|String|
|ad_serving_optimization_status|String|
|advertising_channel_type|String|
|advertising_channel_sub_type|String|
|tracking_url_template|String|
|url_custom_parameters|JSON|
|local_services_campaign_settings|JSON|
|travel_campaign_settings|JSON|
|real_time_bidding_setting|JSON|
|network_settings|JSON|
|hotel_setting|JSON|
|dynamic_search_ads_setting|JSON|
|shopping_setting|JSON|
|audience_setting|JSON|
|geo_target_type_setting|JSON|
|local_campaign_setting|JSON|
|app_campaign_setting|JSON|
|labels|StringArray|
|experiment_type|String|
|base_campaign|String|
|campaign_budget|String|
|bidding_strategy_type|String|
|accessible_bidding_strategy|String|
|start_date|String|
|campaign_group|String|
|end_date|String|
|final_url_suffix|String|
|frequency_caps|JSON|
|video_brand_safety_suitability|String|
|vanity_pharma|JSON|
|selective_optimization|JSON|
|optimization_goal_setting|JSON|
|tracking_setting|JSON|
|payment_mode|String|
|optimization_score|Float|
|excluded_parent_asset_field_types|IntArray|
|excluded_parent_asset_set_types|IntArray|
|url_expansion_opt_out|Bool|
|performance_max_upgrade|JSON|
|hotel_property_asset_set|String|