# Table: facebookmarketing_adsets

This table shows data for Facebook Marketing Ad Sets.

https://developers.facebook.com/docs/marketing-api/reference/ad-campaign#Reading

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|adlabels|`json`|
|adset_schedule|`json`|
|asset_feed_id|`utf8`|
|attribution_spec|`json`|
|bid_adjustments|`json`|
|bid_amount|`int64`|
|bid_constraints|`json`|
|bid_info|`json`|
|bid_strategy|`utf8`|
|billing_event|`utf8`|
|budget_remaining|`utf8`|
|campaign|`json`|
|campaign_attribution|`utf8`|
|campaign_id|`utf8`|
|configured_status|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|creative_sequence|`list<item: utf8, nullable>`|
|daily_budget|`utf8`|
|daily_min_spend_target|`utf8`|
|daily_spend_cap|`utf8`|
|destination_type|`utf8`|
|effective_status|`utf8`|
|end_time|`timestamp[us, tz=UTC]`|
|frequency_control_specs|`json`|
|id (PK)|`utf8`|
|is_dynamic_creative|`bool`|
|issues_info|`json`|
|learning_stage_info|`json`|
|lifetime_budget|`utf8`|
|lifetime_imps|`int64`|
|lifetime_min_spend_target|`utf8`|
|lifetime_spend_cap|`utf8`|
|multi_optimization_goal_weight|`utf8`|
|name|`utf8`|
|optimization_sub_event|`utf8`|
|pacing_type|`list<item: utf8, nullable>`|
|promoted_object|`json`|
|recommendations|`json`|
|recurring_budget_semantics|`bool`|
|review_feedback|`utf8`|
|rf_prediction_id|`utf8`|
|source_adset_id|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|targeting|`json`|
|time_based_ad_rotation_id_blocks|`json`|
|time_based_ad_rotation_intervals|`list<item: int64, nullable>`|
|updated_time|`timestamp[us, tz=UTC]`|
|use_new_app_click|`bool`|