# Table: googleads_ad_group_ads

This table shows data for Google Ads Ad Group Ads.

https://developers.google.com/google-ads/api/reference/rpc/v13/AdGroupAd

The composite primary key for this table is (**customer_id**, **id**, **resource_name**, **ad_group**).

## Relations

This table depends on [googleads_ad_groups](googleads_ad_groups).

The following tables depend on googleads_ad_group_ads:
  - [googleads_ad_group_ad_labels](googleads_ad_group_ad_labels)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|customer_id (PK)|`int64`|
|id (PK)|`int64`|
|resource_name (PK)|`utf8`|
|status|`utf8`|
|ad_group (PK)|`utf8`|
|ad|`json`|
|policy_summary|`json`|
|ad_strength|`utf8`|
|action_items|`list<item: utf8, nullable>`|
|labels|`list<item: utf8, nullable>`|