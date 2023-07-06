# Table: facebookmarketing_campaigns

This table shows data for Facebook Marketing Campaigns.

https://developers.facebook.com/docs/marketing-api/reference/ad-campaign-group/

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|ad_strategy_group_id|`utf8`|
|ad_strategy_id|`utf8`|
|adlabels|`json`|
|bid_strategy|`utf8`|
|boosted_object_id|`utf8`|
|brand_lift_studies|`json`|
|budget_rebalance_flag|`bool`|
|budget_remaining|`utf8`|
|buying_type|`utf8`|
|can_create_brand_lift_study|`bool`|
|can_use_spend_cap|`bool`|
|configured_status|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|daily_budget|`utf8`|
|effective_status|`utf8`|
|has_secondary_skadnetwork_reporting|`bool`|
|id (PK)|`utf8`|
|is_skadnetwork_attribution|`bool`|
|issues_info|`json`|
|last_budget_toggling_time|`timestamp[us, tz=UTC]`|
|lifetime_budget|`utf8`|
|name|`utf8`|
|objective|`utf8`|
|pacing_type|`list<item: utf8, nullable>`|
|primary_attribution|`utf8`|
|promoted_object|`json`|
|recommendations|`json`|
|smart_promotion_type|`utf8`|
|source_campaign_id|`utf8`|
|special_ad_categories|`list<item: utf8, nullable>`|
|special_ad_category|`utf8`|
|special_ad_category_country|`list<item: utf8, nullable>`|
|spend_cap|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|stop_time|`timestamp[us, tz=UTC]`|
|topline_id|`utf8`|
|updated_time|`timestamp[us, tz=UTC]`|