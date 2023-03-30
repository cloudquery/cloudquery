# Table: googleads_ad_group_criterion_labels

This table shows data for Google Ads Ad Group Criterion Labels.

https://developers.google.com/google-ads/api/reference/rpc/v13/AdGroupCriterionLabel

The composite primary key for this table is (**customer_id**, **resource_name**, **ad_group_criterion**).

## Relations

This table depends on [googleads_ad_group_criteria](googleads_ad_group_criteria).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|customer_id (PK)|Int|
|resource_name (PK)|String|
|ad_group_criterion (PK)|String|
|label|String|