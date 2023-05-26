# Table: googleads_ad_group_criteria

This table shows data for Google Ads Ad Group Criteria.

https://developers.google.com/google-ads/api/reference/rpc/v13/AdGroupCriterion

The composite primary key for this table is (**customer_id**, **id**, **resource_name**, **ad_group**).

## Relations

This table depends on [googleads_ad_groups](googleads_ad_groups).

The following tables depend on googleads_ad_group_criteria:
  - [googleads_ad_group_criterion_labels](googleads_ad_group_criterion_labels)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|customer_id (PK)|int64|
|id (PK)|int64|
|resource_name (PK)|utf8|
|criterion_id|int64|
|display_name|utf8|
|status|utf8|
|quality_info|json|
|ad_group (PK)|utf8|
|type|utf8|
|negative|bool|
|system_serving_status|utf8|
|approval_status|utf8|
|disapproval_reasons|list<item: utf8, nullable>|
|labels|list<item: utf8, nullable>|
|bid_modifier|float64|
|cpc_bid_micros|int64|
|cpm_bid_micros|int64|
|cpv_bid_micros|int64|
|percent_cpc_bid_micros|int64|
|effective_cpc_bid_micros|int64|
|effective_cpm_bid_micros|int64|
|effective_cpv_bid_micros|int64|
|effective_percent_cpc_bid_micros|int64|
|effective_cpc_bid_source|utf8|
|effective_cpm_bid_source|utf8|
|effective_cpv_bid_source|utf8|
|effective_percent_cpc_bid_source|utf8|
|position_estimates|json|
|final_urls|list<item: utf8, nullable>|
|final_mobile_urls|list<item: utf8, nullable>|
|final_url_suffix|utf8|
|tracking_url_template|utf8|
|url_custom_parameters|json|