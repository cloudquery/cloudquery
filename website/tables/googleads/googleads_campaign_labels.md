# Table: googleads_campaign_labels

This table shows data for Google Ads Campaign Labels.

https://developers.google.com/google-ads/api/reference/rpc/v13/CampaignLabel

The composite primary key for this table is (**customer_id**, **resource_name**, **campaign**).

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
|resource_name (PK)|`utf8`|
|campaign (PK)|`utf8`|
|label|`utf8`|