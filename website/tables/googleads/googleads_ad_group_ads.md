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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|customer_id (PK)|Int|
|id (PK)|Int|
|resource_name (PK)|String|
|status|String|
|ad_group (PK)|String|
|ad|JSON|
|policy_summary|JSON|
|ad_strength|String|
|action_items|StringArray|
|labels|StringArray|