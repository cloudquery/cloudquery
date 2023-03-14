# Table: facebookmarketing_adsets

This table shows data for Facebookmarketing Adsets.

https://developers.facebook.com/docs/marketing-api/reference/ad-campaign#Reading

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|adlabels|JSON|
|adset_schedule|JSON|
|asset_feed_id|String|
|attribution_spec|JSON|
|bid_adjustments|JSON|
|bid_amount|Int|
|bid_constraints|JSON|
|bid_info|JSON|
|bid_strategy|String|
|billing_event|String|
|budget_remaining|String|
|campaign|JSON|
|campaign_attribution|String|
|campaign_id|String|
|configured_status|String|
|created_time|Timestamp|
|creative_sequence|StringArray|
|daily_budget|String|
|daily_min_spend_target|String|
|daily_spend_cap|String|
|destination_type|String|
|effective_status|String|
|end_time|Timestamp|
|frequency_control_specs|JSON|
|id (PK)|String|
|is_dynamic_creative|Bool|
|issues_info|JSON|
|learning_stage_info|JSON|
|lifetime_budget|String|
|lifetime_imps|Int|
|lifetime_min_spend_target|String|
|lifetime_spend_cap|String|
|multi_optimization_goal_weight|String|
|name|String|
|optimization_sub_event|String|
|pacing_type|StringArray|
|promoted_object|JSON|
|recommendations|JSON|
|recurring_budget_semantics|Bool|
|review_feedback|String|
|rf_prediction_id|String|
|source_adset_id|String|
|start_time|Timestamp|
|status|String|
|targeting|JSON|
|time_based_ad_rotation_id_blocks|JSON|
|time_based_ad_rotation_intervals|IntArray|
|updated_time|Timestamp|
|use_new_app_click|Bool|