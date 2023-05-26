# Table: googleads_ad_group_criterion_labels

This table shows data for Google Ads Ad Group Criterion Labels.

https://developers.google.com/google-ads/api/reference/rpc/v13/AdGroupCriterionLabel

The composite primary key for this table is (**customer_id**, **resource_name**, **ad_group_criterion**).

## Relations

This table depends on [googleads_ad_group_criteria](googleads_ad_group_criteria).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|customer_id (PK)|int64|
|resource_name (PK)|utf8|
|ad_group_criterion (PK)|utf8|
|label|utf8|