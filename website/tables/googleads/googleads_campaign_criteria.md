# Table: googleads_campaign_criteria

This table shows data for Google Ads Campaign Criteria.

https://developers.google.com/google-ads/api/reference/rpc/v13/CampaignCriterion

The composite primary key for this table is (**customer_id**, **id**, **resource_name**, **campaign**).

## Relations

This table depends on [googleads_campaigns](googleads_campaigns).

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
|campaign (PK)|String|
|criterion_id|Int|
|display_name|String|
|bid_modifier|Float|
|negative|Bool|
|type|String|
|status|String|