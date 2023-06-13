# Table: googleads_campaign_criteria

This table shows data for Google Ads Campaign Criteria.

https://developers.google.com/google-ads/api/reference/rpc/v13/CampaignCriterion

The composite primary key for this table is (**customer_id**, **id**, **resource_name**, **campaign**).

## Relations

This table depends on [googleads_campaigns](googleads_campaigns).

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
|campaign (PK)|`utf8`|
|criterion_id|`int64`|
|display_name|`utf8`|
|bid_modifier|`float64`|
|negative|`bool`|
|type|`utf8`|
|status|`utf8`|