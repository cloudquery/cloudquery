# Table: facebookmarketing_ads

This table shows data for Facebook Marketing Ads.

https://developers.facebook.com/docs/marketing-api/reference/adgroup

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|ad_review_feedback|`json`|
|adlabels|`json`|
|adset|`json`|
|adset_id|`utf8`|
|bid_amount|`int64`|
|bid_info|`json`|
|bid_type|`utf8`|
|campaign|`json`|
|campaign_id|`utf8`|
|configured_status|`utf8`|
|conversion_domain|`utf8`|
|conversion_specs|`json`|
|created_time|`timestamp[us, tz=UTC]`|
|creative|`json`|
|demolink_hash|`utf8`|
|display_sequence|`int64`|
|effective_status|`utf8`|
|engagement_audience|`bool`|
|failed_delivery_checks|`json`|
|id (PK)|`utf8`|
|issues_info|`json`|
|last_updated_by_app_id|`utf8`|
|name|`utf8`|
|preview_shareable_link|`utf8`|
|priority|`int64`|
|recommendations|`json`|
|source_ad_id|`utf8`|
|status|`utf8`|
|targeting|`json`|
|tracking_and_conversion_with_defaults|`json`|
|tracking_specs|`json`|
|updated_time|`timestamp[us, tz=UTC]`|