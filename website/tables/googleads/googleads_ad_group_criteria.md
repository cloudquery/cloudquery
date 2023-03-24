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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|customer_id (PK)|Int|
|id (PK)|Int|
|resource_name (PK)|String|
|criterion_id|Int|
|display_name|String|
|status|String|
|quality_info|JSON|
|ad_group (PK)|String|
|type|String|
|negative|Bool|
|system_serving_status|String|
|approval_status|String|
|disapproval_reasons|StringArray|
|labels|StringArray|
|bid_modifier|Float|
|cpc_bid_micros|Int|
|cpm_bid_micros|Int|
|cpv_bid_micros|Int|
|percent_cpc_bid_micros|Int|
|effective_cpc_bid_micros|Int|
|effective_cpm_bid_micros|Int|
|effective_cpv_bid_micros|Int|
|effective_percent_cpc_bid_micros|Int|
|effective_cpc_bid_source|String|
|effective_cpm_bid_source|String|
|effective_cpv_bid_source|String|
|effective_percent_cpc_bid_source|String|
|position_estimates|JSON|
|final_urls|StringArray|
|final_mobile_urls|StringArray|
|final_url_suffix|String|
|tracking_url_template|String|
|url_custom_parameters|JSON|